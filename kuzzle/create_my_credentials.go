package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "errors"
  "reflect"
  "encoding/json"
)

/*
 * Create credentials of the specified strategy for the current user.
 */
func (k Kuzzle) CreateMyCredentials(strategy string, credentials interface{}, options *types.Options) (map[string]interface{}, error) {
  type body struct {
    Strategy string `json:"strategy"`
    Body     interface{} `json:"body"`
  }
  result := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Controller: "auth",
    Action:     "createMyCredentials",
    Body:       &body{strategy, credentials},
  }
  go k.Query(query, options, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
  json.Unmarshal(res.Result, &ref)

  return ref.(map[string]interface{}), nil
}
