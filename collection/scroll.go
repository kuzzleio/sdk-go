package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Scroll passes a "scroll" option to search queries, creating persistent paginated results.
func (dc *Collection) Scroll(scrollId string, options types.QueryOptions) (*SearchResult, error) {
	if scrollId == "" {
		return nil, types.NewError("Collection.Scroll: scroll id required", 400)
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
		return nil, res.Error
	}

	return NewSearchResult(dc, nil, options, res), nil
}
