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

func TestHincrbyfloatEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Hincrbyfloat("", "bar", 1, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Hincrbyfloat: key required", fmt.Sprint(err))
}

func TestHincrbyfloatEmptyField(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Hincrbyfloat("foo", "", 1, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Hincrbyfloat: field required", fmt.Sprint(err))
}

func TestHincrbyfloatError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Hincrbyfloat("foo", "bar", 1, qo)

	assert.NotNil(t, err)
}

func TestHincrbyfloat(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "hincrbyfloat", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "bar", parsedQuery.Body.(map[string]interface{})["field"].(string))
			assert.Equal(t, float64(42), parsedQuery.Body.(map[string]interface{})["value"].(float64))

			r, _ := json.Marshal("1")
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.Hincrbyfloat("foo", "bar", 42, qo)

	assert.Equal(t, float64(1), res)
}
