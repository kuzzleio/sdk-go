package kuzzle

// RemoveListener removes a listener
func (k Kuzzle) RemoveListener(event int, channel chan<- interface{}) {
	k.socket.RemoveListener(event, channel)
}
