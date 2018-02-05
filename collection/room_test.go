package collection

import (
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestInitRoomWithoutOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	assert.NotNil(t, r)
}

func TestRoomFilters(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}

	k, _ = kuzzle.NewKuzzle(c, nil)
	cl := NewCollection(k, "collection", "index")

	type SubscribeFiltersValues struct {
		Values []string `json:"values"`
	}
	type SubscribeFilters struct {
		Ids SubscribeFiltersValues `json:"ids"`
	}
	var filters = SubscribeFilters{
		Ids: SubscribeFiltersValues{
			Values: []string{"1"},
		},
	}

	c.SetState(state.Connected)
	rtc := make(chan *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, filters, res.Room.Filters())
}

func ExampleRoom_Filters() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := NewCollection(k, "collection", "index")

	type SubscribeFiltersValues struct {
		Values []string `json:"values"`
	}
	type SubscribeFilters struct {
		Ids SubscribeFiltersValues `json:"ids"`
	}
	var filters = SubscribeFilters{
		Ids: SubscribeFiltersValues{
			Values: []string{"1"},
		},
	}

	c.SetState(state.Connected)
	rtc := make(chan *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	fmt.Println(res)
}

func TestRoomRealtimeChannel(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}

	k, _ = kuzzle.NewKuzzle(c, nil)
	cl := NewCollection(k, "collection", "index")

	type SubscribeFiltersValues struct {
		Values []string `json:"values"`
	}
	type SubscribeFilters struct {
		Ids SubscribeFiltersValues `json:"ids"`
	}
	var filters = SubscribeFilters{
		Ids: SubscribeFiltersValues{
			Values: []string{"1"},
		},
	}

	c.SetState(state.Connected)
	rtc := make(chan<- *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, rtc, res.Room.RealtimeChannel())
}

func ExampleRoom_RealtimeChannel() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := NewCollection(k, "collection", "index")

	type SubscribeFiltersValues struct {
		Values []string `json:"values"`
	}
	type SubscribeFilters struct {
		Ids SubscribeFiltersValues `json:"ids"`
	}
	var filters = SubscribeFilters{
		Ids: SubscribeFiltersValues{
			Values: []string{"1"},
		},
	}

	c.SetState(state.Connected)
	rtc := make(chan<- *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)
	rtChannel := res.Room.RealtimeChannel()

	fmt.Println(rtChannel)
}

func TestAddListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockAddListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	ch := make(chan interface{})

	r.AddListener(0, ch)
	assert.Equal(t, true, called)
}

func TestRemoveListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockRemoveListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	ch := make(chan interface{})

	r.RemoveListener(0, ch)
	assert.Equal(t, true, called)
}
func TestRemoveAllListener(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockRemoveAllListeners: func(e int) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	r.RemoveAllListeners(0)
	assert.Equal(t, true, called)
}

func TestOnce(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockOnce: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	ch := make(chan interface{})

	r.Once(0, ch)
	assert.Equal(t, true, called)
}

func TestOn(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockAddListener: func(e int, c chan<- interface{}) {
			called = true
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	ch := make(chan interface{})

	r.On(0, ch)
	assert.Equal(t, true, called)
}

func TestListenerCount(t *testing.T) {
	called := false

	c := &internal.MockedConnection{
		MockListenerCount: func(e int) int {
			called = true
			return -1
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)

	r.ListenerCount(0)
	assert.Equal(t, true, called)
}
func TestOnDoneError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return nil
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	ch := make(chan *types.SubscribeResponse)
	done := make(chan bool)

	go func() {
		<-ch
		done <- true
	}()
	r.err = &types.KuzzleError{Message: "Room has an error"}
	r.OnDone(ch)

	<-done
}

func TestOnDoneAlreadyActive(t *testing.T) {
	c := &internal.MockedConnection{}

	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	ch := make(chan *types.SubscribeResponse)
	done := make(chan bool)

	go func() {
		<-ch
		done <- true
	}()

	r.internalState = active
	r.OnDone(ch)

	<-done
}
func TestOnDone(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	ch := make(chan *types.SubscribeResponse)
	r.OnDone(ch)

	r.Subscribe(nil)
	assert.Equal(t, ch, r.subscribeResponseChan)
}
