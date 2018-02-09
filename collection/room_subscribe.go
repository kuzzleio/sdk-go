package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/event"

	"github.com/kuzzleio/sdk-go/types"
)

// Renew the subscription. Force a resubscription using the same filters
// if no new ones are provided.
// Unsubscribes first if this Room was already listening to events.
func (room *Room) Subscribe(realtimeNotificationChannel chan<- types.KuzzleNotification) {
	if room.internalState == active {
		if room.subscribeResponseChan != nil {
			room.subscribeResponseChan <- types.SubscribeResponse{Room: room}
		}
		return
	}

	if room.internalState == subscribing {
		return
	}

	room.err = nil
	room.internalState = subscribing

	go func() {
		result := make(chan *types.KuzzleResponse)

		opts := types.NewQueryOptions()
		opts.SetVolatile(room.Volatile)

		go room.collection.Kuzzle.Query(&types.KuzzleRequest{
			Controller: "realtime",
			Action:     "subscribe",
			Index:      room.collection.index,
			Collection: room.collection.collection,
			Scope:      room.scope,
			State:      room.state,
			Users:      room.users,
			Body:       room.filters,
		}, opts, result)

		res := <-result
		room.subscribing = false
		room.internalState = subscribing

		if res.Error != nil {
			if res.Error.Message == "Not Connected" {
				c := make(chan interface{})
				room.Once(event.Connected, c)
				go func() {
					<-c
					room.internalState = inactive
					room.err = nil
					room.Subscribe(realtimeNotificationChannel)
				}()
				return
			}
			room.internalState = inactive
			if room.subscribeResponseChan != nil {
				room.subscribeResponseChan <- types.SubscribeResponse{Error: res.Error}
			}
			return
		}

		type RoomResult struct {
			RequestId string `json:"requestId"`
			RoomId    string `json:"roomId"`
			Channel   string `json:"channel"`
		}

		var resRoom RoomResult
		json.Unmarshal(res.Result, &resRoom)

		room.requestId = resRoom.RequestId
		room.channel = resRoom.Channel
		room.roomId = resRoom.RoomId
		room.internalState = active

		room.realtimeNotificationChannel = realtimeNotificationChannel
		room.collection.Kuzzle.RegisterRoom(room)

		if room.subscribeResponseChan != nil {
			room.subscribeResponseChan <- types.SubscribeResponse{Room: room}
		}
		if room.isListening {
			go func() {
				<-room.onDisconnect
				room.internalState = inactive
			}()
			go func() {
				<-room.onTokenExpired
				room.internalState = inactive
			}()
			go func() {
				<-room.onReconnect
				room.internalState = inactive

				if room.autoResubscribe {
					room.Subscribe(realtimeNotificationChannel)
				}
			}()

			room.AddListener(event.Disconnected, room.onDisconnect)
			room.AddListener(event.TokenExpired, room.onTokenExpired)
			room.AddListener(event.Reconnected, room.onReconnect)
			room.isListening = true
		}
	}()
}
