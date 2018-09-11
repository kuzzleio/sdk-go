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

package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SinterStore computes the intersection of the provided sets of unique values and stores
// the result in the destination key.
// If the destination key already exists, it is overwritten.
func (ms *Ms) Sinterstore(destination string, keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, types.NewError("Ms.Sinterstore: please provide at least one key to intersect", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Destination string   `json:"destination"`
		Keys        []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sinterstore",
		Body:       &body{Destination: destination, Keys: keys},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
