package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetCredentialsByID recover credentials from given strategy identified by given id
func (s *Security) GetCredentialsByID(strategy, id string, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" || id == "" {
		return nil, types.NewError("Security.GetCredentialById: strategy and id are required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getCredentialsById",
		Strategy:   strategy,
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
