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

func TestSdiffStoreEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SdiffStore("", []string{"bar", "rab"}, "destination", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SdiffStore: key required", fmt.Sprint(err))
}

func TestSdiffStoreEmptySets(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SdiffStore("foo", []string{}, "destination", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SdiffStore: please provide at least one set", fmt.Sprint(err))
}

func TestSdiffStoreEmptyDestination(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SdiffStore("foo", []string{"bar", "rab"}, "", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.SdiffStore: destination required", fmt.Sprint(err))
}

func TestSdiffStoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.SdiffStore("foo", []string{"bar", "rab"}, "destination", qo)

	assert.NotNil(t, err)
}

func TestSdiffStore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sdiffstore", parsedQuery.Action)

			r, _ := json.Marshal(42)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.SdiffStore("foo", []string{"bar", "rab"}, "destination", qo)

	assert.Equal(t, 42, res)
}
