package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  A "scroll" option can be passed to search queries, creating persistent paginated results.
*/
func (dc Collection) Scroll(scrollId string, options types.QueryOptions) (KuzzleSearchResult, error) {
	if scrollId == "" {
		return KuzzleSearchResult{}, errors.New("Collection.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "document",
		Action:     "scroll",
		ScrollId:   scrollId,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := KuzzleSearchResult{
		Collection: dc,
		Options:    options,
	}
	json.Unmarshal(res.Result, &searchResult)

	for _, d := range searchResult.Hits {
		d.collection = dc
	}

	return searchResult, nil
}
