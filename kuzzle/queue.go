package kuzzle

// FlushQueue empties the offline queue without replaying it.
func (k *Kuzzle) FlushQueue() {
  k.socket.ClearQueue()
}

// ReplayQueue replays the requests queued during offline mode. Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
func (k *Kuzzle) ReplayQueue() {
  k.socket.ReplayQueue()
}

func (k *Kuzzle) StartQueuing() {
  k.socket.StartQueuing()
}

func (k *Kuzzle) StopQueuing() {
  k.socket.StopQueuing()
}
