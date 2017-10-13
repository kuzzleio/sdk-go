package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateSelf update the currently authenticated user information.
func (k Kuzzle) UpdateSelf(credentials interface{}, options types.QueryOptions) (*types.User, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "updateSelf",
		Body:       credentials,
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	u := &types.User{}
	json.Unmarshal(res.Result, u)

	return u, nil
}
