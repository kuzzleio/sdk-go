package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchUsers(filters interface{}, options types.QueryOptions) (*UserSearchResult, error) {
	res, err := s.rawSearch("searchUsers", filters, options)

	if err != nil {
		return nil, err
	}

	jsonSearchResult := &jsonUserSearchResult{}
	json.Unmarshal(res, jsonSearchResult)

	searchResult := &UserSearchResult{
		Total:    jsonSearchResult.Total,
		ScrollId: jsonSearchResult.ScrollId,
	}

	for _, j := range jsonSearchResult.Hits {
		searchResult.Hits = append(searchResult.Hits, s.jsonUserToUser(j))
	}

	return searchResult, nil
}
