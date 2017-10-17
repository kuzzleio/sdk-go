package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"errors"
	"encoding/json"
)

func (s *Security) SearchRoles(filters interface{}, options types.QueryOptions) (*RoleSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
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
		return nil, errors.New(res.Error.Message)
	}

	jsonSearchResult := &jsonSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &RoleSearchResult{
		Total: jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		r := s.jsonRoleToRole(j)
		r.Kuzzle = s.Kuzzle
		searchResult.Hits = append(searchResult.Hits, r)
	}

	return searchResult, nil
}
