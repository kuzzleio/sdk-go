package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) rawSearch(action string, filters interface{}, options types.QueryOptions) ([]byte, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
