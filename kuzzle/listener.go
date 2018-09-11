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

// AddListener Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (k *Kuzzle) AddListener(event int, channel chan<- interface{}) {
	k.socket.AddListener(event, channel)
}

// On is an alias to the AddListener function
func (k *Kuzzle) On(event int, channel chan<- interface{}) {
	k.socket.AddListener(event, channel)
}

// RemoveAllListeners removes all listener by event type or all listener if event == -1
func (k *Kuzzle) RemoveAllListeners(event int) {
	k.socket.RemoveAllListeners(event)
}

// RemoveListener removes a listener
func (k *Kuzzle) RemoveListener(event int, channel chan<- interface{}) {
	k.socket.RemoveListener(event, channel)
}

func (k *Kuzzle) Once(event int, channel chan<- interface{}) {
	k.socket.Once(event, channel)
}

func (k *Kuzzle) ListenerCount(event int) int {
	return k.socket.ListenerCount(event)
}
