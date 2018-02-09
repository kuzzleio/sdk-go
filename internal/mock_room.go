package internal

import "github.com/kuzzleio/sdk-go/types"

type MockedRoom struct {
	MockedSubscribe func()
}

func (m MockedRoom) Subscribe(realtimeNotificationChannel chan<- *types.KuzzleNotification) {
	m.MockedSubscribe()
}
func (m MockedRoom) Unsubscribe() error {
	return nil
}

func (m MockedRoom) GetRealtimeChannel() chan<- *types.KuzzleNotification {
	return make(chan<- *types.KuzzleNotification)
}

func (m MockedRoom) GetResponseChannel() chan<- *types.SubscribeResponse {
	return make(chan<- *types.SubscribeResponse)
}

func (m MockedRoom) RoomId() string {
	return ""
}

func (m MockedRoom) Filters() interface{} {
	return nil
}

func (m MockedRoom) Channel() string {
	return ""
}

func (m MockedRoom) Id() string {
	return ""
}

func (m MockedRoom) SubscribeToSelf() bool {
	return true
}
