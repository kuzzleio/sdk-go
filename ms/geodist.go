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

// Geodist gets the distance between two geospatial members of a key (see geoadd)
func (ms *Ms) Geodist(key string, member1 string, member2 string, options types.QueryOptions) (float64, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "geodist",
		Id:         key,
		Member1:    member1,
		Member2:    member2,
	}

	if options != nil && options.Unit() != "" {
		query.Unit = options.Unit()
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return 0, res.Error
	}
	var returnedResult float64
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
