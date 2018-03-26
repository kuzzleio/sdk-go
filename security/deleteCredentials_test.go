package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCredentialsStrategyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	err := k.Security.DeleteCredentials("", "id", nil)
	assert.Error(t, err)
}

func TestDeleteCredentialsIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	err := k.Security.DeleteCredentials("strategy", "", nil)
	assert.Error(t, err)
}

func TestDeleteCredentialsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	err := k.Security.DeleteCredentials("strategy", "id", nil)
	assert.NotNil(t, err)
}

func TestDeleteCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "deleteCredentials", parsedQuery.Action)
			assert.Equal(t, "strategy", parsedQuery.Strategy)
			assert.Equal(t, "id", parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte{}}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	err := k.Security.DeleteCredentials("strategy", "id", nil)

	assert.Nil(t, err)
}
