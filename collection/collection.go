package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

type Collection struct {
	Kuzzle *kuzzle.Kuzzle
}

// NewCollection instanciates a new collection
func NewCollection(kuzzle *kuzzle.Kuzzle) *Collection {
	return &Collection{
		Kuzzle: kuzzle,
	}
}

type ListOptions struct {
	Type string
	From int
	Size int
}

func NewListOptions(t string, from int, size int) *ListOptions {
	return &ListOptions{
		Type: t,
		From: from,
		Size: size,
	}
}
