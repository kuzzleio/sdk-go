package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinsertInvalidPivot(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.MemoryStorage.Linsert("", "position", "pivot", "bar", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Linsert: invalid position argument (must be 'before' or 'after')", fmt.Sprint(err))
}

func TestLinsertError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Linsert("foo", "position", "pivot", "bar", nil)

	assert.NotNil(t, err)
}

func TestLinsert(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "linsert", parsedQuery.Action)
			assert.Equal(t, "bar", parsedQuery.Body.(map[string]interface{})["value"].(string))
			assert.Equal(t, "before", parsedQuery.Body.(map[string]interface{})["position"].(string))
			assert.Equal(t, "pivot", parsedQuery.Body.(map[string]interface{})["pivot"].(string))

			r, _ := json.Marshal(1)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Linsert("foo", "before", "pivot", "bar", nil)

	assert.Equal(t, 1, res)
}

func ExampleMs_Linsert() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := k.MemoryStorage.Linsert("foo", "position", "pivot", "bar", nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
