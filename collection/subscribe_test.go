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

func TestSubscribeError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	room := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-room.ResponseChannel()
	assert.Equal(t, "error", r.Error.Error())
}

func TestSubscribe(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	room := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-room.ResponseChannel()
	assert.Equal(t, "42", r.Room.RoomId())
}

func ExampleCollection_Subscribe() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	room := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-room.ResponseChannel()

	fmt.Println(r.Room.RoomId())
}
