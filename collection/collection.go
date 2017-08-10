package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type Collection struct {
	Kuzzle            *kuzzle.Kuzzle
	index, collection string
	subscribeCallback interface{}
	collectionMapping CollectionMapping
}

func NewCollection(kuzzle *kuzzle.Kuzzle, collection, index string) *Collection {
	return &Collection{
		index:             index,
		collection:        collection,
		Kuzzle:            kuzzle,
		collectionMapping: CollectionMapping{},
	}
}

func (dc Collection) CollectionDocument() CollectionDocument {
	return CollectionDocument{
		Collection: dc,
		Document:   types.Document{Content: []byte(`{}`)},
	}
}
