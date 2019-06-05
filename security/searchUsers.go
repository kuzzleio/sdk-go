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

package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// SearchUsers retrieves the users matching the given query
func (s *Security) SearchUsers(body json.RawMessage, options types.QueryOptions) (*UserSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
		Body:       body,
	}

	if options != nil {
		query.From = options.From()
		query.Size = options.Size()

		scroll := options.Scroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	jsonSearchResult := &jsonUserSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	sr, err := types.NewSearchResult(s.Kuzzle, "scrollUsers", query, options, res)

	if err != nil {
		return nil, err
	}

	searchResult := &UserSearchResult{
		Aggregations: sr.Aggregations,
		Total:        sr.Total,
		Fetched:      sr.Fetched,
		ScrollId:     sr.ScrollId,
		kuzzle:       s.Kuzzle,
		request:      query,
		response:     res,
		options:      options,
	}

	for _, j := range jsonSearchResult.Hits {
		searchResult.Hits = append(searchResult.Hits, j.jsonUserToUser())
	}

	return searchResult, nil
}
