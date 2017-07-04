package kuzzle

import (
  "github.com/kuzzleio/sdk-go/core"
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/utils"
)

func CreateIndex(ik core.IKuzzle, index string, options *types.Options) (*types.AckResponse, error) {
  if index == "" {
    return nil, errors.New("Kuzzle.createIndex: index required")
  }

  result := make(chan types.KuzzleResponse)
  defer close(result)

  type body struct {
     Index string `json:"index"`
  }

  go ik.Query(utils.MakeQuery("index", "create", "", "", &body{index}), result, options)

  res := <- result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ack := &types.AckResponse{}
  json.Unmarshal(res.Result, &ack)

  return ack, nil
}