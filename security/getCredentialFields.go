package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetCredentialFields gets an array of strategy's fieldnames
func (s *Security) GetCredentialFields(strategy string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, types.NewError("Security.GetCredentialFields: strategy is required", 400)
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
		return nil, res.Error
	}

	return res.Result, nil
}
