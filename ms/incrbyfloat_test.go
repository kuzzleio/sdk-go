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

func TestIncrbyfloatEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Incrbyfloat("", 42, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Incrbyfloat: key required", fmt.Sprint(err))
}

func TestIncrbyfloatError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Incrbyfloat("foo", 42, qo)

	assert.NotNil(t, err)
}

func TestIncrbyfloat(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "incrbyfloat", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, float64(2), parsedQuery.Body.(map[string]interface{})["value"].(float64))

			r, _ := json.Marshal("3.14")
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.Incrbyfloat("foo", 42, qo)

	assert.Equal(t, 3.14, res)
}

func ExampleMs_Incrbyfloat() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, err := memoryStorage.Incrbyfloat("foo", 42, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
