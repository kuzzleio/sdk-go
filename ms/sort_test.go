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

func TestSortEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Sort("", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Sort: key required", fmt.Sprint(err))
}

func TestSortError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Sort("foo", qo)

	assert.NotNil(t, err)
}

func TestSort(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sort", parsedQuery.Action)

			r, _ := json.Marshal([]string{"duuude", "iam", "so", "sorted", "right", "now.."})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.Sort("foo", qo)

	assert.Equal(t, []interface{}{"duuude", "iam", "so", "sorted", "right", "now.."}, res)
}

func TestSortWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sort", parsedQuery.Action)

			r, _ := json.Marshal([]string{"duuude", "iam", "so", "sorted", "right", "now.."})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetAlpha(true)
	qo.SetBy("bye")
	qo.SetDirection("DESC")
	qo.SetGet([]string{"jet", "set"})
	qo.SetLimit([]int{0, 42})

	res, _ := memoryStorage.Sort("foo", qo)

	assert.Equal(t, []interface{}{"duuude", "iam", "so", "sorted", "right", "now.."}, res)
}
