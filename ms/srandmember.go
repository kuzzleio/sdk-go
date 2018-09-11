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

// Srandmember returns one or more members of a set of unique values, at random.
// If count is provided and is positive, the returned values are unique.
// If count is negative, a set member can be returned multiple times.
func (ms *Ms) Srandmember(key string, options types.QueryOptions) ([]string, error) {
	count := 1

	if options != nil {
		count = options.Count()

		if count < 1 {
			count = 1
		}
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "srandmember",
		Id:         key,
		Count:      count,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	if count == 1 {
		var returnedResult string
		json.Unmarshal(res.Result, &returnedResult)

		return []string{returnedResult}, nil
	} else {
		var returnedResult []string
		json.Unmarshal(res.Result, &returnedResult)

		return returnedResult, nil
	}
}
