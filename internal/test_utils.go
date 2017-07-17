package internal

import (
  "github.com/stretchr/testify/mock"
  "github.com/kuzzleio/sdk-go/types"
)

type MockedConnection struct {
  mock.Mock
  MockSend func([]byte) types.KuzzleResponse
}

type Connection interface {
  Connect() (bool, error)
  Send([]byte, *types.Options, chan<- types.KuzzleResponse, string) error
  Close() error
}

func (c *MockedConnection) Send(query []byte, options *types.Options, responseChannel chan<- types.KuzzleResponse, requestId string) error {
  if c.MockSend != nil {
    responseChannel <- c.MockSend(query)
  }

  return nil
}

func (c *MockedConnection) Connect() (bool, error) {
  return false, nil
}

func (c *MockedConnection) Close() error {
  return nil
}
