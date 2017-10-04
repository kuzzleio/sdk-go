package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetCredentials gets credential information of the specified strategy for the given user.
func (s Security) GetCredentials(strategy string, kuid string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, errors.New("Security.GetCredentials: strategy is required")
	}

	if kuid == "" {
		return nil, errors.New("Security.GetCredentials: kuid is required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getCredentials",
		Strategy:   strategy,
		Id:         kuid,
	}

	go s.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
