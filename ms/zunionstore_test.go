package ms_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"testing"
	"fmt"
)

func TestZunionStoreEmptyDestination(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZunionStore("", []string{"bar", "rab"}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.ZunionStore: destination required", fmt.Sprint(err))
}

func TestZunionStoreEmptyKeys(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZunionStore("foo", []string{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.ZunionStore: please provide at least one key", fmt.Sprint(err))
}

func TestZunionStoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZunionStore("foo", []string{"bar", "rab"}, qo)

	assert.NotNil(t, err)
}

func TestZunionStore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zunionstore", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.ZunionStore("foo", []string{"bar", "rab"}, qo)

	assert.Equal(t, 2, res)
}

func TestZunionStoreWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zunionstore", parsedQuery.Action)
			assert.Equal(t, "sum", options.GetAggregate())
			assert.Equal(t, []int{1, 2}, options.GetWeights())

			r, _ := json.Marshal(2)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetAggregate("sum")
	qo.SetWeights([]int{1, 2})
	res, _ := memoryStorage.ZunionStore("foo", []string{"bar", "rab"}, qo)

	assert.Equal(t, 2, res)
}
