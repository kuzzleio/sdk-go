package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) rawFetch(action string, id string, options types.QueryOptions) ([]byte, error) {
	if id == "" {
		return nil, types.NewError("Security."+action+": id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Id:         id,
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
