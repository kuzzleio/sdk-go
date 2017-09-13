package kuzzle_test

import (
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

func TestSetDefaultIndexNullIndex(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(internal.MockedConnection{}, nil)
	assert.NotNil(t, k.SetDefaultIndex(""))
}

func TestSetDefaultIndex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "myindex", request.Index)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.SetDefaultIndex("myindex")
	k.ListCollections("", nil)
}

func ExampleKuzzle_SetDefaultIndex() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	err := k.SetDefaultIndex("index")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}