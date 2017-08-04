package kuzzle

/*
 * Unset the authentication token and cancel all subscriptions
 */
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	for _, rooms := range k.socket.GetRooms() {
		for _, room := range rooms {
			room.Renew(room.GetFilters(), room.GetRealtimeChannel(), room.GetResponseChannel())
		}
	}
}
