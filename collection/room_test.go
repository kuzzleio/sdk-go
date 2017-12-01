package collection_test

import (
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRoomWithoutOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	r := collection.NewRoom(collection.NewCollection(k, "collection", "index"), nil)

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

	c.SetState(state.Connected)
	rtc := make(chan *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, filters, res.Room.Filters())
}

func ExampleRoom_Filters() {
	c := &internal.MockedConnection{}
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

	c.SetState(state.Connected)
	rtc := make(chan<- *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)

	assert.Equal(t, rtc, res.Room.RealtimeChannel())
}

func ExampleRoom_RealtimeChannel() {
	c := &internal.MockedConnection{}
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

	c.SetState(state.Connected)
	rtc := make(chan<- *types.KuzzleNotification)
	res := <-cl.Subscribe(filters, types.NewRoomOptions(), rtc)
	rtChannel := res.Room.RealtimeChannel()

	fmt.Println(rtChannel)
}
