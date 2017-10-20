package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestWhoAmIQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getCurrentUser", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.WhoAmI()
	assert.NotNil(t, err)
}

func TestWhoAmI(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getCurrentUser", request.Action)

			toMarshal := types.User{Id: "id"}

			h, err := json.Marshal(toMarshal)
			if err != nil {
				log.Fatal(err)
			}

			return &types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.WhoAmI()

	assert.Equal(t, "id", res.Id)
}

func ExampleKuzzle_WhoAmI() {
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

	res, err := k.WhoAmI()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Strategies, res.Meta, res.Source)
}
