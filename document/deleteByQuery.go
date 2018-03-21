package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) DeleteByQuery(index string, collection string, body string, options types.QueryOptions) ([]string, error) {
	if index == "" {
		return nil, types.NewError("Document.MCreate: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.MCreate: collection required", 400)
	}

	if body == "" {
		return nil, types.NewError("Document.MCreate: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "deleteByQuery",
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var deleted []string
	json.Unmarshal(res.Result, &deleted)

	return deleted, nil

}
