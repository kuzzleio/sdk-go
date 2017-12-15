package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Scroll passes a "scroll" option to search queries, creating persistent paginated results.
func (dc *Collection) Scroll(scrollId string, options types.QueryOptions) (*SearchResult, error) {
	if scrollId == "" {
		return nil, types.NewError("Collection.Scroll: scroll id required", 400)
	}

	return dc.scrollFrom(scrollId, options)
}

// Non-documented function: scrolls from previous search resultsz
func (dc *Collection) scrollFrom(from interface{}, options types.QueryOptions) (*SearchResult, error) {
	var scrollId string
	var previous *SearchResult
	var fetched int

	switch t := from.(type) {
	case string:
		scrollId = t
		fetched = 0
	case *SearchResult:
		previous = t
		scrollId = t.ScrollId
		fetched = t.Fetched
	default:
		return nil, nil
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

	searchResult := &SearchResult{
		Collection: dc,
		Previous:   previous,
		Fetched:    fetched,
		Options:    options,
	}
	json.Unmarshal(res.Result, searchResult)

	for _, d := range searchResult.Documents {
		d.collection = dc
	}

	searchResult.Fetched += len(searchResult.Documents)

	return searchResult, nil
}
