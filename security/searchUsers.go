package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (s *Security) SearchUsers(body string, options types.QueryOptions) (*UserSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
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

	jsonSearchResult := &jsonUserSearchResult{}
	json.Unmarshal(res.Result, jsonSearchResult)

	searchResult := &UserSearchResult{
		Total:    jsonSearchResult.Total,
		ScrollId: jsonSearchResult.ScrollId,
	}

	for _, j := range jsonSearchResult.Hits {
		searchResult.Hits = append(searchResult.Hits, s.jsonUserToUser(j))
	}

	return searchResult, nil
}
