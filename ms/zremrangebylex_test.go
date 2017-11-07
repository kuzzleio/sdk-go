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

func TestZremrangebylexEmptyMin(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zremrangebylex("foo", "", "(f", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zremrangebylex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZremrangebylexEmptyMax(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zremrangebylex("foo", "[b", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zremrangebylex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZremrangebylexError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Zremrangebylex("foo", "[b", "(f", nil)

	assert.NotNil(t, err)
}

func TestZremrangebylex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zremrangebylex", parsedQuery.Action)

			r, _ := json.Marshal(42)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Zremrangebylex("foo", "[b", "(f", nil)

	assert.Equal(t, 42, res)
}

func ExampleMs_Zremrangebylex() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Zremrangebylex("foo", "[b", "(f", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
