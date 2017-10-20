package kuzzle_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListIndexesQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "list", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.ListIndexes(nil)
	assert.NotNil(t, err)
}

func TestListIndexes(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "index", request.Controller)
			assert.Equal(t, "list", request.Action)

			type indexes struct {
				Indexes []string `json:"indexes"`
			}

			list := []string{
				"index1",
				"index2",
			}

			c := indexes{
				Indexes: list,
			}

			h, _ := json.Marshal(c)
			return &types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.ListIndexes(nil)

	assert.Equal(t, "index1", res[0])
	assert.Equal(t, "index2", res[1])
}

func ExampleKuzzle_ListIndexes() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.ListIndexes(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, index := range res {
		fmt.Println(index)
	}
}
