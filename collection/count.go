package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Count returns the number of documents matching the provided set of filters.
// There is a small delay between documents creation and their existence in our advanced search layer,
// usually a couple of seconds.
// That means that a document that was just been created won’t be returned by this function
func (dc Collection) Count(filters *types.SearchFilters, options types.QueryOptions) (int, error) {
	type countResult struct {
		Count int `json:"count"`
	}

	ch := make(chan *types.KuzzleResponse)
	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "count",
		Body:       filters,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return 0, errors.New(res.Error.Message)
	}

	result := &countResult{}
	json.Unmarshal(res.Result, result)

	return result.Count, nil
}
