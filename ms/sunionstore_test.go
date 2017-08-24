package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSunionStoreEmptyDestination(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SunionStore("", []string{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SunionStore: destination required", fmt.Sprint(err))
}

func TestSunionStoreEmptySet(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SunionStore("destination", []string{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SunionStore: please provide at least 2 sets", fmt.Sprint(err))
}

func TestSunionStoreSingleSet(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SunionStore("destination", []string{"foo"}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SunionStore: please provide at least 2 sets", fmt.Sprint(err))
}

func TestSunionStoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SunionStore("destination", []string{"foo", "bar"}, qo)

	assert.NotNil(t, err)
}

func TestSunionStore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sunionstore", parsedQuery.Action)

			r, _ := json.Marshal(4)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.SunionStore("destination", []string{"foo", "bar"}, qo)

	assert.Equal(t, 4, res)
}
