package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchProfiles(body string, options types.QueryOptions) (*ProfileSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchProfiles",
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

	jsonSearchResult := &jsonProfileSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &ProfileSearchResult{
		ScrollId: jsonSearchResult.ScrollId,
		Total:    jsonSearchResult.Total,
	}
	for _, j := range jsonSearchResult.Hits {
		p := s.jsonProfileToProfile(j)
		p.Security = s
		searchResult.Hits = append(searchResult.Hits, p)
	}

	return searchResult, nil
}
