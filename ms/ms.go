package ms

import "github.com/kuzzleio/sdk-go/kuzzle"

type Ms struct {
	kuzzle *kuzzle.Kuzzle
}

func NewMs(kuzzle *kuzzle.Kuzzle) *Ms {
	return &Ms{
		kuzzle: kuzzle,
	}
}
