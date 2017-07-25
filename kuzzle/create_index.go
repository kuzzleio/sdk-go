package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

func (k Kuzzle) CreateIndex(index string, options *types.Options) (*types.AckResponse, error) {
	if index == "" {
		return nil, errors.New("Kuzzle.createIndex: index required")
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
		return nil, errors.New(res.Error.Message)
	}

	ack := &types.AckResponse{}
	json.Unmarshal(res.Result, &ack)

	return ack, nil
}
