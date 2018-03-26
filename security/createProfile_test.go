package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfileIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateProfile("", "body", nil)
	assert.Error(t, err)
}

func TestCreateProfileBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateProfile("id", "", nil)
	assert.Error(t, err)
}

func TestCreateProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.CreateProfile("id", "body", nil)
	assert.NotNil(t, err)
}

func TestCreateProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createProfile", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "newProfile", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte(`{
			    "_id": "id",
			    "_index": "%kuzzle",
			    "_type": "profiles",
			    "_version": 1
		    }`),
			}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.CreateProfile("id", "newProfile", nil)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, "id", res.Id)
}
