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

package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// SearchSpecifications searches specifications across indexes/collections according to the provided filters.
func (dc *Collection) SearchSpecifications(body json.RawMessage, options types.QueryOptions) (*types.SearchResult, error) {
	if body == nil {
		return nil, types.NewError("Document.Search: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
		Body:       body,
	}
	if options != nil {
		query.From = options.From()
		query.Size = options.Size()
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	sr, err := types.NewSearchResult(dc.Kuzzle, "scrollSpecifications", query, options, res)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
