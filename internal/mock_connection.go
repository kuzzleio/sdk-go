package internal

import (
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/mock"
	"time"
)

var OfflineQueue []types.QueryObject

type MockedConnection struct {
	mock.Mock
	MockSend      func([]byte, types.QueryOptions) types.KuzzleResponse
	MockEmitEvent func(int, interface{})
	MockGetRooms  func() *types.RoomList
}

func (c MockedConnection) Send(query []byte, options types.QueryOptions, responseChannel chan<- types.KuzzleResponse, requestId string) error {
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

func (c MockedConnection) AddListener(event int, channel chan<- interface{}) {}

func (c MockedConnection) GetState() *int {
	state := 0
	return &state
}

func (c MockedConnection) GetOfflineQueue() *[]types.QueryObject {
	return &OfflineQueue
}

func (c MockedConnection) EmitEvent(event int, arg interface{}) {
	if c.MockEmitEvent != nil {
		c.MockEmitEvent(event, arg)
	}
}

func (c MockedConnection) RegisterRoom(roomId, id string, room types.IRoom) {
}

func (c MockedConnection) UnregisterRoom(id string) {}

func (c MockedConnection) GetRequestHistory() *map[string]time.Time {
	r := make(map[string]time.Time)

	return &r
}

func (c MockedConnection) RenewSubscriptions() {}

func (c MockedConnection) GetRooms() *types.RoomList {
	v := c.MockGetRooms()

	return v
}

func (c MockedConnection) RemoveListener(event int) {}
func (c MockedConnection) RemoveAllListeners(event int) {
}
