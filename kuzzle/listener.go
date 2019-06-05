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

package kuzzle

import (
	"encoding/json"
)

// AddListener Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (k *Kuzzle) AddListener(event int, channel chan<- json.RawMessage) {
	if k.eventListeners[event] == nil {
		k.eventListeners[event] = make(map[chan<- json.RawMessage]bool)
	}
	k.eventListeners[event][channel] = true
}

// On is an alias to the AddListener function
func (k *Kuzzle) On(event int, channel chan<- json.RawMessage) {
	k.AddListener(event, channel)
}

// RemoveAllListeners removes all listener by event type or all listener if event == -1
func (k *Kuzzle) RemoveAllListeners(event int) {
	for key := range k.eventListeners {
		if event == key || event == -1 {
			delete(k.eventListeners, key)
		}
	}

	for key := range k.eventListenersOnce {
		if event == key || event == -1 {
			delete(k.eventListenersOnce, key)
		}
	}
}

// RemoveListener removes a listener
func (k *Kuzzle) RemoveListener(event int, channel chan<- json.RawMessage) {
	delete(k.eventListeners[event], channel)
	delete(k.eventListenersOnce[event], channel)
}

func (k *Kuzzle) Once(event int, channel chan<- json.RawMessage) {
	if k.eventListenersOnce[event] == nil {
		k.eventListenersOnce[event] = make(map[chan<- json.RawMessage]bool)
	}
	k.eventListenersOnce[event][channel] = true
}

func (k *Kuzzle) ListenerCount(event int) int {
	return len(k.eventListenersOnce[event]) + len(k.eventListeners[event])
}
