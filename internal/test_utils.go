package internal

import (
  "github.com/stretchr/testify/mock"
  "github.com/kuzzleio/sdk-go/types"
)

type MockedKuzzle struct {
  mock.Mock
  MockQuery func() types.KuzzleResponse
}

func (k *MockedKuzzle) Query(query types.KuzzleRequest, res chan<- types.KuzzleResponse, options *types.Options) {
  if k.MockQuery != nil {
    res <- k.MockQuery()
  }
}
