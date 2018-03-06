package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) CreateOrReplace(index string, collection string, _id string, body string, options *DocumentOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.CreateOrReplace: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.CreateOrReplace: collection required", 400)
	}

	if _id == "" {
		return "", types.NewError("Document.CreateOrReplace: id required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.CreateOrReplace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "createOrReplace",
		Id:         _id,
		Body:       body,
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

	var created string
	json.Unmarshal(res.Result, &created)

	return created, nil
}
