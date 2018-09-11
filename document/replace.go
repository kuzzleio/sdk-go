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

// Replace replaces a document in Kuzzle.
func (d *Document) Replace(index string, collection string, _id string, body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Document.Replace: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.Replace: collection required", 400)
	}

	if _id == "" {
		return nil, types.NewError("Document.Replace: id required", 400)
	}

	if body == nil {
		return nil, types.NewError("Document.Replace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "replace",
		Body:       body,
		Id:         _id,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	return res.Result, nil
}
