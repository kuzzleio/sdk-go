package collection

import (
	"errors"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Replaces a document in Kuzzle.
*/
func (dc Collection) ReplaceDocument(id string, document interface{}, options *types.Options) (*types.Document, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "replace",
		Body:       document,
		Id: id,
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	documentResponse := &types.Document{}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}
