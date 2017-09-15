package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRoomWithoutOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	r := collection.NewRoom(*collection.NewCollection(k, "collection", "index"), nil)

	assert.NotNil(t, r)
}

func TestRoomGetFilters(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			room := collection.NewRoom(*collection.NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)
			return types.KuzzleResponse{Result: marshed}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")

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

	*k.State = state.Connected
	rtc := make(chan types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, filters, res.Room.GetFilters())
}

func ExampleRoom_GetFilters() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")

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

	*k.State = state.Connected
	rtc := make(chan types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	fmt.Println(res)
}

func TestRoomGetRealtimeChannel(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			room := collection.NewRoom(*collection.NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)
			return types.KuzzleResponse{Result: marshed}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")

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

	*k.State = state.Connected
	rtc := make(chan<- types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, rtc, res.Room.GetRealtimeChannel())
}

func ExampleRoom_GetRealtimeChannel() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")

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

	*k.State = state.Connected
	rtc := make(chan<- types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)
	rtChannel := res.Room.GetRealtimeChannel()

	fmt.Println(rtChannel)
}
