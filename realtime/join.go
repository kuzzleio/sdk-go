package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Join permits to join a previously created subscription
func (r *Realtime) Join(index, collection, roomID string, cb chan<- types.KuzzleNotification) error {
	if (index == "" || collection == "") || roomID == "" {
		return types.NewError("Realtime.Join: index, collection and roomID required", 400)
	}

	type SubResult struct {
		Error error
	}

	result := make(chan *types.KuzzleResponse)
	queryResult := make(chan SubResult)

	go func() {
		query := &types.KuzzleRequest{
			Controller: "realtime",
			Action:     "join",
			Index:      index,
			Collection: collection,
			Body: struct {
				RoomID string `json:"roomID"`
			}{roomID},
		}

		go r.k.Query(query, nil, result)

		res := <-result
		if res.Error != nil {
			queryResult <- SubResult{Error: res.Error}
			return
		}

		var resSub struct {
			RoomID  string `json:"roomId"`
			Channel string `json:"channel"`
		}

		json.Unmarshal(res.Result, &resSub)
		queryResult <- SubResult{Error: res.Error}

		r.k.RegisterSub(resSub.Channel, resSub.RoomID, cb)
		return
	}()

	res := <-queryResult

	return res.Error
}
