package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

type SearchResult struct {
	Collection   *Collection
	Documents    []*Document
	Total        int
	Fetched      int
	Aggregations map[string]interface{}
	Options      types.QueryOptions
	Filters      *types.SearchFilters
}

func NewSearchResult(collection *Collection, filters *types.SearchFilters, options types.QueryOptions, raw *types.KuzzleResponse) *SearchResult {
	type ParseSearchResult struct {
		Documents    []*Document            `json:"hits"`
		Total        int                    `json:"total"`
		ScrollId     string                 `json:"_scroll_id"`
		Aggregations map[string]interface{} `json:"aggregations"`
	}

	var parsed ParseSearchResult
	json.Unmarshal(raw.Result, &parsed)

	for _, d := range parsed.Documents {
		d.collection = collection
	}

	sr := &SearchResult{
		Collection:   collection,
		Filters:      filters,
		Documents:    parsed.Documents,
		Total:        parsed.Total,
		Fetched:      len(parsed.Documents),
		Aggregations: parsed.Aggregations,
		Options:      types.NewQueryOptions(),
	}

	sr.Options.SetScrollId(parsed.ScrollId)

	if options != nil {
		sr.Options.SetFrom(options.From())
		sr.Options.SetSize(options.Size())
	} else {
		sr.Options.SetFrom(0)
		sr.Options.SetSize(0)
	}

	return sr
}

// FetchNext returns a new SearchResult that corresponds to the next result page
func (ksr *SearchResult) FetchNext() (*SearchResult, error) {
	if ksr.Fetched >= ksr.Total {
		return nil, nil
	}

	if ksr.Options.ScrollId() != "" {
		res, err := ksr.Collection.Scroll(ksr.Options.ScrollId(), nil)
		return ksr.afterFetch(res, err)
	}

	if ksr.Options.Size() > 0 {
		if ksr.Filters != nil && len(ksr.Filters.Sort) > 0 {
			source := ksr.Documents[len(ksr.Documents)-1].SourceToMap()

			filters := &types.SearchFilters{
				Query:        ksr.Filters.Query,
				Sort:         ksr.Filters.Sort,
				Aggregations: ksr.Filters.Aggregations,
			}

			for _, sortRules := range ksr.Filters.Sort {
				switch t := sortRules.(type) {
				case string:
					filters.SearchAfter = append(filters.SearchAfter, source[t])
				case map[string]interface{}:
					for field := range t {
						filters.SearchAfter = append(filters.SearchAfter, source[field])
					}
				}
			}

			res, err := ksr.Collection.Search(filters, ksr.Options)
			return ksr.afterFetch(res, err)
		} else {
			opts := types.NewQueryOptions()
			opts.SetFrom(ksr.Options.From() + ksr.Options.Size())

			if opts.From() >= ksr.Total {
				return nil, nil
			}

			opts.SetSize(ksr.Options.Size())

			res, err := ksr.Collection.Search(ksr.Filters, opts)
			return ksr.afterFetch(res, err)
		}
	}

	return nil, types.NewError("SearchResult.FetchNext: Unable to retrieve results: missing scrollId or from/size parameters", 400)
}

func (ksr *SearchResult) afterFetch(nextResult *SearchResult, err error) (*SearchResult, error) {
	if err != nil {
		return nextResult, err
	}

	nextResult.Fetched = len(nextResult.Documents) + ksr.Fetched

	return nextResult, nil
}
