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

func TestPfmergeEmptySources(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	err := k.MemoryStorage.Pfmerge("foo", []string{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Pfmerge: please provide at least one source to merge", fmt.Sprint(err))
}

func TestPfmergeError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Pfmerge("foo", []string{"bar", "rab"}, nil)

	assert.NotNil(t, err)
}

func TestPfmerge(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "pfmerge", parsedQuery.Action)

			r, _ := json.Marshal("OK")
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Pfmerge("foo", []string{"bar", "rab"}, nil)

	assert.Nil(t, err)
}

func ExampleMs_Pfmerge() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.MemoryStorage.Pfmerge("foo", []string{"bar", "rab"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success")
}
