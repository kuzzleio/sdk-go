package internal

import (
  "github.com/kuzzleio/sdk-go/types"
)

func BuildQuery(collection, index, controller, action string, body interface{}) types.KuzzleRequest {
  return types.KuzzleRequest{
    Controller: controller,
    Action:     action,
    Index:      index,
    Collection: collection,
    Body:       body,
  }
}
