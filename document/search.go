package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Search documents in the given Collection, using provided filters and option.
func (d *Document) Search(index string, collection string, body string, options *types.SearchOptions) (*types.SearchResult, error) {
	if index == "" {
		return nil, types.NewError("Document.Search: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.Search: collection required", 400)
	}

	if body == "" {
		return nil, types.NewError("Document.Search: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "search",
		Body:       body,
	}

	if options != nil {
		query.From = *options.From
		query.Size = *options.Size
		scroll := options.Scroll
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go d.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	sr := &types.SearchResult{}
	json.Unmarshal(res.Result, &sr)

	return sr, nil
}
