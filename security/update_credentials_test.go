package security_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

func TestUpdateCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "updateCredentials", request.Action)
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.UpdateCredentials("local", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.UpdateCredentials("", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateCredentialsEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.UpdateCredentials("local", "", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateCredentials(t *testing.T) {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			ack := myCredentials{Username: "foo", Password: "bar"}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "updateCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "someId", request.Id)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.UpdateCredentials("local", "someId", myCredentials{"foo", "bar"}, nil)
	assert.Nil(t, err)

	assert.Equal(t, "foo", res["username"])
	assert.Equal(t, "bar", res["password"])
}

func ExampleSecurity_UpdateCredentials() {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.UpdateCredentials("local", "someId", myCredentials{"foo", "bar"}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
