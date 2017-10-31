package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteMyCredentials delete credentials of the specified strategy for the current user.
func (k Kuzzle) DeleteMyCredentials(strategy string, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, types.NewError("Kuzzle.DeleteMyCredentials: strategy is required", 400)
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
		return false, res.Error
	}

	return true, nil
}
