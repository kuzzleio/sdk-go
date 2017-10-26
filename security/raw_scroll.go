package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) rawScroll(action string, scrollId string, options types.QueryOptions) ([]byte, error) {
	if scrollId == "" {
		return nil, types.NewError("Security."+action+": scroll id is required", 400)
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
		return nil, res.Error
	}

	return res.Result, nil
}
