package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

func (s *Security) SearchProfiles(filters interface{}, options types.QueryOptions) (*ProfileSearchResult, error) {
	res, err := s.rawSearch("searchProfiles", filters, options)

	if err != nil {
		return nil, err
	}

	jsonSearchResult := &jsonProfileSearchResult{}
	json.Unmarshal(res, jsonSearchResult)

	searchResult := &ProfileSearchResult{
		ScrollId: jsonSearchResult.ScrollId,
		Total: jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		p := s.jsonProfileToProfile(j)
		p.Security = s
		searchResult.Hits = append(searchResult.Hits, p)
	}

	return searchResult, nil
}
