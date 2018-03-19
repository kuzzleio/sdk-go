package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) Get(index string, collection string, _id string, options types.QueryOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.Get: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.Get: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.Get: id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "get",
		Id:         _id,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var doc string
	json.Unmarshal(res.Result, &doc)

	return doc, nil
}
