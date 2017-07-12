package kuzzle

import (
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "errors"
)

func (k *Kuzzle) CreateIndex(index string, options *types.Options) (*types.AckResponse, error) {
  if index == "" {
    return nil, errors.New("Kuzzle.createIndex: index required")
  }

  result := make(chan types.KuzzleResponse)

  go k.Query(buildQuery("index", "create", index, "", nil), options, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ack := &types.AckResponse{}
  json.Unmarshal(res.Result, &ack)

  return ack, nil
}