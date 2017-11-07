package ms_test

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

func TestGeodistError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Geodist("foo", "bar", "barbar", nil)

	assert.NotNil(t, err)
}

func TestGeodist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geodist", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "bar", parsedQuery.Member1)
			assert.Equal(t, "barbar", parsedQuery.Member2)

			r, _ := json.Marshal(float64(42))
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Geodist("foo", "bar", "barbar", nil)

	assert.Equal(t, float64(42), res)
}

func TestGeodistWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geodist", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "bar", parsedQuery.Member1)
			assert.Equal(t, "barbar", parsedQuery.Member2)

			r, _ := json.Marshal(float64(42))
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetUnit("km")

	res, _ := k.MemoryStorage.Geodist("foo", "bar", "barbar", qo)

	assert.Equal(t, float64(42), res)
}

func ExampleMs_Geodist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetUnit("km")

	res, err := k.MemoryStorage.Geodist("foo", "bar", "barbar", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
