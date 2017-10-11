package kuzzle

// FlushQueue empties the offline queue without replaying it.
func (k Kuzzle) FlushQueue() {
  *k.GetOfflineQueue() = (*k.GetOfflineQueue())[:0]
}
