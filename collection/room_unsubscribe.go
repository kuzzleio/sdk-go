package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

// Unsubscribe from Kuzzle. Stop listening immediately.
// If there is no listener left on that Room, sends an unsubscribe request to Kuzzle, once
// pending subscriptions reaches 0, and only if there is still no listener on that Room.
// We wait for pending subscriptions to finish to avoid unsubscribing while another subscription on that Room is
func (room *Room) Unsubscribe() {
	if !room.isReady() {
		room.queue.PushFront(func() {
			room.Unsubscribe()
		})
		return
	}

	if room.RoomId == "" {
		return
	}

	room.collection.Kuzzle.UnregisterRoom(room.id)

	type body struct {
		RoomId string `json:"roomId"`
	}

	query := types.KuzzleRequest{
		Controller: "realtime",
		Action:     "unsubscribe",
		Body:       body{room.RoomId},
	}

	if len(room.pendingSubscriptions) > 0 {
		time.Sleep(100 * time.Millisecond)
		room.Unsubscribe()
	}

	room.collection.Kuzzle.Query(query, nil, nil)
}
