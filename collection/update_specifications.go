package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// UpdateSpecifications updates the current specifications of this collection.
func (dc *Collection) UpdateSpecifications(index string, collection string, body json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Collection.UpdateSpecifications: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Collection.UpdateSpecifications: collection required", 400)
	}

	if body == nil {
		return nil, types.NewError("Collection.UpdateSpecifications: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "updateSpecifications",
		Body:       body,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
