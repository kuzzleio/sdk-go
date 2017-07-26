package collection_test

import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRoomWithoutOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	room := collection.NewCollection(k, "collection", "index").Room(nil)

	newRoom := types.Room{}

	assert.Equal(t, newRoom, room)
}

func TestInitRoomWithDefaultOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	roomOptions := types.DefaultRoomOptions()
	room := collection.NewCollection(k, "collection", "index").Room(roomOptions)

	assert.Equal(t, types.SCOPE_ALL, room.Scope)
	assert.Equal(t, types.STATE_DONE, room.State)
	assert.Equal(t, types.USER_NONE, room.User)
	assert.Equal(t, true, room.SubscribeToSelf)
}

func TestInitRoomWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)

	roomOptions := &types.RoomOptions{
		Scope:           types.SCOPE_OUT,
		State:           types.STATE_PENDING,
		User:            types.USER_ALL,
		SubscribeToSelf: false,
	}
	room := collection.NewCollection(k, "collection", "index").Room(roomOptions)

	assert.Equal(t, types.SCOPE_OUT, room.Scope)
	assert.Equal(t, types.STATE_PENDING, room.State)
	assert.Equal(t, types.USER_ALL, room.User)
	assert.Equal(t, false, room.SubscribeToSelf)
}
