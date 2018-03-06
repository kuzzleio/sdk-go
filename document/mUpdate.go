package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MUpdateDocument mUpdates a document in Kuzzle.
func (d *Document) MUpdate(index string, collection string, body string, options *DocumentOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.MUpdate: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.MUpdate: collection required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.MUpdate: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "mUpdate",
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

	var mUpdated string
	json.Unmarshal(res.Result, &mUpdated)

	return mUpdated, nil
}
