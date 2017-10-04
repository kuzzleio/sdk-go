package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ValidateMyCredentials validate credentials of the specified strategy for the current user.
func (k Kuzzle) ValidateMyCredentials(strategy string, credentials interface{}, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "auth",
		Action:     "validateMyCredentials",
		Strategy:   strategy,
		Body:       credentials,
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
	}

	var r bool
	json.Unmarshal(res.Result, &r)

	return r, nil
}
