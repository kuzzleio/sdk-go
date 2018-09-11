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

// Linsert inserts a value in a list, either before or after the reference pivot value.
func (ms *Ms) Linsert(key string, position string, pivot string, value string, options types.QueryOptions) (int, error) {
	if position != "before" && position != "after" {
		return -1, types.NewError("Ms.Linsert: invalid position argument (must be 'before' or 'after')", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Position string `json:"position"`
		Pivot    string `json:"pivot"`
		Value    string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "linsert",
		Id:         key,
		Body:       &body{Position: position, Pivot: pivot, Value: value},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return -1, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
