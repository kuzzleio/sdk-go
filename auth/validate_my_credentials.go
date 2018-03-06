package auth

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ValidateMyCredentials validate credentials of the specified strategy for the current user.
func (a *Auth) ValidateMyCredentials(strategy string, credentials interface{}, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "validateMyCredentials",
		Strategy:   strategy,
		Body:       credentials,
	}

	go a.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var r bool
	json.Unmarshal(res.Result, &r)

	return r, nil
}
