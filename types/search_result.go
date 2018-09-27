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

package types

import (
	"encoding/json"
	"fmt"
)

// SearchResult is a search result representation
type SearchResult struct {
	Collection   string
	Documents    json.RawMessage
	Total        int
	Fetched      int
	Aggregations json.RawMessage
	Options      QueryOptions
	Filters      json.RawMessage
}

// NewSearchResult Search Result constructor
func NewSearchResult(collection string, filters json.RawMessage, options QueryOptions, raw *KuzzleResponse) (*SearchResult, error) {
	type ParseSearchResult struct {
		Documents    json.RawMessage `json:"hits"`
		Total        int             `json:"total"`
		ScrollID     string          `json:"_scroll_id"`
		Aggregations json.RawMessage `json:"aggregations"`
	}

	var parsed ParseSearchResult
	err := json.Unmarshal(raw.Result, &parsed)

	if err != nil {
		return nil, NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), raw.Result), 500)
	}

	sr := &SearchResult{
		Collection:   collection,
		Filters:      filters,
		Documents:    parsed.Documents,
		Total:        parsed.Total,
		Fetched:      len(parsed.Documents),
		Aggregations: parsed.Aggregations,
		Options:      NewQueryOptions(),
	}

	if options != nil {
		sr.Options.SetFrom(options.From())
		sr.Options.SetSize(options.Size())
	}

	return sr, nil
}
