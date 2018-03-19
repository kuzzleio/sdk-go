package collection

import (
	"github.com/kuzzleio/sdk-go/types"
)

//Collection controller
type Collection struct {
	Kuzzle types.IKuzzle
}

// NewCollection instanciates a new collection
func NewCollection(kuzzle types.IKuzzle) *Collection {
	return &Collection{
		Kuzzle: kuzzle,
	}
}
