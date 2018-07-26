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

func (s *Security) SearchProfiles(body json.RawMessage, options types.QueryOptions) (*ProfileSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchProfiles",
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

	jsonSearchResult := &jsonProfileSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &ProfileSearchResult{
		ScrollId: jsonSearchResult.ScrollId,
		Total:    jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		p := j.jsonProfileToProfile()
		searchResult.Hits = append(searchResult.Hits, p)
	}

	return searchResult, nil
}
