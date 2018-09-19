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

// PexipreAt sets an expiration timestamp on a key.
// After the timestamp has been reached, the key will automatically be deleted.
// The timestamp parameter accepts an Epoch time value, in milliseconds.
func (ms *Ms) Pexpireat(key string, timestamp uint64, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Timestamp uint64 `json:"timestamp"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpireat",
		Id:         key,
		Body:       &body{Timestamp: timestamp},
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
