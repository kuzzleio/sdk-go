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

	result := KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
