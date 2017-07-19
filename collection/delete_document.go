package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Deletes the Document using its provided unique id.
*/
func (dc Collection) DeleteDocument(id string, options *types.Options) (string, error) {
  if id == "" {
    return "", errors.New("Collection.DeleteDocument: document id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Collection: dc.collection,
    Index:      dc.index,
    Controller: "document",
    Action:     "delete",
    Id:         id,
  }
  go dc.kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return "", errors.New(res.Error.Message)
  }

  document := &types.Document{}
  json.Unmarshal(res.Result, document)

  return document.Id, nil
}
