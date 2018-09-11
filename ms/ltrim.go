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

// Ltrim trims an existing list so that it will
// contain only the specified range of elements specified.
func (ms *Ms) Ltrim(key string, start int, stop int, options types.QueryOptions) error {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Start int `json:"start"`
		Stop  int `json:"stop"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "ltrim",
		Id:         key,
		Body:       &body{Start: start, Stop: stop},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}

	return nil
}
