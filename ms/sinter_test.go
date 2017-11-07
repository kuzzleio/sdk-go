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

func TestSinterEmptyKeys(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Sinter([]string{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Sinter: please provide at least one key to intersect", fmt.Sprint(err))
}

func TestSinterError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Sinter([]string{"foo", "bar"}, nil)

	assert.NotNil(t, err)
}

func TestSinter(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sinter", parsedQuery.Action)

			r, _ := json.Marshal([]string{"diff1", "diff2"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Sinter([]string{"foo", "bar"}, nil)

	assert.Equal(t, []string{"diff1", "diff2"}, res)
}

func ExampleMs_Sinter() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Sinter([]string{"foo", "bar"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
