package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// Subscribe permits to join a previously created subscription
func (r *Realtime) Subscribe(index, collection string, filters json.RawMessage, cb chan<- types.KuzzleNotification, options types.RoomOptions) (string, error) {
	if (index == "" || collection == "") || (filters == nil || cb == nil) {
		return "", types.NewError("Realtime.Subscribe: index, collection, filters and notificationChannel required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "subscribe",
		Index:      index,
		Collection: collection,
		Body:       filters,
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
		return "", res.Error
	}

	var resSub struct {
		RequestID string `json:"requestId"`
		RoomID    string `json:"roomId"`
		Channel   string `json:"channel"`
	}

	json.Unmarshal(res.Result, &resSub)

	onReconnect := make(chan interface{})

	r.k.RegisterSub(resSub.Channel, resSub.RoomID, filters, cb, onReconnect)

	go func() {
		<-onReconnect

		if r.k.AutoResubscribe() {
			go r.k.Query(query, opts, result)
		}
	}()

	r.k.AddListener(event.Reconnected, onReconnect)

	return resSub.RoomID, nil
}
