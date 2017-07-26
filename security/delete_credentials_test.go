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

func TestDeleteCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "deleteCredentials", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.DeleteCredentials("local", "someId", nil)
	assert.NotNil(t, err)
}

func TestDeleteCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.DeleteCredentials("", "someId", nil)
	assert.NotNil(t, err)
}

func TestDeleteCredentialsEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.DeleteCredentials("local", "", nil)
	assert.NotNil(t, err)
}

func TestDeleteCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			type ackResult struct {
				Acknowledged       bool
				ShardsAcknowledged bool
			}

			ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "deleteCredentials", request.Action)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	res, _ := s.DeleteCredentials("local", "someId", nil)

	assert.Equal(t, true, res.Acknowledged)
	assert.Equal(t, true, res.ShardsAcknowledged)
}
