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

import "sync"

// SubscribeResponse is a response given after a new subscriptions
type SubscribeResponse struct {
	Room  IRoom
	Error error
}

// IRoom provides functions to manage Room
type IRoom interface {
	Subscribe(realtimeNotificationChannel chan<- NotificationResult)
	Unsubscribe() error
	RealtimeChannel() chan<- NotificationResult
	ResponseChannel() chan SubscribeResponse
	RoomId() string
	Filters() interface{}
	Channel() string
	Id() string
	SubscribeToSelf() bool
}

type RoomList = sync.Map
