package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// SearchSpecifications searches specifications across indexes/collections according to the provided filters.
func (dc *Collection) SearchSpecifications(options types.QueryOptions) (*types.SearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specifications := &types.SearchResult{}
	json.Unmarshal(res.Result, &specifications)

	return specifications, nil
}
