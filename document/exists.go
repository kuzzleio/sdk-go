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

package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Exists returns a boolean indicating whether or not a document with provided ID exists.
func (d *Document) Exists(index string, collection string, _id string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Document.Exists: index required", 400)
	}

	if collection == "" {
		return false, types.NewError("Document.Exists: collection required", 400)
	}

	if _id == "" {
		return false, types.NewError("Document.Exists: id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "exists",
		Id:         _id,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var exists bool
	json.Unmarshal(res.Result, &exists)

	return exists, nil
}
