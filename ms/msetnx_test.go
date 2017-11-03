package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsetnxEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	_, err := memoryStorage.Msetnx([]*types.MSKeyValue{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Msetnx: please provide at least one key/value entry", fmt.Sprint(err))
}

func TestMsetnxError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	_, err := memoryStorage.Msetnx(entries, nil)

	assert.NotNil(t, err)
}

func TestMsetnx(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "msetnx", parsedQuery.Action)

			r, _ := json.Marshal(1)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	res, _ := memoryStorage.Msetnx(entries, nil)

	assert.Equal(t, 1, res)
}

func ExampleMs_Msetnx() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	entries := []*types.MSKeyValue{
		{Key: "foo", Value: "bar"},
		{Key: "bar", Value: "foo"},
	}

	res, err := memoryStorage.Msetnx(entries, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
