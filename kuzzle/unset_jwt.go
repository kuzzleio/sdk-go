package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
	"sync"
)

// UnsetJwt unset the authentication token and cancel all subscriptions
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	rooms := k.socket.GetRooms()
	if rooms != nil {
		k.socket.GetRooms().Range(func(key, value interface{}) bool {
			value.(*sync.Map).Range(func(key, value interface{}) bool {
				room := value.(types.IRoom)
				room.Renew(room.GetFilters(), room.GetRealtimeChannel(), room.GetResponseChannel())

				return true
			})

			return true
		})
	}
}
