package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MReplaceDocument mReplaces a document in Kuzzle.
func (d *Document) MReplace(index string, collection string, body string, options types.QueryOptions) (string, error) {
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

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return "", res.Error
	}

	var mReplaced string
	json.Unmarshal(res.Result, &mReplaced)

	return mReplaced, nil
}
