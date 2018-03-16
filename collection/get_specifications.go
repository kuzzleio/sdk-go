package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetSpecifications retrieves the current specifications of the collection.
func (dc *Collection) GetSpecifications(index string, collection string) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Collection.GetSpecifications: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Collection.GetSpecifications: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "getSpecifications",
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
