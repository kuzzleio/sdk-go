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
			ack := ackResult{Acknowledged: true, ShardsAcknowledged: true}
			r, _ := json.Marshal(ack)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.CreateIndex("index", nil)
	assert.Equal(t, true, res.Acknowledged)
	assert.Equal(t, true, res.ShardsAcknowledged)
}

func ExampleKuzzle_CreateIndex() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.CreateIndex("index", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Acknowledged, res.ShardsAcknowledged)
}
