package kuzzle

func (k Kuzzle) StopQueuing() {
	k.socket.StopQueuing()
}
