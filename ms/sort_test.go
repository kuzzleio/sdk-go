package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	_, err := memoryStorage.Sort("foo", nil)

	assert.NotNil(t, err)
}

func TestSort(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sort", parsedQuery.Action)

			r, _ := json.Marshal([]string{"duuude", "iam", "so", "sorted", "right", "now.."})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)

	res, _ := memoryStorage.Sort("foo", nil)

	assert.Equal(t, []string{"duuude", "iam", "so", "sorted", "right", "now.."}, res)
}

func TestSortWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "sort", parsedQuery.Action)

			r, _ := json.Marshal([]string{"duuude", "iam", "so", "sorted", "right", "now.."})
			return &types.KuzzleResponse{Result: r}
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

	assert.Equal(t, []string{"duuude", "iam", "so", "sorted", "right", "now.."}, res)
}

func ExampleMs_Sort() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetAlpha(true)
	qo.SetBy("bye")
	qo.SetDirection("DESC")
	qo.SetGet([]string{"jet", "set"})
	qo.SetLimit([]int{0, 42})

	res, err := memoryStorage.Sort("foo", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
