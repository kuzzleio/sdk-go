package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Search documents in the given Collection, using provided filters and option.
func (dc *Collection) Search(filters *types.SearchFilters, options types.QueryOptions) (*SearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "search",
		Body:       filters,
	}

	if options != nil {
		query.From = options.From()
		query.Size = options.Size()

		scroll := options.Scroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	searchResult := &SearchResult{
		Collection: dc,
		Options:    options,
		Filters:    filters,
	}
	json.Unmarshal(res.Result, searchResult)

	for _, d := range searchResult.Hits {
		d.collection = dc
	}

	return searchResult, nil
}
