package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoleIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateRole("", "body", nil)
	assert.Error(t, err)
}

func TestCreateRoleBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateRole("id", "", nil)
	assert.Error(t, err)
}

func TestCreateRoleError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.CreateRole("id", "body", nil)
	assert.NotNil(t, err)
}

func TestCreateRole(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRole", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "newRole", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.CreateRole("id", "newRole", nil)

	assert.Nil(t, err)
}
