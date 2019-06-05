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

// Count returns the number of other subscriptions on that room.
func (r *Realtime) Count(roomID string, options types.QueryOptions) (int, error) {
	if roomID == "" {
		return -1, types.NewError("Realtime.Count: roomID required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "count",
		Body: struct {
			RoomID string `json:"roomId"`
		}{roomID},
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return -1, res.Error
	}

	var countRes struct {
		Count int `json:"count"`
	}

	json.Unmarshal(res.Result, &countRes)

	return countRes.Count, nil
}
