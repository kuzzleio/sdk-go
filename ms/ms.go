package ms

import "github.com/kuzzleio/sdk-go/kuzzle"

type Ms struct {
	Kuzzle *kuzzle.Kuzzle
}

func NewMs(kuzzle *kuzzle.Kuzzle) *Ms {
	return &Ms{
		Kuzzle: kuzzle,
	}
}
