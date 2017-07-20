package collection

import (
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Retrieves a Document using its provided unique id.
*/
func (dc Collection) FetchDocument(id string, options *types.Options) (types.Document, error) {
  if id == "" {
    return types.Document{}, errors.New("Collection.FetchDocument: document id required")
  }

  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Collection: dc.collection,
    Index:      dc.index,
    Controller: "document",
    Action:     "get",
    Id:         id,
  }
  go dc.kuzzle.Query(query, options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return types.Document{}, errors.New(res.Error.Message)
  }

  document := types.Document{}
  json.Unmarshal(res.Result, &document)

  return document, nil
}
