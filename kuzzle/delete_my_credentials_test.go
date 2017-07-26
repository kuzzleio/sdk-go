package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteMyCredentialsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "deleteMyCredentials", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.DeleteMyCredentials("local", nil)
	assert.NotNil(t, err)
}

func TestDeleteMyCredentialsEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.DeleteMyCredentials("", nil)
	assert.NotNil(t, err)
}

func TestDeleteMyCredentials(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			type ackResult struct {
				Acknowledged       bool
				ShardsAcknowledged bool
			}

			ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "deleteMyCredentials", request.Action)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.DeleteMyCredentials("local", nil)

	assert.Equal(t, true, res.Acknowledged)
	assert.Equal(t, true, res.ShardsAcknowledged)
}
