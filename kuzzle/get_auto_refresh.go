package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "fmt"
  "encoding/json"
)

func (k Kuzzle) GetAutoRefresh(index string, options *types.Options) (bool, error) {
  if index == "" {
    if k.defaultIndex == "" {
      return false, errors.New("Kuzzle.GetAutoRefresh: index required")
    }
    index = k.defaultIndex
  }

  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "index",
    Action:     "getAutoRefresh",
    Index:      index,
  }
  go k.Query(query, options, result)

  res := <-result

  if res.Error.Message != "" {
    return false, errors.New(res.Error.Message)
  }

  var r bool
  json.Unmarshal(res.Result, &r)

  return r, nil
}
