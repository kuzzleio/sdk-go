package kuzzle

// AddListener Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func AddListener(k *Kuzzle, event int, channel chan<- interface{}) {
	k.socket.AddListener(event, channel)
}
