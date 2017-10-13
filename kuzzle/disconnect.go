package kuzzle

import "github.com/kuzzleio/sdk-go/types"

// Disconnect from Kuzzle and invalidate this instance.
// Does not fire a disconnected event.
func (k Kuzzle) Disconnect() error {
	err := k.socket.Close()

	if err != nil {
		return types.NewError(err.Error())
	}
	k.wasConnected = false

	return nil
}
