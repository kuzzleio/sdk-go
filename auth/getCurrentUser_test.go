package auth_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/auth"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentUserQueryError(t *testing.T) {
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
	auth := auth.NewAuth(k)
	_, err := auth.GetCurrentUser()
	assert.NotNil(t, err)
}

func TestGetCurrentUser(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getCurrentUser", request.Action)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "id"
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	auth := auth.NewAuth(k)
	res, _ := auth.GetCurrentUser()

	assert.Equal(t, "id", res.Id)
}

func ExampleKuzzle_GetCurrentUser() {
	conn := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}

	_, err := k.Auth.Login("local", myCredentials, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	auth := auth.NewAuth(k)
	res, _ := auth.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v\n", res)
}