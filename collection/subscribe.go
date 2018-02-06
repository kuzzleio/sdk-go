package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Subscribe subscribes to this data collection with a set of Kuzzle DSL filters.
func (dc *Collection) Subscribe(filters interface{}, options types.RoomOptions, realtimeNotificationChannel chan<- types.KuzzleNotification) *Room {
	r := NewRoom(dc, filters, options)

	res := make(chan types.SubscribeResponse)
	r.OnDone(res)
	r.Subscribe(realtimeNotificationChannel)

	return r
}
