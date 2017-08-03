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

func TestHscanEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cur := 0
	_, err := memoryStorage.Hscan("", &cur, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Hscan: key required", fmt.Sprint(err))
}

func TestHscanError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cur := 0
	_, err := memoryStorage.Hscan("foo", &cur, qo)

	assert.NotNil(t, err)
}

func TestHscanCursorConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "hscan", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, 1, *parsedQuery.Cursor)

			var result []interface{}
			values := []string{"some", "results"}

			result = append(result, "12abc")
			result = append(result, values)

			r, _ := json.Marshal(result)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 1
	_, err := memoryStorage.Hscan("foo", &cursor, qo)

	assert.NotNil(t, err)
}

func TestHscan(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "hscan", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, 1, *parsedQuery.Cursor)

			var result []interface{}
			values := []string{"some", "results"}

			result = append(result, "12")
			result = append(result, values)

			r, _ := json.Marshal(result)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	cursor := 1
	res, _ := memoryStorage.Hscan("foo", &cursor, qo)

	assert.Equal(t, MemoryStorage.HscanResponse{Cursor: 12, Values: []string{"some", "results"}}, res)
}
