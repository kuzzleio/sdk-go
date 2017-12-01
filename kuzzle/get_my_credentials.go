package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetMyCredentials get credential information of the specified strategy for the current user.
func (k *Kuzzle) GetMyCredentials(strategy string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.GetMyCredentials: strategy is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyCredentials",
		Strategy:   strategy,
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
