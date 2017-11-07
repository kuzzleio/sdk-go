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

func TestZlexcountEmptyMin(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zlexcount("foo", "", "[f", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zlexcount: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZlexcountEmptyMax(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zlexcount("foo", "[b", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zlexcount: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZlexcountError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Zlexcount("foo", "[b", "[f", nil)

	assert.NotNil(t, err)
}

func TestZlexcount(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zlexcount", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Zlexcount("foo", "[b", "[f", nil)

	assert.Equal(t, 2, res)
}

func ExampleMs_Zlexcount() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Zlexcount("foo", "[b", "[f", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
