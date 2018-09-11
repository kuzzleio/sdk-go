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

// Sort sorts and returns elements contained in a list, a set of unique values or a sorted set.
// By default, sorting is numeric and elements are compared by their value interpreted
// as double precision floating point number.
func (ms *Ms) Sort(key string, options types.QueryOptions) ([]string, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sort",
		Id:         key,
	}

	if options != nil {
		type body struct {
			Limit     []int    `json:"limit,omitempty"`
			By        string   `json:"by,omitempty"`
			Direction string   `json:"direction,omitempty"`
			Get       []string `json:"get,omitempty"`
			Alpha     bool     `json:"alpha,omitempty"`
		}

		bodyContent := &body{}

		if options.By() != "" {
			bodyContent.By = options.By()
		}

		if options.Direction() != "" {
			bodyContent.Direction = options.Direction()
		}

		if options.Get() != nil {
			bodyContent.Get = options.Get()
		}

		if options.Limit() != nil {
			bodyContent.Limit = options.Limit()
		}

		bodyContent.Alpha = options.Alpha()

		query.Body = bodyContent
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
