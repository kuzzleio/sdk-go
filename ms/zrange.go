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
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
	"strings"
)

// Zrange returns elements from a sorted set depending on their position in the set, from a start position index to a stop position index (inclusive).
// First position starts at 0.
func (ms *Ms) Zrange(key string, start int, stop int, options types.QueryOptions) ([]*types.MSSortedSet, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zrange",
		Id:         key,
		Start:      start,
		Stop:       stop,
	}

	assignZrangeOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return mapZrangeResults(returnedResult)
}

func assignZrangeOptions(query *types.KuzzleRequest, options types.QueryOptions) {
	opts := make([]interface{}, 0, 1)

	opts = append(opts, "withscores")

	if options != nil {
		if len(options.Limit()) != 0 {
			query.Limit = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(options.Limit())), ","), "[]")
		}
	}

	query.Options = []interface{}(opts)
}

func mapZrangeResults(results []string) ([]*types.MSSortedSet, error) {
	buffer := ""
	sortedSet := make([]*types.MSSortedSet, 0, len(results))

	for _, value := range results {
		if buffer == "" {
			buffer = value
		} else {
			score, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return nil, types.NewError(err.Error())
			}

			sortedSet = append(sortedSet, &types.MSSortedSet{Member: buffer, Score: score})
			buffer = ""
		}
	}

	return sortedSet, nil
}
