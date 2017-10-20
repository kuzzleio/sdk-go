package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

func (s *Security) ScrollProfiles(scrollId string, options types.QueryOptions) (*ProfileSearchResult, error) {
	res, err := s.rawScroll("scrollProfiles", scrollId, options)

	if err != nil {
		return nil, err
	}

	jsonSearchResult := &jsonProfileSearchResult{}
	json.Unmarshal(res, jsonSearchResult)

	searchResult := &ProfileSearchResult{
		Total: jsonSearchResult.Total,
		ScrollId: jsonSearchResult.ScrollId,
	}

	for _, j := range jsonSearchResult.Hits {
		p := s.jsonProfileToProfile(j)
		p.Security = s
		searchResult.Hits = append(searchResult.Hits, p)
	}

	return searchResult, nil
}
