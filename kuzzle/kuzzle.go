// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kuzzle provides a Kuzzle Entry point and main struct for the entire SDK
package kuzzle

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/auth"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/document"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/protocol"
	"github.com/kuzzleio/sdk-go/realtime"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/server"
	"github.com/kuzzleio/sdk-go/types"
)

const (
	version           = "2.0.0"
	MAX_CONNECT_RETRY = 10
)

type OfflineQueueLoader interface {
	Load() []*types.QueryObject
}

type QueueFilter func([]byte) bool

type Kuzzle struct {
	protocol protocol.Protocol

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory map[string]time.Time
	volatile       types.VolatileData

	eventListeners     map[int]map[chan<- json.RawMessage]bool
	eventListenersOnce map[int]map[chan<- json.RawMessage]bool

	autoReconnect         bool
	autoQueue             bool
	autoReplay            bool
	autoResubscribe       bool
	offlineQueue          []*types.QueryObject
	offlineQueueLoader    OfflineQueueLoader
	queueFilter           QueueFilter
	queueMaxSize          int
	queueTTL              time.Duration
	reconnectionDelay     time.Duration
	replayInterval        time.Duration
	retryingToConnect     bool
	stopRetryingToConnect bool
	queuing               bool
	nbConnectionRetried   int

	MemoryStorage *ms.Ms
	Security      *security.Security
	Realtime      *realtime.Realtime
	Auth          *auth.Auth
	Server        *server.Server
	Document      *document.Document
	Index         *index.Index
	Collection    *collection.Collection
}

// NewKuzzle is the Kuzzle constructor
func NewKuzzle(c protocol.Protocol, options types.Options) (*Kuzzle, error) {
	if c == nil {
		return nil, types.NewError("Connection is nil")
	}

	if options == nil {
		options = types.NewOptions()
	}

	k := &Kuzzle{
		protocol:           c,
		version:            version,
		eventListeners:     make(map[int]map[chan<- json.RawMessage]bool),
		eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool),
		autoQueue:          options.AutoQueue(),
		autoReplay:         options.AutoReplay(),
		autoReconnect:      options.AutoReconnect(),
		autoResubscribe:    options.AutoResubscribe(),
		offlineQueue:       []*types.QueryObject{},
		queueMaxSize:       options.QueueMaxSize(),
		queueTTL:           options.QueueTTL(),
		reconnectionDelay:  options.ReconnectionDelay(),
		replayInterval:     options.ReplayInterval(),
		queuing:            false,
	}

	if options.OfflineMode() == types.Auto {
		k.autoQueue = true
		k.autoReplay = true
		k.autoReconnect = true
		k.autoResubscribe = true
	}

	k.RequestHistory = k.protocol.RequestHistory()
	k.MemoryStorage = &ms.Ms{k}
	k.Security = security.NewSecurity(k)
	k.Auth = auth.NewAuth(k)
	k.Realtime = realtime.NewRealtime(k)

	k.Server = server.NewServer(k)
	k.Collection = collection.NewCollection(k)
	k.Document = document.NewDocument(k)
	k.Index = index.NewIndex(k)

	return k, nil
}

// Connect connects to a Kuzzle instance.
func (k *Kuzzle) Connect() error {
	if k.autoQueue {
		k.queuing = true
	}

	wasConnected, err := k.protocol.Connect()

	if err != nil {
		if k.autoReconnect && !k.retryingToConnect && !k.stopRetryingToConnect && k.nbConnectionRetried < MAX_CONNECT_RETRY {
			k.retryingToConnect = true
			time.Sleep(k.reconnectionDelay * time.Millisecond)
			k.retryingToConnect = false
			k.nbConnectionRetried++
			return k.Connect()
		}

		return types.NewError(err.Error())
	}

	k.nbConnectionRetried = 0
	k.stopRetryingToConnect = false
	k.queuing = false

	if k.lastUrl != k.protocol.Host() {
		k.wasConnected = false
		k.lastUrl = k.protocol.Host()
	}

	if wasConnected {
		if k.jwt != "" {
			go func() {
				res, err := k.Auth.CheckToken(k.jwt)

				if err != nil || !res.Valid {
					k.jwt = ""
					k.protocol.EmitEvent(event.TokenExpired, nil)
				}
				k.EmitEvent(event.Reconnected, nil)
			}()
		} else {
			k.EmitEvent(event.Reconnected, nil)
		}
	} else {
		k.EmitEvent(event.Connected, nil)
	}

	if k.autoQueue {
		k.queuing = false
	}
	if k.autoReplay {
		k.PlayQueue()
	}

	// on network error
	ee := make(chan json.RawMessage)
	go func() {
		for err := range <-ee {
			if k.autoQueue {
				k.queuing = true
			}

			k.EmitEvent(event.NetworkError, err)
		}
	}()
	k.protocol.AddListener(event.NetworkError, ee)

	return nil
}

// Jwt get internal jwtToken used to request kuzzle.
func (k *Kuzzle) Jwt() string {
	return k.jwt
}

func (k *Kuzzle) SetJwt(token string) {
	k.jwt = token

	if token != "" {
		k.protocol.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: true})
	}
}

// UnsetJwt unset the authentication token and cancel all subscriptions
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	k.protocol.CancelSubs()
}

func (k *Kuzzle) RegisterSub(channel, roomId string, filters json.RawMessage, subscribeToSelf bool, notifChan chan<- types.NotificationResult, onReconnectChannel chan<- interface{}) {
	k.protocol.RegisterSub(channel, roomId, filters, subscribeToSelf, notifChan, onReconnectChannel)
}

func (k *Kuzzle) UnregisterSub(roomId string) {
	k.protocol.UnregisterSub(roomId)
}

// AutoQueue returns the Kuzzle socket AutoQueue field value
func (k *Kuzzle) AutoQueue() bool {
	return k.autoQueue
}

// AutoReconnect returns the Kuzzle socket AutoReconnect field value
func (k *Kuzzle) AutoReconnect() bool {
	return k.autoReconnect
}

// AutoResubscribe returns the Kuzzle socket AutoQueue field value
func (k *Kuzzle) AutoResubscribe() bool {
	return k.autoResubscribe
}

// AutoReplay returns the Kuzzle socket AutoReplay field value
func (k *Kuzzle) AutoReplay() bool {
	return k.autoReplay
}

// OfflineQueue returns the Kuzzle socket OfflineQueue field value
func (k *Kuzzle) OfflineQueue() []*types.QueryObject {
	return k.offlineQueue
}

// OfflineQueueLoader returns the Kuzzle socket OfflineQueueLoader field value
func (k *Kuzzle) OfflineQueueLoader() OfflineQueueLoader {
	return k.offlineQueueLoader
}

// QueueFilter returns the Kuzzle socket QueueFilter field value
func (k *Kuzzle) QueueFilter() QueueFilter {
	return k.queueFilter
}

// QueueMaxSize returns the Kuzzle socket QueueMaxSize field value
func (k *Kuzzle) QueueMaxSize() int {
	return k.queueMaxSize
}

// QueueTTL returns the Kuzzle socket QueueTTL field value
func (k *Kuzzle) QueueTTL() time.Duration {
	return k.queueTTL
}

// ReplayInterval returns the Kuzzle socket ReplayInterval field value
func (k *Kuzzle) ReplayInterval() time.Duration {
	return k.replayInterval
}

// ReconnectionDelay returns the Kuzzle socket ReconnectionDelay field value
func (k *Kuzzle) ReconnectionDelay() time.Duration {
	return k.reconnectionDelay
}

// SslConnection returns the Kuzzle socket SslConnection field value
func (k *Kuzzle) SslConnection() bool {
	return k.protocol.SslConnection()
}

// SetAutoQueue sets the Kuzzle socket AutoQueue field with the given value
func (k *Kuzzle) SetAutoQueue(v bool) {
	k.autoQueue = v
}

// SetAutoReplay sets the Kuzzle socket AutoReplay field with the given value
func (k *Kuzzle) SetAutoReplay(v bool) {
	k.autoReplay = v
}

// SetOfflineQueueLoader sets the Kuzzle socket OfflineQueueLoader field with given value
func (k *Kuzzle) SetOfflineQueueLoader(v OfflineQueueLoader) {
	k.offlineQueueLoader = v
}

// SetQueueFilter sets the Kuzzle socket QueueFilter field with given value
func (k *Kuzzle) SetQueueFilter(v QueueFilter) {
	k.queueFilter = v
}

// SetQueueMaxSize sets the Kuzzle socket QueueMaxSize field with the given value
func (k *Kuzzle) SetQueueMaxSize(v int) {
	k.queueMaxSize = v
}

// SetQueueTTL sets the Kuzzle socket QueueTTL field with the given value
func (k *Kuzzle) SetQueueTTL(v time.Duration) {
	k.queueTTL = v
}

// SetReplayInterval sets the Kuzzle socket ReplayInterval field with the given value
func (k *Kuzzle) SetReplayInterval(v time.Duration) {
	k.replayInterval = v
}

func (k *Kuzzle) Volatile() types.VolatileData {
	return k.volatile
}

func (k *Kuzzle) SetVolatile(v types.VolatileData) {
	k.volatile = v
}

func (k *Kuzzle) EmitEvent(e int, arg interface{}) {
	json, _ := json.Marshal(arg)
	for c := range k.eventListeners[e] {
		c <- json
	}
	for c := range k.eventListenersOnce[e] {
		c <- json
		close(c)
		delete(k.eventListenersOnce[e], c)
	}
}
