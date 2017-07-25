package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListIndexesQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "list", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.ListIndexes(nil)
	assert.NotNil(t, err)
}

func TestListIndexes(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "list", request.Action)

			type indexes struct {
				Indexes []string `json:"indexes"`
			}

			list := make([]string, 0)
			list = append(list, "index1")
			list = append(list, "index2")

			c := indexes{
				Indexes: list,
			}

			h, _ := json.Marshal(c)
			return types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.ListIndexes(nil)

	assert.Equal(t, "index1", res[0])
	assert.Equal(t, "index2", res[1])
}
