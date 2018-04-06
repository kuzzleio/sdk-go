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

// DeleteByQuery deletes all the documents from Kuzzle that match the given filter or query.
func (d *Document) DeleteByQuery(index string, collection string, body string, options types.QueryOptions) ([]string, error) {
	if index == "" {
		return nil, types.NewError("Document.MCreate: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.MCreate: collection required", 400)
	}

	if body == "" {
		return nil, types.NewError("Document.MCreate: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "deleteByQuery",
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var deleted []string
	json.Unmarshal(res.Result, &deleted)

	return deleted, nil

}
