package kuzzle

// ReplayQueue replays the requests queued during offline mode. Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
func (k Kuzzle) ReplayQueue() {
	k.socket.ReplayQueue()
}
