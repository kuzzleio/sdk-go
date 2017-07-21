package internal

import (
  "github.com/stretchr/testify/mock"
  "github.com/kuzzleio/sdk-go/types"
)

var OfflineQueue []types.QueryObject

type MockedConnection struct {
  mock.Mock
  MockSend func([]byte, *types.Options) types.KuzzleResponse
}

func (c MockedConnection) Send(query []byte, options *types.Options, responseChannel chan<- types.KuzzleResponse, requestId string) error {
  if c.MockSend != nil {
    responseChannel <- c.MockSend(query, options)
  }

  return nil
}

func (c MockedConnection) Connect() (bool, error) {
  OfflineQueue = make([]types.QueryObject, 1)
  return false, nil
}

func (c MockedConnection) Close() error {
  return nil
}

func (c MockedConnection) AddListener(event int, channel chan<- interface{}) {

}

func (c MockedConnection) GetState() *int {
  state := 0
  return &state
}
func (c MockedConnection) GetOfflineQueue() *[]types.QueryObject {
  return &OfflineQueue
}