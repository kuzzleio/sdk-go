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

package websocket

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/stretchr/testify/assert"
)

func TestAddListener(t *testing.T) {
	c := WebSocket{eventListeners: make(map[int]map[chan<- json.RawMessage]bool)}
	c.Connect()
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))
}

func TestOnce(t *testing.T) {
	c := WebSocket{eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool)}
	listener := make(chan json.RawMessage)
	go func() {
		<-listener
	}()
	c.Once(event.LoginAttempt, listener)
	assert.Equal(t, 1, len(c.eventListenersOnce[event.LoginAttempt]))
	c.EmitEvent(event.LoginAttempt, nil)
	assert.Equal(t, 0, len(c.eventListenersOnce[event.LoginAttempt]))
}

func TestRemoveAllListeners(t *testing.T) {
	c := WebSocket{
		eventListeners:     make(map[int]map[chan<- json.RawMessage]bool),
		eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool),
	}
	c.Connect()

	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.Disconnected, nil)
	assert.Equal(t, 2, len(c.eventListeners))

	c.Once(event.Connected, nil)
	c.Once(event.Discarded, nil)
	assert.Equal(t, 2, len(c.eventListenersOnce))

	c.RemoveAllListeners(event.LoginAttempt)
	c.RemoveAllListeners(event.Disconnected)
	c.RemoveAllListeners(event.Connected)
	c.RemoveAllListeners(event.Discarded)

	assert.Equal(t, 0, len(c.eventListeners))
	assert.Equal(t, 0, len(c.eventListenersOnce))
}

func TestRemoveListener(t *testing.T) {
	c := WebSocket{eventListeners: make(map[int]map[chan<- json.RawMessage]bool)}

	listener := make(chan json.RawMessage)
	c.AddListener(event.LoginAttempt, listener)
	c.AddListener(event.Disconnected, listener)
	assert.Equal(t, 1, len(c.eventListeners[event.LoginAttempt]))

	c.RemoveListener(event.LoginAttempt, listener)

	assert.Equal(t, 0, len(c.eventListeners[event.LoginAttempt]))
	assert.Equal(t, 1, len(c.eventListeners[event.Disconnected]))

	c.RemoveListener(event.Disconnected, listener)
	assert.Equal(t, 0, len(c.eventListeners[event.Disconnected]))
}

func TestListenerCount(t *testing.T) {
	c := WebSocket{
		eventListeners:     make(map[int]map[chan<- json.RawMessage]bool),
		eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool),
	}
	c.Connect()

	ch := make(chan<- json.RawMessage)
	c.AddListener(event.LoginAttempt, nil)
	c.AddListener(event.LoginAttempt, ch)

	assert.Equal(t, 2, c.ListenerCount(event.LoginAttempt))

	c.Once(event.LoginAttempt, nil)
	c.Once(event.LoginAttempt, ch)

	assert.Equal(t, 4, c.ListenerCount(event.LoginAttempt))
}

func TestEmitEvent(t *testing.T) {
	c := WebSocket{
		eventListeners:     make(map[int]map[chan<- json.RawMessage]bool),
		eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool),
	}
	listener := make(chan json.RawMessage)
	go func() {
		for {
			<-listener
		}
	}()

	c.AddListener(event.LoginAttempt, listener)
	c.Once(event.LoginAttempt, listener)
	assert.Equal(t, 1, len(c.eventListenersOnce[event.LoginAttempt]))
	c.EmitEvent(event.LoginAttempt, nil)
	assert.Equal(t, 1, len(c.eventListeners[event.LoginAttempt]))
}
