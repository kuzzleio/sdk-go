package types

import (
	"encoding/json"
)

type SearchResult struct {
	Collection   json.RawMessage
	Documents    json.RawMessage
	Total        int
	Fetched      int
	Aggregations json.RawMessage
	Options      QueryOptions
	Filters      *SearchFilters
}

//NewSearchResult Search Result constructor
func NewSearchResult(collection json.RawMessage, filters *SearchFilters, options QueryOptions, raw *KuzzleResponse) *SearchResult {
	type ParseSearchResult struct {
		Documents    json.RawMessage `json:"hits"`
		Total        int             `json:"total"`
		ScrollID     string          `json:"_scroll_id"`
		Aggregations json.RawMessage `json:"aggregations"`
	}

	var parsed ParseSearchResult
	json.Unmarshal(raw.Result, &parsed)

	sr := &SearchResult{
		Collection:   collection,
		Filters:      filters,
		Documents:    parsed.Documents,
		Total:        parsed.Total,
		Fetched:      len(parsed.Documents),
		Aggregations: parsed.Aggregations,
		Options:      NewQueryOptions(),
	}

	sr.Options.SetScrollId(parsed.ScrollID)

	if options != nil {
		sr.Options.SetFrom(options.From())
		sr.Options.SetSize(options.Size())
	}

	return sr
}
