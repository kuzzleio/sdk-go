package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Count returns the number of documents matching the provided set of filters.
// There is a small delay between documents creation and their existence in our advanced search layer,
// usually a couple of seconds.
// That means that a document that was just been created won’t be returned by this function
func (dc *Collection) Count(filters *types.SearchFilters, options types.QueryOptions) (int, error) {
	type countResult struct {
		Count  int `json:"count"`
		Status int
	}

	searchfilters := &types.SearchFilters{}

	if filters != nil {
		searchfilters = filters
	}

	ch := make(chan *types.KuzzleResponse)
	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "count",
		Body:       searchfilters,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return -1, res.Error
	}

	result := &countResult{}
	json.Unmarshal(res.Result, result)

	return result.Count, nil
}
