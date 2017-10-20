package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

func (s *Security) ScrollUsers(scrollId string, options types.QueryOptions) (*UserSearchResult, error) {
	res, err := s.rawScroll("scrollUsers", scrollId, options)

	if err != nil {
		return nil, err
	}

	jsonSearchResult := &jsonUserSearchResult{}
	json.Unmarshal(res, jsonSearchResult)

	searchResult := &UserSearchResult{
		Total: jsonSearchResult.Total,
		ScrollId: jsonSearchResult.ScrollId,
	}

	for _, j := range jsonSearchResult.Hits {
		searchResult.Hits = append(searchResult.Hits, s.jsonUserToUser(j))
	}

	return searchResult, nil
}
