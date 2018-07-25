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

// FlushQueue empties the offline queue without replaying it.
func (k *Kuzzle) FlushQueue() {
	k.socket.ClearQueue()
}

// ReplayQueue replays the requests queued during offline mode.
// Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
func (k *Kuzzle) ReplayQueue() {
	k.socket.ReplayQueue()
}

// StartQueuing start the requests queuing.
func (k *Kuzzle) StartQueuing() {
	k.socket.StartQueuing()
}

// StopQueuing stop the requests queuing.
func (k *Kuzzle) StopQueuing() {
	k.socket.StopQueuing()
}
