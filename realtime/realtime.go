package realtime

import "github.com/kuzzleio/sdk-go/types"

type Realtime struct {
	k types.IKuzzle
}

func NewRealtime(k types.IKuzzle) *Realtime {
	return &Realtime{k}
}
