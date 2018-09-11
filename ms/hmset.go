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

// Hmset sets multiple fields at once in a hash.
func (ms *Ms) Hmset(key string, entries []*types.MsHashField, options types.QueryOptions) error {
	if len(entries) == 0 {
		return types.NewError("Ms.Hmset: at least one entry field to set is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Entries []*types.MsHashField `json:"entries"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hmset",
		Id:         key,
		Body:       &body{Entries: entries},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
