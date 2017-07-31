package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type Collection struct {
	kuzzle            *kuzzle.Kuzzle
	index, collection string
	subscribeCallback interface{}
}

func NewCollection(kuzzle *kuzzle.Kuzzle, collection, index string) *Collection {
	return &Collection{
		index:      index,
		collection: collection,
		kuzzle:     kuzzle,
	}
}

func (dc Collection) CollectionDocument() CollectionDocument {
	return CollectionDocument{
		Collection: dc,
		Document: types.Document{Source: []byte(`{}`)},
	}
}
