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

package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Create a new document in Kuzzle.
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//       Additional information passed to notifications to other users
//   - ifExist (string, allowed values: "error" (default), "replace"):
//       If the same document already exists:
//         - resolves with an error if set to "error".
//         - replaces the existing document if set to "replace"
func (d *Document) Create(index string, collection string, _id string, body string, options types.QueryOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.Create: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.Create: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.Create: id required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.Create: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "create",
		Id:         _id,
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var created string
	json.Unmarshal(res.Result, &created)

	return created, nil
}
