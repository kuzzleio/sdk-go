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
	"strconv"
)

// ZincrBy increments the score of a member in a sorted set by the provided value.
func (ms *Ms) Zincrby(key string, member string, increment float64, options types.QueryOptions) (float64, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Member    string  `json:"member"`
		Increment float64 `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zincrby",
		Id:         key,
		Body:       &body{Member: member, Increment: increment},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return 0, res.Error
	}

	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	converted, err := strconv.ParseFloat(returnedResult, 64)

	if err != nil {
		err = types.NewError(err.Error())
	}

	return converted, err
}
