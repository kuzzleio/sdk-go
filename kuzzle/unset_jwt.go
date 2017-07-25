package kuzzle

/*
 * Unset the authentication token and cancel all subscriptions
 */
func (k *Kuzzle) UnsetJwt() {
  k.jwt = ""

  //todo unsbuscribe every room
}
