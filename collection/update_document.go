package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Updates a document in Kuzzle.
*/
func (dc Collection) UpdateDocument(id string, document interface{}, options *types.Options) (types.Document, error) {
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
  go dc.kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return types.Document{}, errors.New(res.Error.Message)
  }

  documentResponse := types.Document{}
  json.Unmarshal(res.Result, &documentResponse)

  return documentResponse, nil
}
