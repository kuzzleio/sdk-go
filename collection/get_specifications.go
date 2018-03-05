package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// GetSpecifications retrieves the current specifications of the collection.
func (dc *Collection) GetSpecifications(index string, collection string) (string, error) {
	if index == "" {
		return "", types.NewError("Collection.GetSpecifications: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Collection.GetSpecifications: collection required", 400)
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
		return "", res.Error
	}

	var specifications string
	specifications = string(res.Result)

	return specifications, nil
}
