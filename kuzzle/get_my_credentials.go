package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetMyCredentials get credential information of the specified strategy for the current user.
func (k Kuzzle) GetMyCredentials(strategy string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, errors.New("Kuzzle.GetMyCredentials: strategy is required")
	}
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyCredentials",
		Strategy:   strategy,
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
