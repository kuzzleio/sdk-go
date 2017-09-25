package kuzzle

func (k Kuzzle) StartQueuing() {
	k.socket.StartQueuing()
}
