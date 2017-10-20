package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// DocumentExists returns a boolean indicating whether or not a document with provided ID exists.
func (dc Collection) DocumentExists(id string, options types.QueryOptions) (bool, error) {
	if id == "" {
		return false, types.NewError("Collection.DocumentExists: document id required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "exists",
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return false, res.Error
	}

	exists, _ := strconv.ParseBool(string(res.Result))

	return exists, nil
}
