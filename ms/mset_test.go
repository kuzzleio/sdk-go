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

func TestMsetEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Mset([]types.MSKeyValue{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Mset: please provide at least one key/value entry", fmt.Sprint(err))
}

func TestMsetError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []types.MSKeyValue{}
	entries = append(entries, types.MSKeyValue{Key: "foo", Value: "bar"})
	entries = append(entries, types.MSKeyValue{Key: "bar", Value: "foo"})

	_, err := memoryStorage.Mset(entries, qo)

	assert.NotNil(t, err)
}

func TestMset(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "mset", parsedQuery.Action)

			r, _ := json.Marshal("OK")
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []types.MSKeyValue{}
	entries = append(entries, types.MSKeyValue{Key: "foo", Value: "bar"})
	entries = append(entries, types.MSKeyValue{Key: "bar", Value: "foo"})

	res, _ := memoryStorage.Mset(entries, qo)

	assert.Equal(t, "OK", res)
}
