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

func TestGetAllCredentialFieldsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getAllCredentialFields", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.GetAllCredentialFields(nil)
	assert.NotNil(t, err)
}

func TestGetAllCredentialFields(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "getAllCredentialFields", request.Action)

			credFields := types.CredentialFields{}
			credFields["local"] = []string{"username", "password"}
			marsh, _ := json.Marshal(credFields)

			return types.KuzzleResponse{Result: marsh}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	res, err := s.GetAllCredentialFields(nil)
	assert.Nil(t, err)

	assert.Equal(t, "username", res["local"][0])
	assert.Equal(t, "password", res["local"][1])
}
