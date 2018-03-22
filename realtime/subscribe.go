package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Subscribe permits to join a previously created subscription
func (r *Realtime) Subscribe(index, collection, body string, cb chan<- types.KuzzleNotification, options types.RoomOptions) (string, error) {
	if (index == "" || collection == "") || (body == "" || cb == nil) {
		return "", types.NewError("Realtime.Subscribe: index, collection, body and notificationChannel required", 400)
	}

	type SubResult struct {
		Result string
		Error  error
	}

	result := make(chan *types.KuzzleResponse)
	queryResult := make(chan SubResult)

	go func() {
		query := &types.KuzzleRequest{
			Controller: "realtime",
			Action:     "subscribe",
			Index:      index,
			Collection: collection,
			Body:       body,
		}

		if options != nil {
			query.Volatile = options.Volatile()
			query.Users = options.Users()
			query.State = options.State()
			query.Scope = options.Scope()
		}

		go r.k.Query(query, nil, result)

		res := <-result
		if res.Error != nil {
			queryResult <- SubResult{Result: "", Error: res.Error}
			return
		}

		var resSub struct {
			RoomID  string `json:"roomId"`
			Channel string `json:"channel"`
		}

		json.Unmarshal(res.Result, &resSub)
		queryResult <- SubResult{Result: resSub.RoomID, Error: res.Error}

		r.k.RegisterSub(resSub.Channel, resSub.RoomID, cb)
		return
	}()

	res := <-queryResult

	return res.Result, res.Error
}
