package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// CreateMyCredentials create credentials of the specified strategy for the current user.
func (a *Auth) CreateMyCredentials(strategy string, credentials json.RawMessage, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.CreateMyCredentials: strategy is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "createMyCredentials",
		Body:       credentials,
		Strategy:   strategy,
	}
	go a.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
