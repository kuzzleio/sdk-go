package ms_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestScanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 0
	_, err := memoryStorage.Scan(&cursor, qo)

	assert.NotNil(t, err)
}

func TestScan(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "scan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 0
	res, _ := memoryStorage.Scan(&cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func TestScanWithOptions(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "scan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetCount(42)
	qo.SetMatch("*")

	cursor := 0
	res, _ := memoryStorage.Scan(&cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func ExampleMs_Scan() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetCount(42)
	qo.SetMatch("*")

	cursor := 0
	res, err := memoryStorage.Scan(&cursor, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Cursor, res.Values, cursor)
}
