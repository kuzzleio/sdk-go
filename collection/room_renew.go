package collection

import (
	"container/list"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

// Renew the subscription. Force a resubscription using the same filters
// if no new ones are provided.
// Unsubscribes first if this Room was already listening to events.
func (room *Room) Renew(filters interface{}, realtimeNotificationChannel chan<- *types.KuzzleNotification, subscribeResponseChan chan<- *types.SubscribeResponse) {
	if filters != nil {
		room.filters = filters
	}

	if room.collection.Kuzzle.State() != state.Connected {
		room.realtimeNotificationChannel = realtimeNotificationChannel
		room.pendingSubscriptions[room.id] = realtimeNotificationChannel
		return
	}

	if room.subscribing {
		room.queue.PushFront(func(e *list.Element) {
			room.Renew(filters, realtimeNotificationChannel, subscribeResponseChan)
			room.queue.Remove(e)
			time.Sleep(1 * time.Second)
		})
		return
	}

	room.Unsubscribe()
	room.roomId = ""
	room.subscribing = true
	room.pendingSubscriptions[room.id] = realtimeNotificationChannel
	room.realtimeNotificationChannel = realtimeNotificationChannel

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
			Body:       filters,
		}, opts, result)

		res := <-result
		room.subscribing = false

		if res.Error != nil {
			room.queue.Init()
			if subscribeResponseChan != nil {
				subscribeResponseChan <- &types.SubscribeResponse{Error: res.Error}
			}
			return
		}

		delete(room.pendingSubscriptions, room.id)

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

		room.collection.Kuzzle.RegisterRoom(room.channel, room.id, room)
		room.dequeue()

		if room.requestId != "" && !room.collection.Kuzzle.RequestHistory[room.requestId].IsZero() {
			if room.subscribeToSelf {
				if subscribeResponseChan != nil {
					subscribeResponseChan <- &types.SubscribeResponse{Room: room}
				}
			}
			delete(room.collection.Kuzzle.RequestHistory, room.requestId)
		} else {
			if subscribeResponseChan != nil {
				subscribeResponseChan <- &types.SubscribeResponse{Room: room}
			}
		}
	}()

	return
}

func (room *Room) dequeue() {
	if room.queue.Len() > 0 {
		for sub := room.queue.Front(); sub != nil; sub = sub.Next() {
			go sub.Value.(func(e *list.Element))(sub)
		}
	}
}
