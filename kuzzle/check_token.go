package kuzzle

import (
  "errors"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/core"
  "github.com/kuzzleio/sdk-go/utils"
  "encoding/json"
)

// Checks the validity of a JSON Web Token.
type TokenValidity struct {
  Valid bool `json:"valid"`
  State string `json:"state"`
  ExpiresAt int `json:"expiresAt"`
}

func CheckToken(ik core.IKuzzle, token string) (*TokenValidity, error) {
  if token == "" {
    return nil, errors.New("Kuzzle.CheckToken: token required")
  }

  result := make(chan types.KuzzleResponse)
  defer close(result)

  type body struct {
    Token string `json:"token"`
  }

  go ik.Query(utils.MakeQuery("auth", "checkToken", "", "", &body{token}), result, nil)

  res := <- result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  tokenValidity := &TokenValidity{}
  json.Unmarshal(res.Result, &tokenValidity)

  return tokenValidity, nil
}
