package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//Collection controller
type Collection struct {
	Kuzzle *kuzzle.Kuzzle
}

// NewCollection instanciates a new collection
func NewCollection(kuzzle *kuzzle.Kuzzle) *Collection {
	return &Collection{
		Kuzzle: kuzzle,
	}
}

//ListOptions collection list options
type ListOptions struct {
	Type string
	From *int
	Size *int
}
