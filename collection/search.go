package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Searches documents in the given Collection, using provided filters and options.
*/
func (dc Collection) Search(filters interface{}, options *types.Options) (types.KuzzleSearchResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "search",
		Body:       filters,
	}

	if options != nil {
		query.From = options.From
		query.Size = options.Size
		if options.Scroll != "" {
			query.Scroll = options.Scroll
		}
	}

	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}
