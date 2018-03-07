package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateMapping updates the current mapping of this collection.
func (dc *Collection) UpdateMapping(index string, collection string, body string) error {
	if index == "" {
		return types.NewError("Collection.UpdateMapping: index required", 400)
	}

	if collection == "" {
		return types.NewError("Collection.UpdateMapping: collection required", 400)
	}

	if body == "" {
		return types.NewError("Collection.UpdateMapping: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "updateMapping",
		Body:       body,
	}
	go dc.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return res.Error
	}

	return nil
}
