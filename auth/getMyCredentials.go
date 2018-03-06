package auth

import (
	"github.com/kuzzleio/sdk-go/types"
)

// GetMyCredentials get credential information of the specified strategy for the current user.
func (a *Auth) GetMyCredentials(strategy string, options types.QueryOptions) (types.Credentials, error) {
	if strategy == "" {
		return nil, types.NewError("Kuzzle.GetMyCredentials: strategy is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "getMyCredentials",
		Strategy:   strategy,
	}

	go a.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
