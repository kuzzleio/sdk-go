package collection

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// Exists check if a collection exists.
func (dc *Collection) Exists(index string, collection string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Collection.Exists: index required", 400)
	}

	if collection == "" {
		return false, types.NewError("Collection.Exists: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "exists",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	var exists bool

	err := json.Unmarshal(res.Result, &exists)

	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return exists, nil
}
