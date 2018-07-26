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

package connection

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/types"
)

// Connection provides functions to manage many connection type (websocket...)
type Connection interface {
	AddListener(event int, channel chan<- interface{})
	RemoveListener(event int, channel chan<- interface{})
	RemoveAllListeners(event int)
	Once(event int, channel chan<- interface{})
	ListenerCount(event int) int
	Connect() (bool, error)
	Send([]byte, types.QueryOptions, chan<- *types.KuzzleResponse, string) error
	Close() error
	State() int
	EmitEvent(int, interface{})
	RegisterSub(string, string, json.RawMessage, bool, chan<- types.KuzzleNotification, chan<- interface{})
	UnregisterSub(string)
	CancelSubs()
	RequestHistory() map[string]time.Time
	StartQueuing()
	StopQueuing()
	PlayQueue()
	ClearQueue()

	// property getters
	AutoQueue() bool
	AutoReconnect() bool
	AutoResubscribe() bool
	AutoReplay() bool
	Host() string
	OfflineQueue() []*types.QueryObject
	OfflineQueueLoader() OfflineQueueLoader
	Port() int
	QueueFilter() QueueFilter
	QueueMaxSize() int
	QueueTTL() time.Duration
	ReplayInterval() time.Duration
	ReconnectionDelay() time.Duration
	SslConnection() bool

	// property setters
	SetAutoQueue(bool)
	SetAutoReplay(bool)
	SetOfflineQueueLoader(OfflineQueueLoader)
	SetQueueFilter(QueueFilter)
	SetQueueMaxSize(int)
	SetQueueTTL(time.Duration)
	SetReplayInterval(time.Duration)
}

type OfflineQueueLoader interface {
	Load() []*types.QueryObject
}

type QueueFilter func([]byte) bool
