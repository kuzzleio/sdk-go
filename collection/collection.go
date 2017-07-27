package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

type Collection struct {
	kuzzle            *kuzzle.Kuzzle
	index, collection string
	subscribeCallback interface{}
	collectionMapping CollectionMapping
}

func NewCollection(kuzzle *kuzzle.Kuzzle, collection, index string) *Collection {
	return &Collection{
		index:             index,
		collection:        collection,
		kuzzle:            kuzzle,
		collectionMapping: CollectionMapping{},
	}
}
