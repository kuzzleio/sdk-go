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

// Expire sets an expiration timeout on a key
func (ms *Ms) Expire(key string, seconds int, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Seconds int `json:"seconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "expire",
		Id:         key,
		Body:       &body{Seconds: seconds},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
