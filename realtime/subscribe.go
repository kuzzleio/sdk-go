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

package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Subscribe permits to join a previously created subscription
func (r *Realtime) Subscribe(index, collection string, filters json.RawMessage, cb chan<- types.NotificationResult, options types.RoomOptions) (*types.SubscribeResult, error) {
	if (index == "" || collection == "") || (filters == nil || cb == nil) {
		return nil, types.NewError("Realtime.Subscribe: index, collection, filters and notificationChannel required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "subscribe",
		Index:      index,
		Collection: collection,
		Body:       filters,
	}

	opts := types.NewQueryOptions()

	if options != nil {
		query.Users = options.Users()
		query.State = options.State()
		query.Scope = options.Scope()

		opts.SetVolatile(options.Volatile())
	} else {
		options = types.NewRoomOptions()
		query.Users = options.Users()
		query.State = options.State()
		query.Scope = options.Scope()
	}

	go r.k.Query(query, opts, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var resSub types.SubscribeResult

	json.Unmarshal(res.Result, &resSub)

	onReconnect := make(chan interface{})

	r.k.RegisterSub(resSub.Channel, resSub.Room, filters, options.SubscribeToSelf(), cb, onReconnect)

	go func() {
		<-onReconnect

		if r.k.AutoResubscribe() {
			go r.k.Query(query, opts, result)
		}
	}()

	return &resSub, nil
}
