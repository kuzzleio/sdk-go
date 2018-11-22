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

const version = "1.0.0"

type Kuzzle struct {
	socket protocol.Protocol

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory map[string]time.Time
	volatile       types.VolatileData

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
		socket:  c,
		version: version,
	}

	k.RequestHistory = k.socket.RequestHistory()
	k.MemoryStorage = &ms.Ms{k}
	k.Security = security.NewSecurity(k)
	k.Auth = auth.NewAuth(k)
	k.Realtime = realtime.NewRealtime(k)

	k.RequestHistory = k.socket.RequestHistory()

	k.Server = server.NewServer(k)
	k.Collection = collection.NewCollection(k)
	k.Document = document.NewDocument(k)
	k.Index = index.NewIndex(k)

	return k, nil
}

// Connect connects to a Kuzzle instance using the provided host and port.
func (k *Kuzzle) Connect() error {
	wasConnected, err := k.socket.Connect()
	if err == nil {
		if k.lastUrl != k.socket.Host() {
			k.wasConnected = false
			k.lastUrl = k.socket.Host()
		}

		if wasConnected {
			if k.jwt != "" {
				go func() {
					res, err := k.Auth.CheckToken(k.jwt)

					if err != nil {
						k.jwt = ""
						k.socket.EmitEvent(event.TokenExpired, nil)
						return
					}

					if !res.Valid {
						k.jwt = ""
						k.socket.EmitEvent(event.TokenExpired, nil)
					}
				}()
			}
		}
		return nil
	}
	return types.NewError(err.Error())
}

// Jwt get internal jwtToken used to request kuzzle.
func (k *Kuzzle) Jwt() string {
	return k.jwt
}

func (k *Kuzzle) SetJwt(token string) {
	k.jwt = token

	if token != "" {
		k.socket.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: true})
	}
}

// UnsetJwt unset the authentication token and cancel all subscriptions
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	k.socket.CancelSubs()
}

func (k *Kuzzle) RegisterSub(channel, roomId string, filters json.RawMessage, subscribeToSelf bool, notifChan chan<- types.NotificationResult, onReconnectChannel chan<- interface{}) {
	k.socket.RegisterSub(channel, roomId, filters, subscribeToSelf, notifChan, onReconnectChannel)
}

func (k *Kuzzle) UnregisterSub(roomId string) {
	k.socket.UnregisterSub(roomId)
}

// State returns the Kuzzle socket state
func (k *Kuzzle) State() int {
	return k.socket.State()
}

// AutoQueue returns the Kuzzle socket AutoQueue field value
func (k *Kuzzle) AutoQueue() bool {
	return k.socket.AutoQueue()
}

// AutoReconnect returns the Kuzzle socket AutoReconnect field value
func (k *Kuzzle) AutoReconnect() bool {
	return k.socket.AutoReconnect()
}

// AutoResubscribe returns the Kuzzle socket AutoQueue field value
func (k *Kuzzle) AutoResubscribe() bool {
	return k.socket.AutoResubscribe()
}

// AutoReplay returns the Kuzzle socket AutoReplay field value
func (k *Kuzzle) AutoReplay() bool {
	return k.socket.AutoReplay()
}

// Host returns the Kuzzle socket Host field value
func (k *Kuzzle) Host() string {
	return k.socket.Host()
}

// OfflineQueue returns the Kuzzle socket OfflineQueue field value
func (k *Kuzzle) OfflineQueue() []*types.QueryObject {
	return k.socket.OfflineQueue()
}

// OfflineQueueLoader returns the Kuzzle socket OfflineQueueLoader field value
func (k *Kuzzle) OfflineQueueLoader() protocol.OfflineQueueLoader {
	return k.socket.OfflineQueueLoader()
}

// Port returns the Kuzzle socket Port field value
func (k *Kuzzle) Port() int {
	return k.socket.Port()
}

// QueueFilter returns the Kuzzle socket QueueFilter field value
func (k *Kuzzle) QueueFilter() protocol.QueueFilter {
	return k.socket.QueueFilter()
}

// QueueMaxSize returns the Kuzzle socket QueueMaxSize field value
func (k *Kuzzle) QueueMaxSize() int {
	return k.socket.QueueMaxSize()
}

// QueueTTL returns the Kuzzle socket QueueTTL field value
func (k *Kuzzle) QueueTTL() time.Duration {
	return k.socket.QueueTTL()
}

// ReplayInterval returns the Kuzzle socket ReplayInterval field value
func (k *Kuzzle) ReplayInterval() time.Duration {
	return k.socket.ReplayInterval()
}

// ReconnectionDelay returns the Kuzzle socket ReconnectionDelay field value
func (k *Kuzzle) ReconnectionDelay() time.Duration {
	return k.socket.ReconnectionDelay()
}

// SslConnection returns the Kuzzle socket SslConnection field value
func (k *Kuzzle) SslConnection() bool {
	return k.socket.SslConnection()
}

// SetAutoQueue sets the Kuzzle socket AutoQueue field with the given value
func (k *Kuzzle) SetAutoQueue(v bool) {
	k.socket.SetAutoQueue(v)
}

// SetAutoReplay sets the Kuzzle socket AutoReplay field with the given value
func (k *Kuzzle) SetAutoReplay(v bool) {
	k.socket.SetAutoReplay(v)
}

// SetOfflineQueueLoader sets the Kuzzle socket OfflineQueueLoader field with given value
func (k *Kuzzle) SetOfflineQueueLoader(v protocol.OfflineQueueLoader) {
	k.socket.SetOfflineQueueLoader(v)
}

// SetQueueFilter sets the Kuzzle socket QueueFilter field with given value
func (k *Kuzzle) SetQueueFilter(v protocol.QueueFilter) {
	k.socket.SetQueueFilter(v)
}

// SetQueueMaxSize sets the Kuzzle socket QueueMaxSize field with the given value
func (k *Kuzzle) SetQueueMaxSize(v int) {
	k.socket.SetQueueMaxSize(v)
}

// SetQueueTTL sets the Kuzzle socket QueueTTL field with the given value
func (k *Kuzzle) SetQueueTTL(v time.Duration) {
	k.socket.SetQueueTTL(v)
}

// SetReplayInterval sets the Kuzzle socket ReplayInterval field with the given value
func (k *Kuzzle) SetReplayInterval(v time.Duration) {
	k.socket.SetReplayInterval(v)
}

func (k *Kuzzle) Volatile() types.VolatileData {
	return k.volatile
}

func (k *Kuzzle) SetVolatile(v types.VolatileData) {
	k.volatile = v
}

func (k *Kuzzle) EmitEvent(e int, arg interface{}) {
	k.socket.EmitEvent(e, arg)
}
