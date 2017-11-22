package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
	"sync"
)

// UnsetJwt unset the authentication token and cancel all subscriptions
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	rooms := k.socket.Rooms()
	if rooms != nil {
		k.socket.Rooms().Range(func(key, value interface{}) bool {
			value.(*sync.Map).Range(func(key, value interface{}) bool {
				room := value.(types.IRoom)
				room.Renew(room.Filters(), room.RealtimeChannel(), room.ResponseChannel())

				return true
			})

			return true
		})
	}
}
