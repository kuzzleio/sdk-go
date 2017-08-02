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

func TestLinsertEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Linsert("", "position", "pivot", "bar", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Linsert: key required", fmt.Sprint(err))
}

func TestLinsertError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Linsert("foo", "position", "pivot", "bar", qo)

	assert.NotNil(t, err)
}

func TestLinsert(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "linsert", parsedQuery.Action)
			assert.Equal(t, "bar", parsedQuery.Body.(map[string]interface{})["value"].(string))
			assert.Equal(t, "position", parsedQuery.Body.(map[string]interface{})["position"].(string))
			assert.Equal(t, "pivot", parsedQuery.Body.(map[string]interface{})["pivot"].(string))

			r, _ := json.Marshal(1)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.Linsert("foo", "position", "pivot", "bar", qo)

	assert.Equal(t, 1, res)
}
