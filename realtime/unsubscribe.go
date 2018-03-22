package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Unsubscribe instructs Kuzzle to detach you from its subscribers for the given room
func (r *Realtime) Unsubscribe(roomID string) error {
	if roomID == "" {
		return types.NewError("Realtime.Unsubscribe: roomID required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "unsubscribe",
		Body: struct {
			RoomID string `json:"roomID"`
		}{roomID},
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return res.Error
	}

	var oldRoomID string
	json.Unmarshal(res.Result, &oldRoomID)

	r.k.UnregisterSub(oldRoomID)

	return nil
}
