package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

type SearchResult struct {
	Collection   *Collection `json:"-"`
	Documents    []*Document `json:"hits"`
	Total        int         `json:"total"`
	Fetched      int
	ScrollId     string                 `json:"_scroll_id"`
	Aggregations map[string]interface{} `json:"aggregations"`
	Options      types.QueryOptions
	Filters      *types.SearchFilters
	Previous     *SearchResult
}

// FetchNext returns a new SearchResult that corresponds to the next result page
func (ksr *SearchResult) FetchNext() (*SearchResult, error) {
	if ksr.Fetched >= ksr.Total {
		return nil, nil
	}

	if ksr.ScrollId != "" {
		if ksr.Options != nil {
			ksr.Options.SetFrom(0)
			ksr.Options.SetSize(0)
		}

		return ksr.Collection.scrollFrom(ksr, ksr.Options)
	}

	if ksr.Options != nil && ksr.Filters != nil {
		if ksr.Options.Size() > 0 && len(ksr.Filters.Sort) > 0 {
			var filters = ksr.Filters
			var source = ksr.Documents[len(ksr.Documents)-1].SourceToMap()

			for _, sortRules := range filters.Sort {
				switch t := sortRules.(type) {
				case string:
					filters.SearchAfter = append(filters.SearchAfter, source[t])
				case map[string]interface{}:
					for field := range t {
						filters.SearchAfter = append(filters.SearchAfter, source[field])
					}
				}
			}

			ksr.Options.SetFrom(0)

			return ksr.Collection.Search(filters, ksr.Options)
		}

		if ksr.Options.Size() > 0 {
			ksr.Options.SetFrom(ksr.Options.From() + ksr.Options.Size())

			if ksr.Options.From() >= ksr.Total {
				return nil, nil
			}

			return ksr.Collection.Search(ksr.Filters, ksr.Options)
		}
	}

	return nil, types.NewError("SearchResult.FetchNext: Unable to retrieve results: missing scrollId or from/size parameters", 400)
}
