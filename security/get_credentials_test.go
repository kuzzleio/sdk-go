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

func TestGetCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "someId", request.Id)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetCredentials("local", "someId", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetCredentials("", "someId", nil)
	assert.NotNil(t, err)
}

func TestGetCredentialsEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetCredentials("local", "", nil)
	assert.NotNil(t, err)
}

func TestGetCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "someId", request.Id)

			type credentials struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			cred := credentials{"admin", "test"}
			marsh, _ := json.Marshal(cred)

			return types.KuzzleResponse{Result: marsh}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	res, err := s.GetCredentials("local", "someId", nil)
	assert.Nil(t, err)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	cred := credentials{}
	json.Unmarshal(res, &cred)

	assert.Equal(t, "admin", cred.Username)
	assert.Equal(t, "test", cred.Password)
}
