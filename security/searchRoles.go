package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchRoles(body json.RawMessage, options types.QueryOptions) (*RoleSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchRoles",
		Body:       body,
	}

	if options != nil {
		query.From = options.From()
		query.Size = options.Size()

		scroll := options.Scroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}
	jsonSearchResult := &jsonRoleSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &RoleSearchResult{
		Total: jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		r := j.jsonRoleToRole()
		searchResult.Hits = append(searchResult.Hits, r)
	}

	return searchResult, nil
}
