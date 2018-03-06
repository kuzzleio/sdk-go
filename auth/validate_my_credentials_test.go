package auth_test

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

func TestValidateMyCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "validateMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Auth.ValidateMyCredentials("local", nil, nil)
	assert.NotNil(t, err)
}

func TestValidateMyCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "validateMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "foo", request.Body.(map[string]interface{})["username"])
			assert.Equal(t, "bar", request.Body.(map[string]interface{})["password"])

			ret, _ := json.Marshal(true)
			return &types.KuzzleResponse{Result: ret}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Auth.ValidateMyCredentials("local", struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{"foo", "bar"}, nil)
	assert.Nil(t, err)

	assert.Equal(t, true, res)
}

func ExampleKuzzle_ValidateMyCredentials() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
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

	res, err := k.Auth.ValidateMyCredentials("local", myCredentials, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}