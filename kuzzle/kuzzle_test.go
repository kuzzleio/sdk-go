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

func ExampleKuzzle_Connect() {
	opts := types.NewOptions()
	opts.SetConnect(types.Manual)

	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, opts)

	err := k.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(k.State)
}

func ExampleKuzzle_Jwt() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	jwt := k.Jwt()
	fmt.Println(jwt)
}

func ExampleKuzzle_OfflineQueue() {
	//todo
}

func TestUnsetJwt(t *testing.T) {
	var k *kuzzle.Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			assert.Equal(t, 0, request.ExpiresIn)

			type loginResult struct {
				Jwt string `json:"jwt"`
			}

			loginRes := loginResult{"token"}
			marsh, _ := json.Marshal(loginRes)

			return &types.KuzzleResponse{Result: marsh}
		},
		MockGetRooms: func() *types.RoomList {
			return nil
		},
	}

	k, _ = kuzzle.NewKuzzle(c, nil)

	res, _ := k.Login("local", nil, nil)
	assert.Equal(t, "token", res)
	assert.Equal(t, "token", k.Jwt())
	k.UnsetJwt()
	assert.Equal(t, "", k.Jwt())
}

func ExampleKuzzle_UnsetJwt() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}

	_, err := k.Login("local", myCredentials, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	k.UnsetJwt()
	fmt.Println(k.Jwt())
}

func TestSetDefaultIndexNullIndex(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	assert.NotNil(t, k.SetDefaultIndex(""))
}

func TestSetDefaultIndex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "myindex", request.Index)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
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
