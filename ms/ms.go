// Package ms provides methods to interact with the Kuzzle memory storage
package ms

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

type Ms struct {
	Kuzzle *kuzzle.Kuzzle
}

// NewMs initializes a new Ms struct
func NewMs(kuzzle *kuzzle.Kuzzle) *Ms {
	return &Ms{
		Kuzzle: kuzzle,
	}
}
