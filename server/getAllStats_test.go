package server_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetAllStatsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Server.GetAllStats(nil)
	assert.NotNil(t, err)
}

func TestGetAllStats(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)

			return &types.KuzzleResponse{Result: json.RawMessage(`{"foo": "bar"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Server.GetAllStats(nil)

	assert.Equal(t, json.RawMessage(`{"foo": "bar"}`), res)
}

func ExampleKuzzle_GetAllStats() {
	conn := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.Server.GetAllStats(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
