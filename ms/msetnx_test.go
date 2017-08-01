package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"testing"
)

func TestMsetNxEmptyEntries(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.MsetNx([]types.MSKeyValue{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.MsetNx: please provide at least one key/value entry", fmt.Sprint(err))
}

func TestMsetNxError(t *testing.T) {
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

	_, err := memoryStorage.MsetNx(entries, qo)

	assert.NotNil(t, err)
}

func TestMsetNx(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "msetnx", parsedQuery.Action)

			r, _ := json.Marshal(1)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	entries := []types.MSKeyValue{}
	entries = append(entries, types.MSKeyValue{Key: "foo", Value: "bar"})
	entries = append(entries, types.MSKeyValue{Key: "bar", Value: "foo"})

	res, _ := memoryStorage.MsetNx(entries, qo)

	assert.Equal(t, 1, res)
}
