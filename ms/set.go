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

// Set creates a key holding the provided value, or overwrites it if it already exists.
func (ms *Ms) Set(key string, value interface{}, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ex    int         `json:"ex,omitempty"`
		Px    int         `json:"px,omitempty"`
		Nx    bool        `json:"nx"`
		Xx    bool        `json:"xx"`
	}

	bodyContent := body{Value: value}

	if options != nil {
		if options.Ex() != 0 {
			bodyContent.Ex = options.Ex()
		}

		if options.Px() != 0 {
			bodyContent.Px = options.Px()
		}

		bodyContent.Nx = options.Nx()
		bodyContent.Xx = options.Xx()
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "set",
		Id:         key,
		Body:       &bodyContent,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
