package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//Count returns the number of other subscriptions on that room.
func (r *Realtime) Count(index, collection, roomId string) (int, error) {
	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "count",
		Index:      index,
		Collection: collection,
		Body: struct {
			RoomId string `json:"roomId"`
		}{roomId},
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return -1, res.Error
	}

	type countRes struct {
		Count int `json:"count"`
	}
	c := &countRes{}
	json.Unmarshal(res.Result, c)

	return c.Count, nil
}
