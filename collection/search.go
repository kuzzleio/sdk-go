package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Searches documents in the given Collection, using provided filters and option.
*/
func (dc Collection) Search(filters types.SearchFilters, options types.QueryOptions) (KuzzleSearchResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "search",
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := KuzzleSearchResult{
		Collection: dc,
		Options:    options,
		Filters:    filters,
	}
	json.Unmarshal(res.Result, &searchResult)

	for _, d := range searchResult.Hits {
		d.collection = dc
	}

	return searchResult, nil
}
