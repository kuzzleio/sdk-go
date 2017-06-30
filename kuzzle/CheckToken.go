package kuzzle

import (
  "github.com/kuzzleio/sdk-core"
  "encoding/json"
  "errors"
  "github.com/kuzzleio/sdk-core/types"
)

func CheckToken(ik core.IKuzzle, token string) (*core.TokenValidity, error) {
  if token == "" {
    return nil, errors.New("Kuzzle.CheckToken: token required")
  }

  result := make(chan types.KuzzleResponse)

  type body struct {
    Token string `json:"token"`
  }
  ik.Query(core.MakeQuery("auth", "checkToken", "", "", &body{token}), result, nil)

  res := <- result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  tokenValidity := &core.TokenValidity{}
  json.Unmarshal(res.Result, &tokenValidity)

  return tokenValidity, nil
}
