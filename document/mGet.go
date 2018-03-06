package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) MGet(index string, collection string, ids []string, includeTrash bool) (string, error) {
	if index == "" {
		return "", types.NewError("Document.MGet: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.MGet: collection required", 400)
	}

	if len(ids) == 0 {
		return "", types.NewError("Document.MGet: ids filled array required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Ids          []string `json:"ids"`
		IncludeTrash bool     `json:"includeTrash"`
	}

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mGet",
		Body:       &body{Ids: ids, IncludeTrash: includeTrash},
	}

	go d.Kuzzle.Query(query, nil, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var docs string
	json.Unmarshal(res.Result, &docs)

	return docs, nil
}
