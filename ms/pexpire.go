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

// Pexpire sets a timeout (in milliseconds) on a key.
// After the timeout has expired, the key will automatically be deleted.
func (ms *Ms) Pexpire(key string, ttl int, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Milliseconds int `json:"milliseconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpire",
		Id:         key,
		Body:       &body{Milliseconds: ttl},
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
