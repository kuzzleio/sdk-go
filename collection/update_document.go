package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// UpdateDocument updates a document in Kuzzle.
func (dc Collection) UpdateDocument(id string, document *Document, options types.QueryOptions) (*Document, error) {
	if id == "" {
		return &Document{}, errors.New("Collection.UpdateDocument: document id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "update",
		Body:       document.Content,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return &Document{}, errors.New(res.Error.Message)
	}

	documentResponse := &Document{collection: &dc}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}

// MUpdateDocument update the provided documents.
func (dc *Collection) MUpdateDocument(documents []*Document, options types.QueryOptions) (*SearchResult, error) {
	result := &SearchResult{}

	if len(documents) == 0 {
		return result, errors.New("Collection.MUpdateDocument: please provide at least one document to update")
	}

	ch := make(chan *types.KuzzleResponse)

	type CreationDocument struct {
		Id   string       `json:"_id"`
		Body interface{}  `json:"body"`
	}
	docs := []*CreationDocument{}

	type body struct {
		Documents []*CreationDocument `json:"documents"`
	}

	for _, doc := range documents {
		docs = append(docs, &CreationDocument{
			Id:   doc.Id,
			Body: doc.Content,
		})
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mUpdate",
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return result, errors.New(res.Error.Message)
	}

	json.Unmarshal(res.Result, result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
