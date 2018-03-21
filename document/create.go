package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) Create(index string, collection string, _id string, body string, options types.QueryOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.Create: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.Create: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.Create: id required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.Create: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "create",
		Id:         _id,
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var created string
	json.Unmarshal(res.Result, &created)

	return created, nil
}
