package test

import (
  "testing"
  "github.com/kuzzleio/sdk-core/types"
  "github.com/kuzzleio/sdk-core/kuzzle"
)

type MockedKuzzle struct {}

func (k MockedKuzzle) Query(query types.KuzzleRequest, res chan<- types.KuzzleResponse, subscription chan<- types.KuzzleNotification) {
  println("mock ok")
}

func TestCheckToken(t *testing.T) {
  k := &MockedKuzzle{}

  kuzzle.CheckToken(k, "token")
}
