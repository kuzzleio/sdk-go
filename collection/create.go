package collection

import (
  "github.com/kuzzleio/sdk-go/internal"
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Create a new empty data collection, with no associated mapping.
*/
func (dc Collection) Create(options *types.Options) (*types.AckResponse, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(internal.BuildQuery(dc.collection, dc.index, "collection", "create", nil), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ack := &types.AckResponse{}
  json.Unmarshal(res.Result, &ack)

  return ack, nil
}
