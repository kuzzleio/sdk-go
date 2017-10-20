package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteMyCredentials delete credentials of the specified strategy for the current user.
func (k Kuzzle) DeleteMyCredentials(strategy string, options types.QueryOptions) (*types.AckResponse, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.DeleteMyCredentials: strategy is required", 400)
	}

	type body struct {
		Strategy string `json:"strategy"`
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "deleteMyCredentials",
		Strategy:   strategy,
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
