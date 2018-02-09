package connection

import (
	"time"

	"github.com/kuzzleio/sdk-go/types"
)

type Connection interface {
	AddListener(event int, channel chan<- interface{})
	RemoveListener(event int, channel chan<- interface{})
	RemoveAllListeners(event int)
	Once(event int, channel chan<- interface{})
	ListenerCount(event int) int
	Connect() (bool, error)
	Send([]byte, types.QueryOptions, chan<- *types.KuzzleResponse, string) error
	Close() error
	State() int
	EmitEvent(int, interface{})
	RegisterRoom(types.IRoom)
	UnregisterRoom(string)
	RequestHistory() map[string]time.Time
	Rooms() *types.RoomList
	StartQueuing()
	StopQueuing()
	ReplayQueue()
	ClearQueue()

	// property getters
	AutoQueue() bool
	AutoReconnect() bool
	AutoResubscribe() bool
	AutoReplay() bool
	Host() string
	OfflineQueue() []*types.QueryObject
	OfflineQueueLoader() OfflineQueueLoader
	Port() int
	QueueFilter() QueueFilter
	QueueMaxSize() int
	QueueTTL() time.Duration
	ReplayInterval() time.Duration
	ReconnectionDelay() time.Duration
	SslConnection() bool

	// property setters
	SetAutoQueue(bool)
	SetAutoReplay(bool)
	SetOfflineQueueLoader(OfflineQueueLoader)
	SetQueueFilter(QueueFilter)
	SetQueueMaxSize(int)
	SetQueueTTL(time.Duration)
	SetReplayInterval(time.Duration)
}

type OfflineQueueLoader interface {
	Load() []*types.QueryObject
}

type QueueFilter func([]byte) bool
