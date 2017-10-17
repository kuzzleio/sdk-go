// Package ms provides methods to interact with the Kuzzle memory storage
package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Ms struct {
	Kuzzle types.IKuzzle
}

// NewMs initializes a new Ms struct
func NewMs(kuzzle types.IKuzzle) *Ms {
	return &Ms{
		Kuzzle: kuzzle,
	}
}
