package collection

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
)

// Create a new document in Kuzzle.
// Takes an optional argument object with the following properties:
//   - volatile (object, default: null):
//       Additional information passed to notifications to other users
//   - ifExist (string, allowed values: "error" (default), "replace"):
//       If the same document already exists:
//         - resolves with an error if set to "error".
//         - replaces the existing document if set to "replace"
func (dc *Collection) CreateDocument(id string, document *Document, options types.QueryOptions) (*Document, error) {
	ch := make(chan *types.KuzzleResponse)

	action := "create"

	if options != nil {
		if options.GetIfExist() == "replace" {
			action = "createOrReplace"
		} else if options.GetIfExist() != "error" {
			return nil, types.NewError(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.GetIfExist()), 400)
		}
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     action,
		Body:       document.Content,
		Id:         id,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	documentResponse := &Document{collection: dc}
	json.Unmarshal(res.Result, documentResponse)

	return documentResponse, nil
}

/*
  Creates the provided documents.
*/
func (dc *Collection) MCreateDocument(documents []*Document, options types.QueryOptions) (*SearchResult, error) {
	return performMultipleCreate(dc, documents, "mCreate", options)
}

/*
  Creates or replaces the provided documents.
*/
func (dc *Collection) MCreateOrReplaceDocument(documents []*Document, options types.QueryOptions) (*SearchResult, error) {
	return performMultipleCreate(dc, documents, "mCreateOrReplace", options)
}

func performMultipleCreate(dc *Collection, documents []*Document, action string, options types.QueryOptions) (*SearchResult, error) {
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
		Action:     action,
		Body:       &body{docs},
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	result := &SearchResult{}
	json.Unmarshal(res.Result, result)

	for _, d := range result.Hits {
		d.collection = dc
	}

	return result, nil
}
