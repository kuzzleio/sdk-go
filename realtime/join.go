package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Join permits to join a previously created subscription
func (r *Realtime) Join(index, collection, roomID string, options types.RoomOptions, cb chan<- types.KuzzleNotification) error {
	if (index == "" || collection == "" || roomID == "") || (cb == nil) {
		return types.NewError("Realtime.Subscribe: index, collection, filters and notificationChannel required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		RoomId string `json:"roomId"`
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "join",
		Index:      index,
		Collection: collection,
		Body:       body{roomID},
	}

	opts := types.NewQueryOptions()

	if options != nil {
		query.Users = options.Users()
		query.State = options.State()
		query.Scope = options.Scope()

		opts.SetVolatile(options.Volatile())
	}

	go r.k.Query(query, opts, result)

	res := <-result
	if res.Error != nil {
		return res.Error
	}

	var resSub struct {
		RequestID string `json:"requestId"`
		RoomID    string `json:"roomId"`
		Channel   string `json:"channel"`
	}

	json.Unmarshal(res.Result, &resSub)

	onReconnect := make(chan interface{})

	r.k.RegisterSub(resSub.Channel, resSub.RoomID, nil, cb, onReconnect)

	go func() {
		<-onReconnect

		if r.k.AutoResubscribe() {
			go r.k.Query(query, opts, result)
		}
	}()

	r.k.AddListener(event.Reconnected, onReconnect)

	return nil
}
