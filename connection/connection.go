package connection

import (
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

type Connection interface {
	AddListener(event int, channel chan<- interface{})
	RemoveListener(event int)
	Connect() (bool, error)
	Send([]byte, types.QueryOptions, chan<- types.KuzzleResponse, string) error
	Close() error
	GetOfflineQueue() *[]types.QueryObject
	GetState() *int
	EmitEvent(int, interface{})
	RegisterRoom(string, string, types.IRoom)
	UnregisterRoom(string)
	GetRequestHistory() *map[string]time.Time
	RenewSubscriptions()
	GetRooms() *types.RoomList
	ReplayQueue()
}
