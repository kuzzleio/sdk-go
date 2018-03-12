package index_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRefreshInternalQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	err := i.RefreshInternal()
	assert.NotNil(t, err)
}

func TestRefreshInternal(t *testing.T) {
	type ackResult struct {
		Acknowledged       bool
		ShardsAcknowledged bool
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "index", q.Controller)
			assert.Equal(t, "refreshInternal", q.Action)

			return &types.KuzzleResponse{Result: []byte(`"result":true`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	err := i.RefreshInternal()

	assert.Nil(t, err)
}

func ExampleIndex_RefreshInternal() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)
	i := index.NewIndex(k)
	err := i.RefreshInternal()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
