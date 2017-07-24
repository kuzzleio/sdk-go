package security

import (
  "github.com/kuzzleio/sdk-go/kuzzle"
)

type Security struct {
  kuzzle *kuzzle.Kuzzle
}

func NewSecurity(kuzzle *kuzzle.Kuzzle) *Security {
  return &Security{kuzzle}
}
