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

// Sscan is identical to scan, except that sscan iterates the members held by a set of unique values.
func (ms *Ms) Sscan(key string, cursor int, options types.QueryOptions) (*types.MSScanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sscan",
		Id:         key,
		Cursor:     cursor,
	}

	if options != nil {
		if options.Count() != 0 {
			query.Count = options.Count()
		}

		if options.Match() != "" {
			query.Match = options.Match()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var sscanResponse = &types.MSScanResponse{}
	json.Unmarshal(res.Result, sscanResponse)

	return sscanResponse, nil
}
