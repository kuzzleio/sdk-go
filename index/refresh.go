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

package index

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Refresh forces the provided data index to refresh on each modification
func (i *Index) Refresh(index string, options types.QueryOptions) error {
	if index == "" {
		return types.NewError("Index.Refresh: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refresh",
		Index:      index,
	}
	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}
	return nil
}
