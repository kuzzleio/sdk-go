package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetCredentialFields gets an array of strategy's fieldnames
func (s Security) GetCredentialFields(strategy string, options types.QueryOptions) (types.CredentialStrategyFields, error) {
	if strategy == "" {
		return types.CredentialStrategyFields{}, errors.New("Security.GetCredentialFields: strategy is required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getCredentialFields",
		Strategy:   strategy,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return types.CredentialStrategyFields{}, errors.New(res.Error.Message)
	}

	credentialFields := types.CredentialStrategyFields{}
	json.Unmarshal(res.Result, &credentialFields)

	return credentialFields, nil
}
