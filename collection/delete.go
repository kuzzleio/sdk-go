package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Delete deletes a collection
func (dc *Collection) Delete(index string, collection string) error {
	if index == "" {
		return types.NewError("Collection.Delete: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.Delete: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "delete",
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
