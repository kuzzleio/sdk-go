package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// SearchSpecifications searches specifications across indexes/collections according to the provided filters.
func (dc *Collection) SearchSpecifications(options types.SpecificationSearchOptions) (*types.SpecificationSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
		Body: struct {
			Query interface{} `json:"query"`
		}{Query: filters},
	}

	if options != nil {
		query.From = options.From()
		query.Size = options.Size()
		scroll := options.Scroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specifications := &types.SpecificationSearchResult{}
	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}
