package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ReplaceDocument replaces a document in Kuzzle.
func (d *Document) Replace(index string, collection string, _id string, body string, options types.QueryOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.Replace: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.Replace: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.Replace: id required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.Replace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "replace",
		Body:       body,
		Id:         _id,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var replaced string
	json.Unmarshal(res.Result, &replaced)

	return replaced, nil
}
