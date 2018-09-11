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
	"github.com/kuzzleio/sdk-go/types"
)

// Georadiusbymember returns the geospatial members of a key inside the provided radius
func (ms *Ms) Georadiusbymember(key string, member string, distance float64, unit string, options types.QueryOptions) ([]*types.Georadius, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return responseToGeoradius(res, options)
}
