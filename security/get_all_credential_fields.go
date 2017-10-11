package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetAllCredentialFields gets an array of strategy's fieldnames for each strategies
func (s Security) GetAllCredentialFields(options types.QueryOptions) (types.CredentialFields, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getAllCredentialFields",
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return types.CredentialFields{}, res.Error
	}

	credentialFields := types.CredentialFields{}
	json.Unmarshal(res.Result, &credentialFields)

	return credentialFields, nil
}
