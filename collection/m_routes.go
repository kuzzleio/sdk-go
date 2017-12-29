package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

type mActionResult struct {
	Hits  []json.RawMessage `json:"hits"`
	Total int               `json:"total"`
}

/*
  Creates the provided documents.
*/
func (dc *Collection) MCreateDocument(documents []*Document, options types.QueryOptions) ([]*Document, error) {
	return performMultipleCreate(dc, documents, "mCreate", options)
}

/*
  Creates or replaces the provided documents.
*/
func (dc *Collection) MCreateOrReplaceDocument(documents []*Document, options types.QueryOptions) ([]*Document, error) {
	return performMultipleCreate(dc, documents, "mCreateOrReplace", options)
}

func (dc *Collection) parseMultiActionsResult(response *types.KuzzleResponse) ([]*Document, error) {
	if response.Error != nil {
		return nil, response.Error
	}

	var rawResult mActionResult
	json.Unmarshal(response.Result, &rawResult)

	result := make([]*Document, rawResult.Total)

	for i, d := range rawResult.Hits {
		json.Unmarshal(d, &result[i])
		result[i].collection = dc
	}

	return result, nil
}

func performMultipleCreate(dc *Collection, documents []*Document, action string, options types.QueryOptions) ([]*Document, error) {
	ch := make(chan *types.KuzzleResponse)

	type CreationDocument struct {
		Id   string      `json:"_id"`
		Body interface{} `json:"body"`
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
		Action:     action,
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	return dc.parseMultiActionsResult(res)
}

// MReplaceDocument replaces multiple documents at once
func (dc *Collection) MReplaceDocument(documents []*Document, options types.QueryOptions) ([]*Document, error) {
	if len(documents) == 0 {
		return nil, types.NewError("Collection.MReplaceDocument: please provide at least one document to replace", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type CreationDocument struct {
		Id   string      `json:"_id"`
		Body interface{} `json:"body"`
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

	return dc.parseMultiActionsResult(res)
}

// MDeleteDocument deletes multiple documents at once
func (dc *Collection) MDeleteDocument(ids []string, options types.QueryOptions) ([]string, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Collection.MDeleteDocument: please provide at least one id of document to delete", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mDelete",
		Body:       &body{Ids: ids},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	result := []string{}
	json.Unmarshal(res.Result, &result)

	return result, nil
}

// MGetDocument fetches multiple documents at once
func (dc *Collection) MGetDocument(ids []string, options types.QueryOptions) ([]*Document, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Collection.MGetDocument: please provide at least one id of document to retrieve", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     "mGet",
		Body:       &body{Ids: ids},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	return dc.parseMultiActionsResult(res)
}

// MUpdateDocument updates multiple documents at once
func (dc *Collection) MUpdateDocument(documents []*Document, options types.QueryOptions) ([]*Document, error) {
	if len(documents) == 0 {
		return nil, types.NewError("Collection.MUpdateDocument: please provide at least one document to update", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type CreationDocument struct {
		Id   string      `json:"_id"`
		Body interface{} `json:"body"`
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

	return dc.parseMultiActionsResult(res)
}
