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

type ZScanResponse struct {
	Cursor int
	Values []string
}

// Zscan is identical to scan, except that zscan iterates the members held by a sorted set.
func (ms *Ms) Zscan(key string, cursor int, options types.QueryOptions) (*types.MSScanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "zscan",
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

	var scanResponse []interface{}
	json.Unmarshal(res.Result, &scanResponse)

	return formatZscanResponse(scanResponse), nil
}

func formatZscanResponse(response []interface{}) *types.MSScanResponse {
	formatedResponse := &types.MSScanResponse{}

	for _, element := range response {
		switch vf := element.(type) {
		case string:
			formatedResponse.Cursor, _ = strconv.Atoi(vf)
		case []interface{}:
			values := make([]string, 0, len(vf))

			for _, v := range vf {
				switch vv := v.(type) {
				case string:
					values = append(values, vv)
				}
			}

			formatedResponse.Values = values
		}
	}

	return formatedResponse
}
