package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProfileIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.UpdateProfile("", "body", nil)
	assert.Error(t, err)
}

func TestUpdateProfileBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.UpdateProfile("id", "", nil)
	assert.Error(t, err)
}

func TestUpdateProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.UpdateProfile("id", "body", nil)
	assert.Error(t, err)
}

func TestUpdateProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateProfile", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "body", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte(`{
          "_id": "id",
          "_index": "%kuzzle",
          "_type": "profile",
          "_version": 2
        }`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.UpdateProfile("id", "body", nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "id", res.Id)
}
