package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrReplaceProfileBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateOrReplaceProfile("id", nil, nil)
	assert.Error(t, err)
}

func TestCreateOrReplaceProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.CreateOrReplaceProfile("id", []byte(`{"body": "test"}`), nil)
	assert.NotNil(t, err)
}

func TestCreateOrReplaceProfile(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createOrReplaceProfile", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)

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
	res, err := k.Security.CreateOrReplaceProfile("id", []byte(`{"body": "test"}`), nil)
	assert.Equal(t, "id", res.Id)
	assert.Nil(t, err)
}
