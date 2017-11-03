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

func TestZinterStoreEmptyKeys(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.ZinterStore("foo", []string{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.ZinterStore: please provide at least one key", fmt.Sprint(err))
}

func TestZinterStoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.ZinterStore("foo", []string{"bar", "rab"}, nil)

	assert.NotNil(t, err)
}

func TestZinterStore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zinterstore", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.ZinterStore("foo", []string{"bar", "rab"}, nil)

	assert.Equal(t, 2, res)
}

func TestZinterStoreWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zinterstore", parsedQuery.Action)
			assert.Equal(t, "sum", options.GetAggregate())
			assert.Equal(t, []int{1, 2}, options.GetWeights())

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetAggregate("sum")
	qo.SetWeights([]int{1, 2})
	res, _ := k.MemoryStorage.ZinterStore("foo", []string{"bar", "rab"}, qo)

	assert.Equal(t, 2, res)
}

func ExampleMs_ZinterStore() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetAggregate("sum")
	qo.SetWeights([]int{1, 2})
	res, err := k.MemoryStorage.ZinterStore("foo", []string{"bar", "rab"}, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
