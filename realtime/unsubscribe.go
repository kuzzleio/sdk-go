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
	"github.com/kuzzleio/sdk-go/types"
)

// Unsubscribe instructs Kuzzle to detach you from its subscribers for the given room
func (r *Realtime) Unsubscribe(roomID string, options types.QueryOptions) error {
	if roomID == "" {
		return types.NewError("Realtime.Unsubscribe: roomID required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "unsubscribe",
		Body: struct {
			RoomID string `json:"roomId"`
		}{roomID},
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}

	r.k.UnregisterSub(roomID)

	return nil
}
