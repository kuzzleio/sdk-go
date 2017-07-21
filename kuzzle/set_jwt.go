package kuzzle

import (
  "github.com/kuzzleio/sdk-go/event"
  "github.com/kuzzleio/sdk-go/types"
)

func (k *Kuzzle) SetJwt(jwt string) {
  k.socket.EmitEvent(event.LoginAttempt, types.LoginAttempt{Success: true})
  k.jwt = jwt
}