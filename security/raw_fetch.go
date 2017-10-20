package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"errors"
)

func (s *Security) rawFetch(action string, id string, options types.QueryOptions) ([]byte, error) {
	if id == "" {
		return nil, errors.New("Security." + action + ": id is required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action: action,
		Id: id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
