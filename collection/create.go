package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Create creates a new empty data collection
func (dc *Collection) Create(index string, collection string, options types.QueryOptions) error {
	if index == "" {
		return types.NewError("Collection.Create: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.Create: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "create",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	return res.Error
}
