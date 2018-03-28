package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestReplaceUserIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	res, err := k.Security.ReplaceUser("", []byte(`{"body": "test"}`), nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUserContentNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	res, err := k.Security.ReplaceUser("id", nil, nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ReplaceUser("id", []byte(`{"body": "test"}`), nil)
	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestReplaceUser(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "replaceUser", request.Action)
			assert.Equal(t, "id", request.Id)

			return &types.KuzzleResponse{Result: []byte(`{
          "_id": "id",
          "_source": {
            "profileIds": ["profileId"],
            "name": "John Doe"
          }
        }`),
			}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.ReplaceUser("id", []byte(`{"body": "test"}`), nil)
	assert.NoError(t, err)
	assert.Equal(t, "id", res.Id)
}
