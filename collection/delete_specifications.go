package collection

import "github.com/kuzzleio/sdk-go/types"

// DeleteSpecifications deletes the current specifications of this collection.
func (dc *Collection) DeleteSpecifications(index string, collection string, options types.QueryOptions) error {
	if index == "" {
		return types.NewError("Collection.DeleteSpecifications: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.DeleteSpecifications: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "deleteSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
