package kuzzle

func (k *Kuzzle) UnsetJwt() {
  k.jwt = ""

  //todo unsbuscribe every room
}
