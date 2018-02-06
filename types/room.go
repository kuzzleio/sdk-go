package types

import "sync"

type SubscribeResponse struct {
	Room  IRoom
	Error error
}

type IRoom interface {
	Subscribe(realtimeNotificationChannel chan<- *KuzzleNotification)
	Unsubscribe() error
	RealtimeChannel() chan<- *KuzzleNotification
	ResponseChannel() chan *SubscribeResponse
	RoomId() string
	Filters() interface{}
	Channel() string
	Id() string
	SubscribeToSelf() bool
}

type RoomList = sync.Map
