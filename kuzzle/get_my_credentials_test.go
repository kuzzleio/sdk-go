package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMyCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.GetMyCredentials("local", nil)
	assert.NotNil(t, err)
}

func TestGetMyCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.GetMyCredentials("", nil)
	assert.Equal(t, "Kuzzle.GetMyCredentials: strategy is required", fmt.Sprint(err))
}

func TestGetMyCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getMyCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)

			type myCredentials struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			myCred := myCredentials{"admin", "test"}
			marsh, _ := json.Marshal(myCred)

			return types.KuzzleResponse{Result: marsh}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.GetMyCredentials("local", nil)
	assert.Nil(t, err)

	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	cred := myCredentials{}
	json.Unmarshal(res, &cred)

	assert.Equal(t, "admin", cred.Username)
	assert.Equal(t, "test", cred.Password)
}
