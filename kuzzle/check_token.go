package kuzzle

import (
  "encoding/json"
  "errors"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/internal"
)

// Checks the validity of a JSON Web Token.
type TokenValidity struct {
  Valid bool `json:"valid"`
  State string `json:"state"`
  ExpiresAt int `json:"expiresAt"`
}

func (k Kuzzle) CheckToken(token string) (*TokenValidity, error) {
  if token == "" {
    return nil, errors.New("Kuzzle.CheckToken: token required")
  }

  result := make(chan types.KuzzleResponse)

  type body struct {
    Token string `json:"token"`
  }

  go k.Query(internal.BuildQuery("auth", "checkToken", "", "", &body{token}), nil, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  tokenValidity := &TokenValidity{}
  json.Unmarshal(res.Result, &tokenValidity)

  return tokenValidity, nil
}
