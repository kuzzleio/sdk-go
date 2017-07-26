package security_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCredentialFieldsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetCredentialFields("", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialFieldsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getCredentialFields", request.Action)
			assert.Equal(t, "local", request.Strategy)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetCredentialFields("local", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialFields(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getCredentialFields", request.Action)
			assert.Equal(t, "local", request.Strategy)

			marsh, _ := json.Marshal([]string{"username", "password"})

			return types.KuzzleResponse{Result: marsh}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	res, err := s.GetCredentialFields("local", nil)
	assert.Nil(t, err)

	assert.Equal(t, "username", res[0])
	assert.Equal(t, "password", res[1])
}
