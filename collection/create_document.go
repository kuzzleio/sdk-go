package collection

import (
  "github.com/kuzzleio/sdk-go/internal"
  "errors"
  "encoding/json"
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
func (dc *Collection) CreateDocument(id string, document interface{}, options *types.Options) (*types.Document, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(internal.BuildQuery(dc.collection, dc.index, "document", "create", document), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  documentResponse := &types.Document{}
  json.Unmarshal(res.Result, documentResponse)

  return documentResponse, nil
}