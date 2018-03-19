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

func TestExistsNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	i := index.NewIndex(k)
	_, err := i.Exists("")
	assert.NotNil(t, err)
}

func TestExistsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	_, err := i.Exists("index")
	assert.NotNil(t, err)
}

func TestExists(t *testing.T) {
	type ackResult struct {
		Acknowledged       bool
		ShardsAcknowledged bool
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			q := &types.KuzzleRequest{}
			json.Unmarshal(query, q)

			assert.Equal(t, "index", q.Controller)
			assert.Equal(t, "exists", q.Action)
			assert.Equal(t, "index", q.Index)

			return &types.KuzzleResponse{Result: []byte(`true`)}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, nil)
	i := index.NewIndex(k)
	res, err := i.Exists("index")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res)
}

func ExampleIndex_Exists() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)
	i := index.NewIndex(k)
	i.Create("index")
	_, err := i.Exists("index")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}