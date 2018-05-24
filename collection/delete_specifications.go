// Copyright 2015-2017 Kuzzle
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

package collection

import "github.com/kuzzleio/sdk-go/types"

// DeleteSpecifications deletes the current specifications of this collection.
func (dc *Collection) DeleteSpecifications(index string, collection string, options types.QueryOptions) error {
	if index == "" {
		return types.NewError("Collection.DeleteSpecifications: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.DeleteSpecifications: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "deleteSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return res.Error
	}

	return nil
}
