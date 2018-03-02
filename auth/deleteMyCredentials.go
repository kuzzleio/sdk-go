package auth

import (
	"github.com/kuzzleio/sdk-go/types"
)

// DeleteMyCredentials delete credentials of the specified strategy for the current user.
func (a *Auth) DeleteMyCredentials(strategy string, options types.QueryOptions) error {
	if strategy == "" {
		return types.NewError("Auth.DeleteMyCredentials: strategy is required", 400)
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

	go a.k.Query(query, options, result)

	res := <-result

	return res.Error
}
