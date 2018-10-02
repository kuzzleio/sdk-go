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

// MCreate creates the provided documents.
func (d *Document) MCreate(index string, collection string, documents json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Document.MCreate: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.MCreate: collection required", 400)
	}

	if documents == nil {
		return nil, types.NewError("Document.MCreate: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Documents json.RawMessage `json:"documents"`
	}

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mCreate",
		Body:       &body{documents},
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	type r struct {
		Hits json.RawMessage `json:"hits"`
	}
	var docs r
	json.Unmarshal(res.Result, &docs)

	return docs.Hits, nil
}
