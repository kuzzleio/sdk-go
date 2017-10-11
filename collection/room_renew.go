package collection

import (
	"container/list"
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"time"
)

// Renew the subscription. Force a resubscription using the same filters
// if no new ones are provided.
// Unsubscribes first if this Room was already listening to events.
func (room Room) Renew(filters interface{}, realtimeNotificationChannel chan<- *types.KuzzleNotification, subscribeResponseChan chan<- *types.SubscribeResponse) {
	if filters != nil {
		room.filters = filters
	}

	if room.collection.Kuzzle.State != state.Connected {
		room.RealtimeNotificationChannel = realtimeNotificationChannel
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
	room.RoomId = ""
	room.subscribing = true
	room.pendingSubscriptions[room.id] = realtimeNotificationChannel
	room.RealtimeNotificationChannel = realtimeNotificationChannel

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
			User:       room.user,
			Body:       filters,
		}, opts, result)

		res := <-result

		room.subscribing = false

		if res.Error != nil {
			room.queue.Init()
			if subscribeResponseChan != nil {
				subscribeResponseChan <- &types.SubscribeResponse{Error: errors.New(res.Error.Message)}
			}
			return
		}

		delete(room.pendingSubscriptions, room.id)

		resRoom := NewRoom(room.collection, nil)
		json.Unmarshal(res.Result, resRoom)

		room.RequestId = res.RequestId
		room.Channel = resRoom.Channel
		room.RoomId = resRoom.RoomId

		room.collection.Kuzzle.RegisterRoom(room.Channel, room.id, room)
		room.dequeue()

		if room.RequestId != "" && !room.collection.Kuzzle.RequestHistory[room.RequestId].IsZero() {
			if room.subscribeToSelf {
				if subscribeResponseChan != nil {
					subscribeResponseChan <- &types.SubscribeResponse{Room: room}
				}
			}
			delete(room.collection.Kuzzle.RequestHistory, room.RequestId)
		} else {
			if subscribeResponseChan != nil {
				subscribeResponseChan <- &types.SubscribeResponse{Room: room}
			}
		}
	}()

	return
}

func (room Room) dequeue() {
	if room.queue.Len() > 0 {
		for sub := room.queue.Front(); sub != nil; sub = sub.Next() {
			go sub.Value.(func(e *list.Element))(sub)
		}
	}
}
