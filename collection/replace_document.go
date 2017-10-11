package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// ReplaceDocument replaces a document in Kuzzle.
func (dc *Collection) ReplaceDocument(id string, document interface{}, options types.QueryOptions) (*Document, error) {
	if id == "" {
		return &Document{}, types.NewError("Collection.ReplaceDocument: document id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "replace",
		Body:       document,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return &Document{}, res.Error
	}

	d := &Document{collection: dc}
	json.Unmarshal(res.Result, d)

	return d, nil
}

// MReplaceDocument replaces the provided documents.
func (dc *Collection) MReplaceDocument(documents []*Document, options types.QueryOptions) (*SearchResult, error) {
	result := &SearchResult{}

	if len(documents) == 0 {
		return result, types.NewError("Collection.MReplaceDocument: please provide at least one document to replace")
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
		Action:     "mReplace",
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return result, res.Error
	}

	json.Unmarshal(res.Result, result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
