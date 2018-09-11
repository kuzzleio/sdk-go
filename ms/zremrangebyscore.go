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

// Zremrangebyscore removes members from a sorted set with a score between min and max (inclusive by default).
func (ms *Ms) Zremrangebyscore(key string, min float64, max float64, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zremrangebyscore",
		Id:         key,
		Body:       &body{Min: strconv.FormatFloat(min, 'f', 6, 64), Max: strconv.FormatFloat(max, 'f', 6, 64)},
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
