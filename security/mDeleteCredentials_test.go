package security_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestMDeleteCredentialsEmptyIDs(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.Security.MDeleteCredentials([]string(nil), nil)
	assert.Error(t, err)
}

func TestMDeleteCredentialsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.Security.MDeleteCredentials([]string{"id"}, nil)
	assert.NotNil(t, err)
}

func TestMDeleteCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "mDeleteCredentials", parsedQuery.Action)

			return &types.KuzzleResponse{Result: []byte(`["id"]`)}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.MDeleteCredentials([]string{"id"}, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "id", res[0])
}
