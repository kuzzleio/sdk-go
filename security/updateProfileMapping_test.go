package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProfileMappingBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	err := k.Security.UpdateProfileMapping(nil, nil)
	assert.Error(t, err)
}

func TestUpdateProfileMappingError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	err := k.Security.UpdateProfileMapping([]byte(`{"body": "test"}`), nil)
	assert.Error(t, err)
}

func TestUpdateProfileMapping(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateProfileMapping", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`{ "acknowledged" : true}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.Security.UpdateProfileMapping([]byte(`{"body": "test"}`), nil)
	assert.NoError(t, err)
}
