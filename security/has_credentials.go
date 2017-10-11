package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

// HasCredentials gets credential information of the specified strategy for the given user.
func (s Security) HasCredentials(strategy string, kuid string, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, types.NewError("Security.HasCredentials: strategy is required")
	}

	if kuid == "" {
		return false, types.NewError("Security.HasCredentials: kuid is required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "hasCredentials",
		Strategy:   strategy,
		Id:         kuid,
	}

	go s.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var r bool
	json.Unmarshal(res.Result, &r)

	return r, nil
}