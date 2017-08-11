package collection

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

type Collection struct {
	Kuzzle            *kuzzle.Kuzzle
	index, collection string
	subscribeCallback interface{}
}

func NewCollection(kuzzle *kuzzle.Kuzzle, collection, index string) *Collection {
	return &Collection{
		index:             index,
		collection:        collection,
		Kuzzle:            kuzzle,
	}
}

func (dc Collection) Document() Document {
	return Document{
		Content:    []byte(`{}`),
		collection: dc,
	}
}
