package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// List retrieves the list of known data collections contained in a specified index.
func (dc *Collection) List(index string, options types.QueryOptions) (json.RawMessage, error) {
	if index == "" {
		return nil, types.NewError("Collection.List: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "list",
		Index:      index,
	}

	go dc.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
