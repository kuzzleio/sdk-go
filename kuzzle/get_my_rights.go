package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetMyRights gets the rights array for the currently logged user.
func (k Kuzzle) GetMyRights(options types.QueryOptions) ([]types.Rights, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyRights",
	}

	type rights struct {
		Hits []types.Rights `json:"hits"`
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	r := rights{}
	json.Unmarshal(res.Result, &r)

	return r.Hits, nil
}
