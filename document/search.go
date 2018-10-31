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

// Search documents in the given Collection, using provided filters and option.
func (d *Document) Search(index string, collection string, body json.RawMessage, options types.QueryOptions) (*types.SearchResult, error) {
	if index == "" {
		return nil, types.NewError("Document.Search: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.Search: collection required", 400)
	}

	if body == nil {
		return nil, types.NewError("Document.Search: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "search",
		Body:       body,
	}
	if options != nil {
		query.IncludeTrash = options.IncludeTrash()
		query.From = options.From()
		query.Size = options.Size()
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	sr, err := types.NewSearchResult(d.Kuzzle, "scroll", query, options, res)

	if err != nil {
		return nil, err
	}

	return sr, nil
}
