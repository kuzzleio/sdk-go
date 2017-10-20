package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Count returns the number of other subscriptions on that room.
func (room Room) Count() (int, error) {
	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "count",
		Body: struct {
			RoomId string `json:"roomId"`
		}{room.RoomId},
	}

	result := make(chan *types.KuzzleResponse)

	go room.collection.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	count := struct {
		Count int `json:"count"`
	}{}

	json.Unmarshal(res.Result, &count)

	return count.Count, nil
}
