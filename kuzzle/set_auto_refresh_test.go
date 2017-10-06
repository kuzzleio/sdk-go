package kuzzle_test

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

func TestSetAutoRefreshDefaultIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.SetAutoRefresh("", true, nil)
	assert.NotNil(t, err)
}

func TestSetAutoRefreshIndexNull(t *testing.T) {
	opts := types.NewOptions()
	opts.SetDefaultIndex("myIndex")

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: json.RawMessage("true")}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, opts)
	_, err := k.SetAutoRefresh("", true, nil)
	assert.Nil(t, err)
}

func TestSetAutoRefreshQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "setAutoRefresh", request.Action)

			type autoRefresh struct {
				AutoRefresh bool `json:"autoRefresh"`
			}
			assert.Equal(t, true, request.Body.(map[string]interface{})["autoRefresh"])

			return &types.KuzzleResponse{Error: &types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.SetAutoRefresh("index", true, nil)
	assert.NotNil(t, err)
}

func TestSetAutoRefresh(t *testing.T) {
	type autorefresh struct {
		Response bool `json:"response"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "setAutoRefresh", request.Action)

			ar := autorefresh{true}
			marshaled, _ := json.Marshal(ar)

			return &types.KuzzleResponse{Result: marshaled}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.SetAutoRefresh("index", true, nil)

	assert.Equal(t, true, res)
}

func ExampleKuzzle_SetAutoRefresh() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.SetAutoRefresh("index", true, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
