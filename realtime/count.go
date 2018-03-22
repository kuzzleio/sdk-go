package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//Count returns the number of other subscriptions on that room.
func (r *Realtime) Count(index, collection, roomID string) (int, error) {
	if index == "" || collection == "" || roomID == "" {
		return -1, types.NewError("Realtime.Count: index, collection and roomID required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "count",
		Index:      index,
		Collection: collection,
		Body: struct {
			RoomID string `json:"roomID"`
		}{roomID},
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	var countRes struct {
		Count int `json:"count"`
	}

	json.Unmarshal(res.Result, &countRes)

	return countRes.Count, nil
}
