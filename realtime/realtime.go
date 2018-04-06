package realtime

import "github.com/kuzzleio/sdk-go/types"

// Realtime is a realtime controller
type Realtime struct {
	k types.IKuzzle
}

// NewRealtime is a realtime controller constructor
func NewRealtime(k types.IKuzzle) *Realtime {
	return &Realtime{k}
}
