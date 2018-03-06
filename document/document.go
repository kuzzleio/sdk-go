package document

import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

type Document struct {
	Kuzzle *kuzzle.Kuzzle
}

func NewDocument(kuzzle *kuzzle.Kuzzle) *Document {
	return &Document{
		Kuzzle: kuzzle,
	}
}

type DocumentOptions struct {
	Volatile string
	WaitFor  bool
}
