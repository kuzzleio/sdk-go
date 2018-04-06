package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// List lists all subscriptions on all indexes and all collections.
func (r *Realtime) List(index string, collection string) (json.RawMessage, error) {
	if index == "" || collection == "" {
		return nil, types.NewError("Realtime.List: index and collection required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "list",
		Index:      index,
		Collection: collection,
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
