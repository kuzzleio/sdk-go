package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateRestrictedUserBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.CreateRestrictedUser("", nil)
	assert.Error(t, err)
}

func TestCreateRestrictedUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.CreateRestrictedUser("body", nil)
	assert.NotNil(t, err)
}

func TestCreateRestrictedUser(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "createRestrictedUser", parsedQuery.Action)
			assert.Equal(t, "myRestrictedUser", parsedQuery.Body)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.CreateRestrictedUser("myRestrictedUser", nil)

	assert.Nil(t, err)
}
