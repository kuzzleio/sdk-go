package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Deletes the Document using its provided unique id.
*/
func (dc Collection) DeleteDocument(id string, options types.QueryOptions) (string, error) {
	if id == "" {
		return "", errors.New("Collection.DeleteDocument: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "delete",
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	document := Document{collection: dc}
	json.Unmarshal(res.Result, &document)

	return document.Id, nil
}

/*
  Delete specific documents according to given IDs.
*/
func (dc Collection) MDeleteDocument(ids []string, options types.QueryOptions) ([]string, error) {
	if len(ids) == 0 {
		return []string{}, errors.New("Collection.MDeleteDocument: please provide at least one id of document to delete")
	}

	ch := make(chan types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mDelete",
		Body:       &body{Ids: ids},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return []string{}, errors.New(res.Error.Message)
	}

	result := []string{}
	json.Unmarshal(res.Result, &result)

	return result, nil
}
