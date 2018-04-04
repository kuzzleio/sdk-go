package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCredentialsByIDStrategyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentialsByID("", "id", nil)
	assert.Error(t, err)
}

func TestGetCredentialsByIDIDNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.GetCredentialsByID("strategy", "", nil)
	assert.Error(t, err)
}

func TestGetCredentialsByIDError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.GetCredentialsByID("strategy", "id", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialsByID(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getCredentialsById", parsedQuery.Action)
			assert.Equal(t, "id", parsedQuery.Id)
			assert.Equal(t, "strategy", parsedQuery.Strategy)

			return &types.KuzzleResponse{Result: []byte(`{"username": "user", "kuid": "id"}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.Security.GetCredentialsByID("strategy", "id", nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
