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

// SetEx sets a key with the provided value, and an expiration delay expressed in seconds.
// If the key does not exist, it is created beforehand.
func (ms *Ms) Setex(key string, value interface{}, ttl int, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ttk   int         `json:"seconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "setex",
		Id:         key,
		Body:       &body{Value: value, Ttk: ttl},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
