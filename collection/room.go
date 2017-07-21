package collection

import (
  "github.com/kuzzleio/sdk-go/types"
)

/*
  Instantiates a new Room object.
*/
func (dc Collection) Room(options *types.RoomOptions) (types.Room) {
  room := types.Room{}

  scopes := map[string]bool {types.SCOPE_ALL: true, types.SCOPE_IN: true, types.SCOPE_OUT: true, types.SCOPE_NONE: true, types.SCOPE_UNKNOWN: true}
  states := map[string]bool {types.STATE_ALL: true, types.STATE_PENDING: true, types.STATE_DONE: true}
  user := map[string]bool {types.USER_ALL: true, types.USER_IN: true, types.USER_OUT: true, types.USER_NONE: true}

  if options != nil {
    if scopes[options.Scope] {
      room.Scope = options.Scope
    }
    if states[options.State] {
      room.State = options.State
    }
    if user[options.User] {
      room.User = options.User
    }
    room.SubscribeToSelf = options.SubscribeToSelf
  }

  return room
}
