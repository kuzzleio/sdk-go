package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Get credential information of the specified strategy for the given user.
 */
func (s Security) HasCredentials(strategy string, kuid string, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, errors.New("Security.HasCredentials: strategy is required")
	}

	if kuid == "" {
		return false, errors.New("Security.HasCredentials: kuid is required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "hasCredentials",
		Strategy:   strategy,
		Id:         kuid,
	}

	go s.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
	}

	var hasCredentials bool
	json.Unmarshal(res.Result, &hasCredentials)

	return true, nil
}