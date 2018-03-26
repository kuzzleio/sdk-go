package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateCredentials updates credentials of the specified strategy for the given user.
func (s *Security) UpdateCredentials(strategy string, kuid string, body string, options types.QueryOptions) error {
	if strategy == "" || kuid == "" {
		return types.NewError("Security.UpdateCredentials: strategy and kuid are required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateCredentials",
		Body:       body,
		Strategy:   strategy,
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	return nil
}
