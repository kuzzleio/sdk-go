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

// Zadd adds the specified elements to the sorted set stored at key.
// If the key does not exist, it is created, holding an empty sorted set. If it already exists and does not hold a sorted set, an error is returned.
// Scores are expressed as floating point numbers.
// If a member to insert is already in the sorted set,
// its score is updated and the member is reinserted at the right position in the set.
func (ms *Ms) Zadd(key string, elements []*types.MSSortedSet, options types.QueryOptions) (int, error) {
	if len(elements) == 0 {
		return 0, types.NewError("Ms.Zadd: please provide at least one element", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Elements []*types.MSSortedSet `json:"elements"`
		Nx       bool                 `json:"nx,omitempty"`
		Xx       bool                 `json:"xx,omitempty"`
		Ch       bool                 `json:"ch,omitempty"`
		Incr     bool                 `json:"incr,omitempty"`
	}

	bodyContent := body{Elements: elements}

	if options != nil {
		bodyContent.Nx = options.Nx()
		bodyContent.Xx = options.Xx()
		bodyContent.Ch = options.Ch()
		bodyContent.Incr = options.Incr()
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zadd",
		Id:         key,
		Body:       &bodyContent,
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
