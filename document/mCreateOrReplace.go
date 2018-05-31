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

// MCreateOrReplace creates or replaces the provided documents.
func (d *Document) MCreateOrReplace(index string, collection string, body json.RawMessage, options types.QueryOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.MCreateOrReplace: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.MCreateOrReplace: collection required", 400)
	}

	if body == nil {
		return "", types.NewError("Document.MCreateOrReplace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mCreateOrReplace",
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return "", res.Error
	}

	var mCreated string
	json.Unmarshal(res.Result, &mCreated)

	return mCreated, nil
}
