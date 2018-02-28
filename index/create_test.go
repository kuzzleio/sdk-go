package index_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.CreateIndex("", nil)
	assert.NotNil(t, err)
}

func TestCreateIndexQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.CreateIndex("index", nil)
	assert.NotNil(t, err)
}

func TestCreateIndex(t *testing.T) {
	type ackResult struct {
		Acknowledged       bool
		ShardsAcknowledged bool
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "index", q.Controller)
			assert.Equal(t, "create", q.Action)
			assert.Equal(t, "index", q.Index)

			return &types.KuzzleResponse{Result: []byte(`{"acknowledged":true}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.CreateIndex("index", nil)

	assert.Nil(t, err)
}

func ExampleKuzzle_CreateIndex() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.CreateIndex("index", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%#v\n", res)
}
