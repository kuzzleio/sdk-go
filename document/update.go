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

// Update updates a document in Kuzzle.
func (d *Document) Update(index string, collection string, id string, body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if id == "" {
		return nil, types.NewError("Document.update: id required", 400)
	}

	if index == "" {
		return nil, types.NewError("Document.Update: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.Update: collection required", 400)
	}

	if body == nil {
		return nil, types.NewError("Document.Update: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "update",
		Id:         id,
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return res.Result, nil
}
