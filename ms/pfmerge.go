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

// Pfmerge merges multiple HyperLogLog data structures into an unique HyperLogLog
// structure stored at key, approximating the cardinality of the union of the source structures.
func (ms *Ms) Pfmerge(key string, sources []string, options types.QueryOptions) error {
	if len(sources) == 0 {
		return types.NewError("Ms.Pfmerge: please provide at least one source to merge", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Sources []string `json:"sources"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pfmerge",
		Id:         key,
		Body:       &body{Sources: sources},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
