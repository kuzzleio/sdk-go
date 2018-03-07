package document

import (
	"encoding/json"
	"strconv"

	"github.com/kuzzleio/sdk-go/types"
)

// MReplaceDocument mReplaces a document in Kuzzle.
func (d *Document) MReplace(index string, collection string, body string, options *DocumentOptions) (string, error) {
	if index == "" {
		return "", types.NewError("Document.MReplace: index required", 400)
	}

	if collection == "" {
		return "", types.NewError("Document.MReplace: collection required", 400)
	}

	if body == "" {
		return "", types.NewError("Document.MReplace: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "mReplace",
		Body:       body,
	}

	queryOpts := types.NewQueryOptions()

	if options != nil {
		queryOpts.SetRefresh(strconv.FormatBool(options.WaitFor))
	}

	go d.Kuzzle.Query(query, queryOpts, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var mReplaced string
	json.Unmarshal(res.Result, &mReplaced)

	return mReplaced, nil
}
