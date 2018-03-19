package document

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Document is a Controller
type Document struct {
	Kuzzle types.IKuzzle
}

// NewDocument is a Document Controller constructor
func NewDocument(kuzzle types.IKuzzle) *Document {
	return &Document{
		Kuzzle: kuzzle,
	}
}
