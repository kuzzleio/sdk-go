package kuzzle

// Remove all listener by event type or all listener if event == -1
func (k Kuzzle) RemoveAllListeners(event int) {
	k.socket.RemoveAllListeners(event)
}
