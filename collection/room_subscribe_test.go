package collection

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenewNotConnected(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	room := NewRoom(NewCollection(k, "collection", "index"), nil)
	room.Renew(nil, nil, nil)

	assert.Equal(t, 1, len(room.pendingSubscriptions))
}

func TestRenewSubscribing(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	room := NewRoom(NewCollection(k, "collection", "index"), nil)
	room.subscribing = true
	room.Renew(nil, nil, nil)

	assert.Equal(t, 1, room.queue.Len())
}

func TestRenewQueryError(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "ah!"}}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	subResChan := make(chan *types.SubscribeResponse)
	NewRoom(NewCollection(k, "collection", "index"), nil).Renew(nil, nil, subResChan)

	res := <-subResChan
	assert.Equal(t, "ah!", res.Error.Error())
}

func TestRenewWithSubscribeToSelf(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{RequestId: "ah!", Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	subResChan := make(chan *types.SubscribeResponse)
	NewRoom(NewCollection(k, "collection", "index"), nil).Renew(nil, nil, subResChan)

	res := <-subResChan
	assert.Equal(t, "42", res.Room.RoomId())
}

func TestRenew(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	NewRoom(NewCollection(k, "collection", "index"), nil).Renew(nil, nil, nil)
}

func ExampleRoom_Renew() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	NewRoom(NewCollection(k, "collection", "index"), nil).Renew(nil, nil, nil)
}
