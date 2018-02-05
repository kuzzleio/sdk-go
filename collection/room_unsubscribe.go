package collection

import (
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Unsubscribe from Kuzzle. Stop listening immediately.
// If there is no listener left on that Room, sends an unsubscribe request to Kuzzle, once
// pending subscriptions reaches 0, and only if there is still no listener on that Room.
// We wait for pending subscriptions to finish to avoid unsubscribing while another subscription on that Room is
// Unsubscribe from Kuzzle. Stop listening immediately.
func (room *Room) Unsubscribe() error {
	if room.internalState == subscribing {
		return types.NewError("Cannot unsubscribe a room while a subscription attempt is underway")
	}

	if room.isListening {
		room.collection.Kuzzle.RemoveListener(event.Disconnected, room.onDisconnect)
		room.collection.Kuzzle.RemoveListener(event.TokenExpired, room.onTokenExpired)
		room.collection.Kuzzle.RemoveListener(event.Reconnected, room.onReconnect)
		room.isListening = false
	}

	if room.internalState == active {

	}
	// if !room.isReady() {
	// 	room.queue.PushFront(func() {
	// 		room.Unsubscribe()
	// 	})
	// 	return
	// }

	// if room.roomId == "" {
	// 	return
	// }

	// room.collection.Kuzzle.UnregisterRoom(room.id)

	// type body struct {
	// 	RoomId string `json:"roomId"`
	// }

	// query := &types.KuzzleRequest{
	// 	Controller: "realtime",
	// 	Action:     "unsubscribe",
	// 	Body:       &body{room.roomId},
	// }

	// if len(room.pendingSubscriptions) > 0 {
	// 	time.Sleep(100 * time.Millisecond)
	// 	room.Unsubscribe()
	// }

	// room.collection.Kuzzle.Query(query, nil, nil)
	return nil
}
