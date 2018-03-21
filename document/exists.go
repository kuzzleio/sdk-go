package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) Exists(index string, collection string, _id string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Document.Exists: index required", 400)
	}

	if collection == "" {
		return false, types.NewError("Document.Exists: collection required", 400)
	}

	if _id == "" {
		return false, types.NewError("Document.Exists: id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "exists",
		Id:         _id,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	var exists bool
	json.Unmarshal(res.Result, &exists)

	return exists, nil
}
