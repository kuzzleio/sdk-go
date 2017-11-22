package connection

import (
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

type Connection interface {
	AddListener(event int, channel chan<- interface{})
	RemoveListener(event int)
	RemoveAllListeners(event int)
	Connect() (bool, error)
	Send([]byte, types.QueryOptions, chan<- *types.KuzzleResponse, string) error
	Close() error
	OfflineQueue() *[]*types.QueryObject
	State() int
	EmitEvent(int, interface{})
	RegisterRoom(string, string, types.IRoom)
	UnregisterRoom(string)
	RequestHistory() map[string]time.Time
	RenewSubscriptions()
	Rooms() *types.RoomList
	StartQueuing()
	StopQueuing()
	ReplayQueue()
}
