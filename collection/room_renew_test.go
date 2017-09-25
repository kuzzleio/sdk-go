package collection

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRenewNotConnected(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	room := NewRoom(*NewCollection(k, "collection", "index"), nil)
	room.Renew(nil, nil, nil)

	assert.Equal(t, 1, len(room.pendingSubscriptions))
}

func TestRenewSubscribing(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	*k.State = state.Connected

	room := NewRoom(*NewCollection(k, "collection", "index"), nil)
	room.subscribing = true
	room.Renew(nil, nil, nil)

	assert.Equal(t, 1, room.queue.Len())
}

func TestRenewQueryError(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "ah!"}}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected

	subResChan := make(chan types.SubscribeResponse)
	NewRoom(*NewCollection(k, "collection", "index"), nil).Renew(nil, nil, subResChan)

	res := <-subResChan
	assert.Equal(t, "ah!", res.Error.Error())
}

func TestRenewWithSubscribeToSelf(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			room := NewRoom(*NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"
			(*room.collection.Kuzzle.RequestHistory)["ah!"] = time.Now()
			marshed, _ := json.Marshal(room)

			return types.KuzzleResponse{RequestId: "ah!", Result: marshed}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected

	subResChan := make(chan types.SubscribeResponse)
	NewRoom(*NewCollection(k, "collection", "index"), nil).Renew(nil, nil, subResChan)

	res := <-subResChan
	assert.Equal(t, "42", res.Room.GetRoomId())
}

func TestRenew(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			room := NewRoom(*NewCollection(k, "collection", "index"), nil)
			room.RoomId = "42"

			marshed, _ := json.Marshal(room)
			return types.KuzzleResponse{Result: marshed}
		},
	}
	k, _ = kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected

	NewRoom(*NewCollection(k, "collection", "index"), nil).Renew(nil, nil, nil)
}

func ExampleRoom_Renew() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	*k.State = state.Connected

	NewRoom(*NewCollection(k, "collection", "index"), nil).Renew(nil, nil, nil)
}
