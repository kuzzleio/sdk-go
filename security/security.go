package security

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Security struct {
	Kuzzle types.IKuzzle
}

// NewSecurity returns a new instance of Security
func New(kuzzle types.IKuzzle) *Security {
	return &Security{
		Kuzzle: kuzzle,
	}
}
