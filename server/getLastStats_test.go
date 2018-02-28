package server_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetLastStatsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getLastStats", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Server.GetLastStats(nil, nil, nil)
	assert.NotNil(t, err)
}

func TestGetLastStats(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getLastStats", request.Action)

			return &types.KuzzleResponse{Result: json.RawMessage(`{"foo": "bar"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Server.GetLastStats(nil, nil, nil)

	assert.Equal(t, json.RawMessage(`{"foo": "bar"}`), res)
}

func ExampleKuzzle_GetLastStats() {
	conn := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	now := time.Now()
	res, err := k.Server.GetLastStats(&now, nil, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
