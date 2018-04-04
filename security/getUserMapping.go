package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// GetUserMapping gets mapping for Users
func (s *Security) GetUserMapping(options types.QueryOptions) (json.RawMessage, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserMapping",
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
