package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// CreateIndex create a new empty data index, with no associated mapping.
func (k Kuzzle) CreateIndex(index string, options types.QueryOptions) (*types.AckResponse, error) {
	if index == "" {
		return nil, types.NewError("Kuzzle.createIndex: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "create",
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, ack)

	return ack, nil
}
