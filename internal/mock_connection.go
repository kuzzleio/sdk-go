package internal

import (
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/mock"
	"time"
)

var offlineQueue []*types.QueryObject

type MockedConnection struct {
	mock.Mock
	MockSend      func([]byte, types.QueryOptions) *types.KuzzleResponse
	MockEmitEvent func(int, interface{})
	MockGetRooms  func() *types.RoomList
	state         int
}

func (c *MockedConnection) Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error {
	if c.MockSend != nil {
		responseChannel <- c.MockSend(query, options)
	}

	return nil
}

func (c *MockedConnection) Connect() (bool, error) {
	offlineQueue = make([]*types.QueryObject, 1)
	return false, nil
}

func (c *MockedConnection) Close() error {
	return nil
}

func (c *MockedConnection) AddListener(event int, channel chan<- interface{}) {}

func (c *MockedConnection) State() int {
	return c.state
}

func (c *MockedConnection) EmitEvent(event int, arg interface{}) {
	if c.MockEmitEvent != nil {
		c.MockEmitEvent(event, arg)
	}
}

func (c *MockedConnection) RegisterRoom(roomId, id string, room types.IRoom) {
}

func (c *MockedConnection) UnregisterRoom(id string) {}

func (c *MockedConnection) RequestHistory() map[string]time.Time {
	r := make(map[string]time.Time)

	return r
}

func (c *MockedConnection) RenewSubscriptions() {}

func (c *MockedConnection) Rooms() *types.RoomList {
	return c.MockGetRooms()
}

func (c *MockedConnection) StartQueuing() {}

func (c *MockedConnection) StopQueuing()             {}
func (c *MockedConnection) RemoveListener(event int) {}

func (c *MockedConnection) ReplayQueue() {}
func (c *MockedConnection) RemoveAllListeners(event int) {
}

func (c *MockedConnection) ClearQueue() {
	offlineQueue = nil
}

func (c *MockedConnection) AutoQueue() bool {
	return true
}

func (c *MockedConnection) AutoReconnect() bool {
	return true
}

func (c *MockedConnection) AutoResubscribe() bool {
	return true
}

func (c *MockedConnection) AutoReplay() bool {
	return true
}

func (c *MockedConnection) Host() string {
	return ""
}

func (c *MockedConnection) OfflineQueue() []*types.QueryObject {
	return offlineQueue
}

func (c *MockedConnection) OfflineQueueLoader() connection.OfflineQueueLoader {
	return nil
}

func (c *MockedConnection) Port() int {
	return 0
}

func (c *MockedConnection) QueueFilter() connection.QueueFilter {
	return nil
}

func (c *MockedConnection) QueueMaxSize() int {
	return 0
}

func (c *MockedConnection) QueueTTL() time.Duration {
	return 0
}

func (c *MockedConnection) ReplayInterval() time.Duration {
	return 0
}

func (c *MockedConnection) ReconnectionDelay() time.Duration {
	return 0
}

func (c *MockedConnection) SslConnection() bool {
	return false
}

func (c *MockedConnection) SetAutoQueue(v bool) {
}

func (c *MockedConnection) SetAutoReplay(v bool) {
}

func (c *MockedConnection) SetOfflineQueueLoader(v connection.OfflineQueueLoader) {
}

func (c *MockedConnection) SetQueueFilter(v connection.QueueFilter) {
}

func (c *MockedConnection) SetQueueMaxSize(v int) {
}

func (c *MockedConnection) SetQueueTTL(v time.Duration) {
}

func (c *MockedConnection) SetReplayInterval(v time.Duration) {
}

// mock specific functions
func (c *MockedConnection) SetState(value int) {
	c.state = value
}

func (c *MockedConnection) AddToOfflineQueue(q *types.QueryObject) {
	offlineQueue = append(offlineQueue, q)
}
