package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * Get an array of strategy's fieldnames
 */
func (s Security) GetCredentialFields(strategy string, options *types.Options) ([]string, error) {
	if strategy == "" {
		return make([]string, 0), errors.New("Security.GetCredentialFields: strategy is required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getCredentialFields",
		Strategy:   strategy,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return make([]string, 0), errors.New(res.Error.Message)
	}

	credentialFields := make([]string, 0)
	json.Unmarshal(res.Result, &credentialFields)

	return credentialFields, nil
}
