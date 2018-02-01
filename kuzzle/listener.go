package kuzzle

// AddListener Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (k *Kuzzle) AddListener(event int, channel chan<- interface{}) {
	k.socket.AddListener(event, channel)
}

// On is an alias to the AddListener function
func (k *Kuzzle) On(event int, channel chan<- interface{}) {
	k.socket.AddListener(event, channel)
}

// Remove all listener by event type or all listener if event == -1
func (k *Kuzzle) RemoveAllListeners(event int) {
	k.socket.RemoveAllListeners(event)
}

// RemoveListener removes a listener
func (k *Kuzzle) RemoveListener(event int, channel chan<- interface{}) {
	k.socket.RemoveListener(event, channel)
}

func (k *Kuzzle) Once(event int, channel chan<- interface{}) {
	k.socket.Once(event, channel)
}

func (k *Kuzzle) ListenerCount(event int) int {
	return k.socket.ListenerCount(event)
}
