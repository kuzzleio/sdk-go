package internal

import "github.com/kuzzleio/sdk-go/types"

type MockedRoom struct {
	MockedRenew func()
}

func (m MockedRoom) Renew(filters interface{}, realtimeNotificationChannel chan<- types.KuzzleNotification, subscribeResponseChan chan<- types.SubscribeResponse) {
	m.MockedRenew()
}
func (m MockedRoom) Unsubscribe() {}

func (m MockedRoom) GetRealtimeChannel() chan<- types.KuzzleNotification {
	return make(chan<- types.KuzzleNotification)
}

func (m MockedRoom) GetResponseChannel() chan<- types.SubscribeResponse {
	return make(chan<- types.SubscribeResponse)
}

func (m MockedRoom) GetRoomId() string {
	return ""
}

func (m MockedRoom) GetFilters() interface{} {
	return nil
}
