package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Updates a document in Kuzzle.
*/
func (dc Collection) UpdateDocument(id string, document interface{}, options types.QueryOptions) (types.Document, error) {
	if id == "" {
		return types.Document{}, errors.New("Collection.UpdateDocument: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "update",
		Body:       document,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.Document{}, errors.New(res.Error.Message)
	}

	documentResponse := types.Document{}
	json.Unmarshal(res.Result, &documentResponse)

	return documentResponse, nil
}

/*
  Update the provided documents.
*/
func (dc Collection) MUpdateDocument(documents []types.Document, options types.QueryOptions) (types.KuzzleSearchResult, error) {
	if len(documents) == 0 {
		return types.KuzzleSearchResult{}, errors.New("Collection.MUpdateDocument: please provide at least one document to update")
	}

	ch := make(chan types.KuzzleResponse)

	type CreationDocument struct {
		Id   string      `json:"_id"`
		Body interface{} `json:"body"`
	}
	docs := []CreationDocument{}

	type body struct {
		Documents []CreationDocument `json:"documents"`
	}

	for _, doc := range documents {
		docs = append(docs, CreationDocument{
			Id:   doc.Id,
			Body: doc.Content,
		})
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mUpdate",
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	result := types.KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	return result, nil
}
