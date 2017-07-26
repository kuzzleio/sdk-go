package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefreshIndexQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "refresh", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.RefreshIndex("index", nil)
	assert.NotNil(t, err)
}

func TestRefreshIndex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			type shards struct {
				Shards types.Shards `json:"_shards"`
			}

			ack := shards{types.Shards{10, 9, 8}}
			r, _ := json.Marshal(ack)

			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "refresh", request.Action)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.RefreshIndex("index", nil)

	assert.Equal(t, 10, res.Total)
	assert.Equal(t, 9, res.Successful)
	assert.Equal(t, 8, res.Failed)
}
