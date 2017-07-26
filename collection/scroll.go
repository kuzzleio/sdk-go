package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  A "scroll" option can be passed to search queries, creating persistent paginated results.
*/
func (dc Collection) Scroll(scrollId string, options types.QueryOptions) (types.KuzzleSearchResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchResult{}, errors.New("Collection.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "document",
		Action:     "scroll",
		ScrollId:   scrollId,
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
