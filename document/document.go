package document

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Document struct {
	Kuzzle types.IKuzzle
}

func NewDocument(kuzzle types.IKuzzle) *Document {
	return &Document{
		Kuzzle: kuzzle,
	}
}

type DocumentOptions struct {
	Volatile string
	WaitFor  bool
}
