package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogoutError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "logout", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	error := k.Logout()
	assert.NotNil(t, error)
}

func TestLogout(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "logout", request.Action)

			return types.KuzzleResponse{}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	error := k.Logout()
	assert.Nil(t, error)
}
