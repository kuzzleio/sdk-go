package kuzzle_test

import (
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlushQueue(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	*k.OfflineQueue() = append(*k.OfflineQueue(), &types.QueryObject{RequestId: "test"})
	assert.NotEmpty(t, *k.OfflineQueue())

	k.FlushQueue()
	assert.Empty(t, *k.OfflineQueue())
}

func ExampleKuzzle_FlushQueue() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	k.FlushQueue()
}
