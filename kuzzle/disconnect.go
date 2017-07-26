package kuzzle

// Disconnect from Kuzzle and invalidate this instance.
// Does not fire a disconnected event.
func (k *Kuzzle) Disconnect() error {
	err := k.socket.Close()

	if err != nil {
		return err
	}
	k.wasConnected = false

	return nil
}
