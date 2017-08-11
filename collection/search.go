package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type KuzzleSearchResult struct {
Hits     []Document `json:"hits"`
Total    int        `json:"total"`
ScrollId string     `json:"_scroll_id"`
}

/*
  Searches documents in the given Collection, using provided filters and option.
*/
func (dc Collection) Search(filters interface{}, options types.QueryOptions) (KuzzleSearchResult, error) {
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

	result := KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
