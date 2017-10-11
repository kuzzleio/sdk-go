package collection

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubscribeError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.State = state.Connected

	subRes := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-subRes
	assert.Equal(t, "error", r.Error.Error())
}

func TestSubscribe(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			room := NewRoom(NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)
			return &types.KuzzleResponse{Result: marshed}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	k.State = state.Connected

	subRes := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-subRes
	assert.Equal(t, "42", r.Room.GetRoomId())
}

func ExampleCollection_Subscribe() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.State = state.Connected

	subRes := NewCollection(k, "collection", "index").Subscribe(nil, nil, nil)

	r := <-subRes

	fmt.Println(r.Room.GetRoomId())
}
