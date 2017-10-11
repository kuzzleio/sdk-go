package kuzzle

import (
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

func (k *Kuzzle) SetJwt(token string) {
	k.jwt = token

	if token != "" {
		k.socket.RenewSubscriptions()
		k.socket.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: true})
	}
}
