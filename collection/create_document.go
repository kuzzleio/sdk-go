package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Create a new document in Kuzzle.
  Takes an optional argument object with the following properties:
     - volatile (object, default: null):
         Additional information passed to notifications to other users
     - ifExist (string, allowed values: "error" (default), "replace"):
         If the same document already exists:
           - resolves with an error if set to "error".
           - replaces the existing document if set to "replace"
*/
func (dc Collection) CreateDocument(id string, document interface{}, options *types.Options) (*types.Document, error) {
	ch := make(chan types.KuzzleResponse)

	action := "create"

	if options != nil {
		if options.IfExist == "replace" {
			action = "createOrReplace"
		} else if options.IfExist != "error" {
			return nil, errors.New(fmt.Sprintf("Invalid value for the 'ifExist' option: '%s'", options.IfExist))
		}
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     action,
		Body:       document,
		Id:         id,
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

/*
  Creates the provided documents.
*/
func (dc Collection) MCreateDocument(documents []types.Document, options *types.Options) (types.KuzzleSearchResult, error) {
	return performMultipleCreate(dc, documents, "mCreate", options)
}

/*
  Creates or replaces the provided documents.
*/
func (dc Collection) MCreateOrReplaceDocument(documents []types.Document, options *types.Options) (types.KuzzleSearchResult, error) {
	return performMultipleCreate(dc, documents, "mCreateOrReplace", options)
}

func performMultipleCreate(dc Collection, documents []types.Document, action string, options *types.Options) (types.KuzzleSearchResult, error) {
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
			Body: doc.Source,
		})
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "document",
		Action:     action,
		Body:       &body{docs},
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
