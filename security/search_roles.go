package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchRoles(filters interface{}, options types.QueryOptions) (*RoleSearchResult, error) {
	res, err := s.rawSearch("searchRoles", filters, options)

	if err != nil {
		return nil, err
	}

	jsonSearchResult := &jsonRoleSearchResult{}
	json.Unmarshal(res, jsonSearchResult)

	searchResult := &RoleSearchResult{
		Total: jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		r := s.jsonRoleToRole(j)
		searchResult.Hits = append(searchResult.Hits, r)
	}

	return searchResult, nil
}
