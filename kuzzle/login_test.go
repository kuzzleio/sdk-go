package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginNoStrategy(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Login("", struct{}{}, nil)
	assert.NotNil(t, err)
}

func TestLoginError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "error"}}
		},
		MockEmitEvent: func(e int, arg interface{}) {
			assert.Equal(t, event.LoginAttempt, e)
			assert.Equal(t, "error", arg.(*types.LoginAttempt).Error.Error())
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Login("local", struct{}{}, nil)
}

func TestLogin(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			assert.Equal(t, 42, request.ExpiresIn)

			type loginResult struct {
				Jwt string `json:"jwt"`
			}

			loginRes := loginResult{"token"}
			marsh, _ := json.Marshal(loginRes)

			return &types.KuzzleResponse{Result: marsh}
		},
		MockEmitEvent: func(e int, arg interface{}) {
			assert.Equal(t, event.LoginAttempt, e)
			assert.Equal(t, true, arg.(*types.LoginAttempt).Success)
			assert.Nil(t, arg.(*types.LoginAttempt).Error)
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	expiresIn := 42
	token, _ := k.Login("local", struct{}{}, &expiresIn)
	assert.Equal(t, "token", token)
}

func ExampleKuzzle_Login() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}

	jwt, err := k.Login("local", myCredentials, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(jwt)
}
