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

func TestBitposError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Bitpos("foo", 0, nil)

	assert.NotNil(t, err)
}

func TestBitpos(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "bitpos", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, 2, parsedQuery.Bit)

			r, _ := json.Marshal(1)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Bitpos("foo", 2, nil)

	assert.Equal(t, 1, res)
}

func TestBitposOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "bitpos", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, 2, parsedQuery.Bit)
			assert.Equal(t, 1, parsedQuery.Start)
			assert.Equal(t, 2, parsedQuery.End)

			r, _ := json.Marshal(1)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetStart(1).SetEnd(2)

	res, _ := k.MemoryStorage.Bitpos("foo", 2, qo)

	assert.Equal(t, 1, res)
}

func ExampleMs_Bitpos() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Bitpos("foo", 2, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
