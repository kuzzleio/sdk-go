package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "fmt"
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
