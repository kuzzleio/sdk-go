package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Search documents in the given Collection, using provided filters and option.
func (d *Document) Search(index string, collection string, body string, options types.QueryOptions) (*types.SearchResult, error) {
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

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	sr := &types.SearchResult{}
	json.Unmarshal(res.Result, &sr)

	return sr, nil
}
