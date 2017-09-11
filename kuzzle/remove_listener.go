package kuzzle

func (k Kuzzle) RemoveListener(event int) {
	k.socket.RemoveListener(event)
}
