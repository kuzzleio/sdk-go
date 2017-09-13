package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Create a new empty data index, with no associated mapping.
func (k Kuzzle) CreateIndex(index string, options types.QueryOptions) (types.AckResponse, error) {
	if index == "" {
		return types.AckResponse{}, errors.New("Kuzzle.createIndex: index required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "create",
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return types.AckResponse{}, errors.New(res.Error.Message)
	}

	ack := types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
