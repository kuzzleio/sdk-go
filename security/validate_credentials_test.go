package security_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"fmt"
)

func TestValidateCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "validateCredentials", request.Action)
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ValidateCredentials("local", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestValidateCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ValidateCredentials("", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestValidateCredentialsEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.ValidateCredentials("local", "", nil, nil)
	assert.NotNil(t, err)
}

func TestValidateCredentials(t *testing.T) {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "validateCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "someId", request.Id)

			marsh, _ := json.Marshal(true)

			return &types.KuzzleResponse{Result: marsh}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ValidateCredentials("local", "someId", myCredentials{"foo", "bar"}, nil)
	assert.Nil(t, err)
	assert.Equal(t, true, res)
}

func ExampleSecurity_ValidateCredentials() {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ValidateCredentials("local", "someId", myCredentials{"foo", "bar"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
