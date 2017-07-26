package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Get Kuzzle usage statistics
 */
func (k Kuzzle) GetStatistics(options types.QueryOptions) (types.Statistics, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "server",
		Action:     "getLastStats",
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return types.Statistics{}, errors.New(res.Error.Message)
	}

	s := types.Statistics{}
	json.Unmarshal(res.Result, &s)

	return s, nil
}
