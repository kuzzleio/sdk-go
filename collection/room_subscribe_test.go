package collection

import (
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeQueryError(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "ah!"}}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	subResChan := make(chan *types.SubscribeResponse)
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.OnDone(subResChan)
	r.Subscribe(nil)

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
	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.OnDone(subResChan)
	r.Subscribe(nil)

	res := <-subResChan
	assert.Equal(t, "42", res.Room.RoomId())
}

func TestRoomSubscribe(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	NewRoom(NewCollection(k, "collection", "index"), nil, nil).Subscribe(nil)
}

func TestRoomSubscribeAlreadyActive(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			roomRaw := []byte(`{"requestId": "rqid", "channel": "foo", "roomId": "42"}`)
			return &types.KuzzleResponse{Result: roomRaw}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.internalState = active
	r.subscribeResponseChan = make(chan *types.SubscribeResponse)
	done := make(chan bool)

	go func() {
		<-r.subscribeResponseChan
		done <- true
	}()
	r.Subscribe(nil)

	<-done
}

func ExampleRoom_Subscribe() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	c.SetState(state.Connected)

	NewRoom(NewCollection(k, "collection", "index"), nil, nil).Subscribe(nil)
}
