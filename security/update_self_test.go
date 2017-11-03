package security_test

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

func TestUpdateSelfQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.UpdateSelf(nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateSelf(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateSelf", request.Action)

			assert.Equal(t, "foo", request.Body.(map[string]interface{})["username"])

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "login",
				"_source": {
					"username": "foo"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.UpdateSelf(&types.UserData{
		Content: map[string]interface{}{
			"username": "foo",
		},
	}, nil)

	assert.Equal(t, "login", res.Id)
}

func ExampleKuzzle_UpdateSelf() {
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

	res, err := k.UpdateSelf(&types.UserData{
		Content: map[string]interface{}{
			"foo": "bar",
		},
	}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
