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

func TestZscanEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 42
	_, err := memoryStorage.Zscan("", &cursor, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Zscan: key required", fmt.Sprint(err))
}

func TestZscanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 42
	_, err := memoryStorage.Zscan("foo", &cursor, qo)

	assert.NotNil(t, err)
}

func TestZscan(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 42,
		Values: []string{"bar", "5", "foo", "10"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zscan", parsedQuery.Action)

			r, _ := json.Marshal([]interface{}{"42", []string{"bar", "5", "foo", "10"}})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 42
	res, _ := memoryStorage.Zscan("foo", &cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func TestZscanWithOptions(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 42,
		Values: []string{"bar", "5", "foo", "10"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zscan", parsedQuery.Action)
			assert.Equal(t, 10, options.GetCount())
			assert.Equal(t, "*", options.GetMatch())

			r, _ := json.Marshal([]interface{}{"42", []string{"bar", "5", "foo", "10"}})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 42
	qo.SetCount(10)
	qo.SetMatch("*")
	res, _ := memoryStorage.Zscan("foo", &cursor, qo)

	assert.Equal(t, scanResponse, res)
}

func ExampleMs_Zscan() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()


	cursor := 42
	qo.SetCount(10)
	qo.SetMatch("*")
	res, err := memoryStorage.Zscan("foo", &cursor, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Cursor, res.Values, cursor)
}
