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

package types

type roomOptions struct {
	scope           string
	state           string
	users           string
	subscribeToSelf bool
	volatile        VolatileData
	autoResubscribe bool
}

// RoomOptions provides a public access to roomOptions private struct
type RoomOptions interface {
	Scope() string
	SetScope(string) *roomOptions
	State() string
	SetState(string) *roomOptions
	Users() string
	SetUsers(string) *roomOptions
	SubscribeToSelf() bool
	SetSubscribeToSelf(bool) *roomOptions
	Volatile() VolatileData
	SetVolatile(VolatileData) *roomOptions
	AutoResubscribe() bool
	SetAutoResubscribe(bool)
}

func (ro roomOptions) Scope() string {
	return ro.scope
}

func (ro *roomOptions) SetScope(scope string) *roomOptions {
	ro.scope = scope
	return ro
}

func (ro roomOptions) State() string {
	return ro.state
}

func (ro *roomOptions) SetState(state string) *roomOptions {
	ro.state = state
	return ro
}

func (ro roomOptions) Users() string {
	return ro.users
}

func (ro *roomOptions) SetUsers(users string) *roomOptions {
	ro.users = users
	return ro
}

func (ro roomOptions) SubscribeToSelf() bool {
	return ro.subscribeToSelf
}

func (ro *roomOptions) SetSubscribeToSelf(subscribeToSelf bool) *roomOptions {
	ro.subscribeToSelf = subscribeToSelf
	return ro
}

func (ro roomOptions) Volatile() VolatileData {
	return ro.volatile
}

func (ro *roomOptions) SetVolatile(volatile VolatileData) *roomOptions {
	ro.volatile = volatile
	return ro
}

func (ro roomOptions) AutoResubscribe() bool {
	return ro.autoResubscribe
}

func (ro *roomOptions) SetAutoResubscribe(v bool) {
	ro.autoResubscribe = v
}

// NewRoomOptions instanciates a new RoomOptions
func NewRoomOptions() RoomOptions {
	return &roomOptions{
		scope:           SCOPE_ALL,
		state:           STATE_DONE,
		users:           USERS_NONE,
		subscribeToSelf: true,
	}
}
