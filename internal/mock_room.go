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

import "github.com/kuzzleio/sdk-go/types"

type MockedRoom struct {
	MockedSubscribe func()
}

func (m MockedRoom) Subscribe(realtimeNotificationChannel chan<- *types.NotificationResult) {
	m.MockedSubscribe()
}
func (m MockedRoom) Unsubscribe() error {
	return nil
}

func (m MockedRoom) GetRealtimeChannel() chan<- *types.NotificationResult {
	return make(chan<- *types.NotificationResult)
}

func (m MockedRoom) GetResponseChannel() chan<- *types.SubscribeResponse {
	return make(chan<- *types.SubscribeResponse)
}

func (m MockedRoom) RoomId() string {
	return ""
}

func (m MockedRoom) Filters() interface{} {
	return nil
}

func (m MockedRoom) Channel() string {
	return ""
}

func (m MockedRoom) Id() string {
	return ""
}

func (m MockedRoom) SubscribeToSelf() bool {
	return true
}
