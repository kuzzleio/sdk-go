package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Get credential information of the specified strategy for the current user.
 */
func (k Kuzzle) GetMyCredentials(strategy string, options *types.Options) (json.RawMessage, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyCredentials",
		Strategy:   strategy,
	}

	go k.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
