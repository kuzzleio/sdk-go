package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Replaces a document in Kuzzle.
*/
func (dc Collection) ReplaceDocument(id string, document interface{}, options types.QueryOptions) (Document, error) {
	if id == "" {
		return Document{}, errors.New("Collection.ReplaceDocument: document id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "replace",
		Body:       document,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return Document{}, errors.New(res.Error.Message)
	}

	d := Document{collection: dc}
	json.Unmarshal(res.Result, &d)

	return d, nil
}

/*
  Replace the provided documents.
*/
func (dc Collection) MReplaceDocument(documents []Document, options types.QueryOptions) (KuzzleSearchResult, error) {
	if len(documents) == 0 {
		return KuzzleSearchResult{}, errors.New("Collection.MReplaceDocument: please provide at least one document to replace")
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
		Action:     "mReplace",
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return KuzzleSearchResult{}, errors.New(res.Error.Message)
	}

	result := KuzzleSearchResult{}
	json.Unmarshal(res.Result, &result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
