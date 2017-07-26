package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateMyCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateMyCredentials", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.UpdateMyCredentials("local", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateMyCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.UpdateMyCredentials("", nil, nil)
	assert.NotNil(t, err)
}

func TestUpdateMyCredentials(t *testing.T) {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			ack := myCredentials{Username: "foo", Password: "bar"}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "updateMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.UpdateMyCredentials("local", myCredentials{"foo", "bar"}, nil)

	assert.Equal(t, "foo", res["username"])
	assert.Equal(t, "bar", res["password"])
}
