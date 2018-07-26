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

// Geopos returns the longitude/latitude values for the provided key's members
func (ms *Ms) Geopos(key string, members []string, options types.QueryOptions) ([]*types.GeoPoint, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "geopos",
		Id:         key,
		Members:    members,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}
	var stringResults [][]string
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]*types.GeoPoint, len(stringResults))

	for i := 0; i < len(stringResults); i++ {
		returnedResults[i] = &types.GeoPoint{}
		tmp, err := strconv.ParseFloat(stringResults[i][0], 64)
		if err != nil {
			return nil, types.NewError(err.Error())
		}
		returnedResults[i].Lon = tmp

		tmp, err = strconv.ParseFloat(stringResults[i][1], 64)
		if err != nil {
			return nil, types.NewError(err.Error())
		}
		returnedResults[i].Lat = tmp

		returnedResults[i].Name = members[i]
	}

	return returnedResults, nil
}
