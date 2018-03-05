package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// GetMapping retrieves the current mapping of the collection.
func (dc *Collection) GetMapping(index string, collection string) (string, error) {
	if index == "" {
		return "", types.NewError("Collection.GetMapping: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Collection.GetMapping: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "getMapping",
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var mapping string
	mapping = string(res.Result)

	return mapping, nil
}
