package collection

import (
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

/*
  Returns a boolean indicating whether or not a document with provided ID exists.
*/
func (dc Collection) DocumentExists(id string, options types.QueryOptions) (bool, error) {
	if id == "" {
		return false, errors.New("Collection.DocumentExists: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "exists",
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return false, errors.New(res.Error.Message)
	}

	exists, _ := strconv.ParseBool(string(res.Result))

	return exists, nil
}
