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

// Lset sets the list element at index with the provided value.
func (ms *Ms) Lset(key string, index int, value string, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
		Index int    `json:"index"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "lset",
		Id:         key,
		Body:       &body{Value: value, Index: index},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
