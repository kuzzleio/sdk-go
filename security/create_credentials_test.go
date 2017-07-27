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

func TestCreateCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "createCredentials", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.CreateCredentials("local", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestCreateCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.CreateCredentials("", "someId", nil, nil)
	assert.NotNil(t, err)
}

func TestCreateCredentialsEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	s := security.NewSecurity(k)
	_, err := s.CreateCredentials("local", "", nil, nil)
	assert.NotNil(t, err)
}

func TestCreateCredentials(t *testing.T) {
	type myCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			ack := myCredentials{Username: "foo", Password: "bar"}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "security", request.Controller)
			assert.Equal(t, "createCredentials", request.Action)
			assert.Equal(t, "local", request.Strategy)
			assert.Equal(t, "someId", request.Id)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	s := security.NewSecurity(k)
	res, err := s.CreateCredentials("local", "someId", myCredentials{"foo", "bar"}, nil)
	assert.Nil(t, err)

	assert.Equal(t, "foo", res["username"])
	assert.Equal(t, "bar", res["password"])
}