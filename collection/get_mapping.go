package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetMapping retrieves the current mapping of the collection.
func (dc *Collection) GetMapping(index string, collection string, options types.QueryOptions) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Collection.GetMapping: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Collection.GetMapping: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "getMapping",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
