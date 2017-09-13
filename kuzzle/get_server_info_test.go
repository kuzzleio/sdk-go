package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestGetServerInfoQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "info", request.Action)

			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.GetServerInfo(nil)
	assert.NotNil(t, err)
}

func TestGetServerInfo(t *testing.T) {
	type myServerInfo struct {
		Kuzzle struct {
			Version string `json:"version"`
		} `json:"kuzzle"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "info", request.Action)

			msi := myServerInfo{struct {
				Version string `json:"version"`
			}{"1.0.1"}}
			msiMarsh, _ := json.Marshal(msi)

			type serverInfo struct {
				ServerInfo json.RawMessage `json:"serverInfo"`
			}
			si := serverInfo{msiMarsh}

			info, _ := json.Marshal(si)
			return types.KuzzleResponse{Result: info}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.GetServerInfo(nil)
	fmt.Printf("%s\n", res)

	var info myServerInfo
	json.Unmarshal(res, &info)
	assert.Equal(t, "1.0.1", info.Kuzzle.Version)
}

func ExampleKuzzle_GetServerInfo() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.GetServerInfo(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}