package index

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Index struct {
	kuzzle types.IKuzzle
}

func NewIndex(k types.IKuzzle) *Index {
	return &Index{k}
}
