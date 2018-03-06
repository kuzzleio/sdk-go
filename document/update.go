package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// UpdateDocument updates a document in Kuzzle.
func (d *Document) Update(index string, collection string, _id string, body string, options *DocumentOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.Update: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.Update: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.Update: id required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.Update: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "update",
		Body:       body,
		Id:         _id,
	}

	queryOpts := types.NewQueryOptions()

	if options != nil {
		queryOpts.SetVolatile(options.Volatile)
		queryOpts.SetRefresh(options.WaitFor)
	}

	go d.Kuzzle.Query(query, queryOpts, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var updated string
	json.Unmarshal(res.Result, &updated)

	return updated, nil
}
