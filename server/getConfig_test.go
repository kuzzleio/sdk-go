package server_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getConfig", request.Action)

			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Server.GetConfig(nil)
	assert.NotNil(t, err)
}

func TestGetConfig(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getConfig", request.Action)

			return &types.KuzzleResponse{Result: json.RawMessage(`{"foo": "bar"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Server.GetConfig(nil)
	assert.Equal(t, json.RawMessage(`{"foo": "bar"}`), res)
}

func ExampleGetConfig() {
	c := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Server.GetConfig(nil)
	println(res)
}