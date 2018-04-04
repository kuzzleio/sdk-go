package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCredentialsStrategyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentials("", "id", nil)
	assert.Error(t, err)
}

func TestGetCredentialsIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentials("strategy", "", nil)
	assert.Error(t, err)
}

func TestGetCredentialsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.GetCredentials("strategy", "id", nil)
	assert.NotNil(t, err)
}

func TestGetCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getCredentials", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "strategy", parsedQuery.Strategy)

			return &types.KuzzleResponse{Result: []byte(`{"username": "user", "kuid": "id"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.GetCredentials("strategy", "id", nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
