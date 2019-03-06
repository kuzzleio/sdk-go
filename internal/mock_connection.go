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

package internal

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/mock"
)

type MockedConnection struct {
	mock.Mock
	MockSend               func([]byte, types.QueryOptions) *types.KuzzleResponse
	MockEmitEvent          func(int, interface{})
	MockAddListener        func(int, chan<- json.RawMessage)
	MockRemoveAllListeners func(int)
	MockRemoveListener     func(int, chan<- json.RawMessage)
	MockOnce               func(int, chan<- json.RawMessage)
	MockListenerCount      func(int) int

	state int
}

func (c *MockedConnection) Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error {
	if c.MockSend != nil {
		responseChannel <- c.MockSend(query, options)
	}

	return nil
}

func (c *MockedConnection) Connect() (bool, error) {
	return false, nil
}

func (c *MockedConnection) Close() error {
	return nil
}

func (c *MockedConnection) AddListener(event int, channel chan<- json.RawMessage) {
	if c.MockAddListener != nil {
		c.MockAddListener(event, channel)
	}
}

func (c *MockedConnection) Once(event int, channel chan<- json.RawMessage) {
	if c.MockOnce != nil {
		c.MockOnce(event, channel)
	}
}

func (c *MockedConnection) ListenerCount(event int) int {
	if c.MockListenerCount != nil {
		return c.MockListenerCount(event)
	}
	return -1
}

func (c *MockedConnection) State() int {
	return c.state
}

func (c *MockedConnection) EmitEvent(event int, arg interface{}) {
	if c.MockEmitEvent != nil {
		c.MockEmitEvent(event, arg)
	}
}

func (c *MockedConnection) RegisterSub(channel, roomID string, filters json.RawMessage, subscribeToSelf bool, notifChan chan<- types.NotificationResult, onReconnectChan chan<- interface{}) {
}

func (c *MockedConnection) CancelSubs() {}

func (c *MockedConnection) UnregisterSub(roomID string) {}

func (c *MockedConnection) RequestHistory() map[string]time.Time {
	r := make(map[string]time.Time)

	return r
}

func (c *MockedConnection) RenewSubscriptions() {}

func (c *MockedConnection) RemoveListener(event int, channel chan<- json.RawMessage) {
	if c.MockRemoveListener != nil {
		c.MockRemoveListener(event, channel)
	}
}

func (c *MockedConnection) RemoveAllListeners(event int) {
	if c.MockRemoveAllListeners != nil {
		c.MockRemoveAllListeners(event)
	}
}

func (c *MockedConnection) AutoResubscribe() bool {
	return true
}

func (c *MockedConnection) Host() string {
	return ""
}

func (c *MockedConnection) IsReady() bool {
	return true
}

func (c *MockedConnection) Port() int {
	return 0
}

func (c *MockedConnection) ReconnectionDelay() time.Duration {
	return 0
}

func (c *MockedConnection) SslConnection() bool {
	return false
}

// mock specific functions
func (c *MockedConnection) SetState(value int) {
	c.state = value
}
