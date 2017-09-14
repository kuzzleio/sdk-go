package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestSscanEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 0
	_, err := memoryStorage.Sscan("", &cursor, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Sscan: key required", fmt.Sprint(err))
}

func TestSscanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 0
	_, err := memoryStorage.Sscan("foo", &cursor, qo)

	assert.NotNil(t, err)
}

func TestSscan(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sscan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 10
	res, _ := memoryStorage.Sscan("foo", &cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func TestSscanWithOptions(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sscan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetCount(42)
	qo.SetMatch("*")

	cursor := 10
	res, _ := memoryStorage.Sscan("foo", &cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func ExampleMs_Sscan() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()


	qo.SetCount(42)
	qo.SetMatch("*")

	cursor := 10
	res, err := memoryStorage.Sscan("foo", &cursor, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Cursor, res.Values, cursor)
}
