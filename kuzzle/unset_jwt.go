package kuzzle

import (
	"github.com/kuzzleio/sdk-go/types"
	"sync"
)

/*
 * Unset the authentication token and cancel all subscriptions
 */
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	rooms := k.socket.GetRooms()
	if rooms != nil {
		k.socket.GetRooms().Range(func(key, value interface{}) bool {
			value.(*sync.Map).Range(func(key, value interface{}) bool {
				value.(types.IRoom).Renew(value.(types.IRoom).GetFilters(), value.(types.IRoom).GetRealtimeChannel(), value.(types.IRoom).GetResponseChannel())

				return true
			})

			return true
		})
	}
}
