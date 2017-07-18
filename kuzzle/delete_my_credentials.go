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
func (k Kuzzle) DeleteMyCredentials(strategy string, credentials interface{}, options *types.Options) (map[string]interface{}, error) {
  type body struct {
    Strategy string `json:"strategy"`
    Body     interface{} `json:"body"`
  }
  result := make(chan types.KuzzleResponse)

  go k.Query(buildQuery("", "", "auth", "deleteMyCredentials", &body{strategy, credentials}), options, result)

  res := <-result

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ref := reflect.New(reflect.TypeOf(credentials)).Elem().Interface()
  json.Unmarshal(res.Result, &ref)

  return ref.(map[string]interface{}), nil
}
