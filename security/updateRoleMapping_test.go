package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRoleMappingBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	err := k.Security.UpdateRoleMapping("", nil)
	assert.Error(t, err)
}

func TestUpdateRoleMappingError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	err := k.Security.UpdateRoleMapping("body", nil)
	assert.Error(t, err)
}

func TestUpdateRoleMapping(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "updateRoleMapping", parsedQuery.Action)
			assert.Equal(t, "body", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte(`{ "acknowledged" : true}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.Security.UpdateRoleMapping("body", nil)
	assert.NoError(t, err)
}
