package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Subscribe permits to join a previously created subscription
func (r *Realtime) Subscribe(index, collection, body string, cb chan<- types.KuzzleNotification, options types.RoomOptions) (string, error) {
	if (index == "" || collection == "") || (body == "" || cb == nil) {
		return "", types.NewError("Realtime.Subscribe: index, collection, body and notificationChannel required", 400)
	}

	if r.k.InternalState() == internalState.active {
		if room.SubscribeResponseChan != nil {
			room.SubscribeResponseChan <- types.SubscribeResponse{Room: room}
		}
		return room.roomId, nil
	}

	if room.internalState == subscribing {
		return room.id, nil
	}

	queryResult := make(chan types.SubscribeResponse)

	go func() {
		result := make(chan *types.KuzzleResponse)

		query := &types.KuzzleRequest{
			Controller: "realtime",
			Action:     "subscribe",
			Index:      index,
			Collection: collection,
			Body:       body,
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
			queryResult <- types.SubscribeResponse{Room: "", Error: res.Error}
			return
		}

		var resSub struct {
			RequestID string `json:"requestId"`
			RoomID    string `json:"roomId"`
			Channel   string `json:"channel"`
		}

		json.Unmarshal(res.Result, &resSub)
		queryResult <- types.SubscribeResponse{Room: resSub.RoomID, Error: res.Error}

		r.k.RegisterSub(resSub.Channel, resSub.RoomID, cb)

		go func() {
			<-room.OnDisconnect
			r.k.SetInternalState(resSub.Channel, resSub.RoomID, inactive)

		}()
		go func() {
			<-room.OnTokenExpired
		}()
		go func() {
			<-room.OnReconnect

			if r.k.AutoResubscribe() {
				r.Subscribe(index, collection, body, cb, options)
			}
		}()

		r.k.AddListener(event.Disconnected, cb)
		r.k.AddListener(event.TokenExpired, cb)
		r.k.AddListener(event.Reconnected, cb)
	}()

	res := <-queryResult

	return res.Result, res.Error
}
