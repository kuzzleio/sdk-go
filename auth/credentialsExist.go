package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//CredentialsExist check the existence of the specified <strategy>'s credentials for the current user.
func (a *Auth) CredentialsExist(strategy string, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, types.NewError("Auth.CredentialsExist: strategy is required", 400)
	}
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "credentialsExist",
		Strategy:   strategy,
	}
	go a.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var exists bool
	json.Unmarshal(res.Result, &exists)

	return exists, nil
}
