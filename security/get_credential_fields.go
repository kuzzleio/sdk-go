package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetCredentialFields gets an array of strategy's fieldnames
func (s Security) GetCredentialFields(strategy string, options types.QueryOptions) (types.CredentialStrategyFields, error) {
	if strategy == "" {
		return types.CredentialStrategyFields{}, types.NewError("Security.GetCredentialFields: strategy is required")
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
		return types.CredentialStrategyFields{}, res.Error
	}

	credentialFields := types.CredentialStrategyFields{}
	json.Unmarshal(res.Result, &credentialFields)

	return credentialFields, nil
}
