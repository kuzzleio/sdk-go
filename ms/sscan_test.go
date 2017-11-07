package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSscanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.MemoryStorage.Sscan("foo", 0, nil)

	assert.NotNil(t, err)
}

func TestSscan(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sscan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.MemoryStorage.Sscan("foo", 0, nil)

	assert.Equal(t, &scanResponse, res)
}

func TestSscanWithOptions(t *testing.T) {
	scanResponse := types.MSScanResponse{
		Cursor: 10,
		Values: []string{"foo", "bar"},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sscan", parsedQuery.Action)

			r, _ := json.Marshal(scanResponse)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetCount(42)
	qo.SetMatch("*")

	res, _ := k.MemoryStorage.Sscan("foo", 10, qo)

	assert.Equal(t, &scanResponse, res)
}

func ExampleMs_Sscan() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetCount(42)
	qo.SetMatch("*")

	res, err := k.MemoryStorage.Sscan("foo", 0, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Cursor, res.Values)
}
