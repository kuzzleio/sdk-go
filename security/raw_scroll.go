package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"errors"
)

func (s *Security) rawScroll(action string, scrollId string, options types.QueryOptions) ([]byte, error) {
	if scrollId == "" {
		return nil, errors.New("Security." + action + " : scroll id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		ScrollId:   scrollId,
	}

	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
