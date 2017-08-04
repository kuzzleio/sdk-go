package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/**
 * Returns the number of other subscriptions on that room.
 */
func (room *Room) Count() (int, error) {
	query := types.KuzzleRequest{
		Controller: "realtime",
		Action:     "count",
		Body: struct {
			RoomId string `json:"roomId"`
		}{room.RoomId},
	}

	result := make(chan types.KuzzleResponse)

	go room.collection.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}

	count := struct {
		Count int `json:"count"`
	}{}

	json.Unmarshal(res.Result, &count)

	return count.Count, nil
}
