package types

import "sync"

type IRoom interface {
	Renew(filters interface{}, realtimeNotificationChannel chan<- *KuzzleNotification, subscribeResponseChan chan<- *SubscribeResponse)
	Unsubscribe()
	RealtimeChannel() chan<- *KuzzleNotification
	ResponseChannel() chan<- *SubscribeResponse
	RoomId() string
	Filters() interface{}
}

type SubscribeResponse struct {
	Room  IRoom
	Error error
}

type RoomList = sync.Map
