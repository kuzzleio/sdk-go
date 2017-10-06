package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ValidateCredentials validates credentials of the specified strategy for the given user.
func (s Security) ValidateCredentials(strategy string, kuid string, credentials interface{}, options types.QueryOptions) (bool, error) {
	if strategy == "" {
		return false, errors.New("Security.ValidateCredentials: strategy is required")
	}

	if kuid == "" {
		return false, errors.New("Security.ValidateCredentials: kuid is required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "validateCredentials",
		Body:       credentials,
		Strategy:   strategy,
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, errors.New(res.Error.Message)
	}

	var hasCredentials bool
	json.Unmarshal(res.Result, &hasCredentials)

	return true, nil
}
