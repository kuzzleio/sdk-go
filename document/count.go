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
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// Count returns the number of documents matching the provided set of filters.
// There is a small delay between documents creation and their existence in our advanced search layer,
// usually a couple of seconds.
// That means that a document that was just been created won’t be returned by this function
func (d *Document) Count(index string, collection string, body json.RawMessage, options types.QueryOptions) (int, error) {
	if index == "" {
		return 0, types.NewError("Document.Count: index required", 400)
	}

	if collection == "" {
		return 0, types.NewError("Document.Count: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "count",
		Body:       body,
	}
	if options != nil {
		query.IncludeTrash = options.IncludeTrash()
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return 0, res.Error
	}

	type countResult struct {
		Count int `json:"count"`
	}
	var count countResult
	err := json.Unmarshal(res.Result, &count)

	if err != nil {
		return 0, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return count.Count, nil
}
