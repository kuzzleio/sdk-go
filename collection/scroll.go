package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Scroll passes a "scroll" option to search queries, creating persistent paginated results.
func (dc *Collection) Scroll(scrollId string, options types.QueryOptions) (*SearchResult, error) {
	if scrollId == "" {
		return &SearchResult{}, types.NewError("Collection.Scroll: scroll id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "document",
		Action:     "scroll",
		ScrollId:   scrollId,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return &SearchResult{}, res.Error
	}

	searchResult := &SearchResult{
		Collection: dc,
		Options:    options,
	}
	json.Unmarshal(res.Result, searchResult)

	for _, d := range searchResult.Hits {
		d.collection = dc
	}

	return searchResult, nil
}
