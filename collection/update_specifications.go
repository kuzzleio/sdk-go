package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateSpecifications updates the current specifications of this collection.
func (dc *Collection) UpdateSpecifications(index string, collection string, body string) (string, error) {
	if index == "" {
		return "", types.NewError("Collection.UpdateSpecifications: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Collection.UpdateSpecifications: collection required", 400)
	}

	if body == "" {
		return "", types.NewError("Collection.UpdateSpecifications: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "updateSpecifications",
		Body:       body,
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var result string
	result = string(res.Result)

	return result, nil
}
