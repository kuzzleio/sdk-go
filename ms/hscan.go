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

type HscanResponse struct {
	Cursor int
	Values []string
}

// Hscan is identical to scan, except that hscan iterates the fields contained in a hash.
func (ms *Ms) Hscan(key string, cursor int, options types.QueryOptions) (*HscanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hscan",
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

	var stringResult []interface{}
	json.Unmarshal(res.Result, &stringResult)

	returnedResult := &HscanResponse{}

	tmp, err := strconv.ParseInt(stringResult[0].(string), 10, 0)
	if err != nil {
		return returnedResult, types.NewError(err.Error())
	}
	returnedResult.Cursor = int(tmp)

	tmpS := stringResult[1].([]interface{})

	for _, value := range tmpS {
		returnedResult.Values = append(returnedResult.Values, value.(string))
	}

	return returnedResult, nil
}
