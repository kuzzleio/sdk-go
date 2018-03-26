package realtime

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Join permits to join a previously created subscription
func (r *Realtime) Join(index, collection, roomID string, cb chan<- types.KuzzleNotification) error {
	return nil
}
