package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Truncate delete every Documents from the provided Collection.
func (dc *Collection) Truncate(index string, collection string) error {
	if index == "" {
		return types.NewError("Collection.Truncate: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.Truncate: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "truncate",
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
