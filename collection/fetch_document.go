package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Retrieves a Document using its provided unique id.
*/
func (dc Collection) FetchDocument(id string, options types.QueryOptions) (types.Document, error) {
	if id == "" {
		return types.Document{}, errors.New("Collection.FetchDocument: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "get",
		Id:         id,
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Document{}, errors.New(res.Error.Message)
	}

	document := types.Document{}
	json.Unmarshal(res.Result, &document)

	return document, nil
}

/*
  Get specific documents according to given IDs.
*/
func (dc Collection) MGetDocument(ids []string, options types.QueryOptions) (types.KuzzleSearchResult, error) {
	if len(ids) == 0 {
		return types.KuzzleSearchResult{}, errors.New("Collection.MGetDocument: please provide at least one id of document to retrieve")
	}

	ch := make(chan types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mGet",
		Body:       &body{Ids: ids},
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	result := types.KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	return result, nil
}
