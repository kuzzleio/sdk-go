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

func TestZrevRangeByLexEmptyMin(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.ZrevRangeByLex("foo", "", "(g", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.ZrevRangeByLex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZrevRangeByLexEmptyMax(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.ZrevRangeByLex("foo", "-", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.ZrevRangeByLex: an empty string is not a valid string range item", fmt.Sprint(err))
}

func TestZrevRangeByLexError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.ZrevRangeByLex("foo", "-", "(g", nil)

	assert.NotNil(t, err)
}

func TestZrevRangeByLex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrevrangebylex", parsedQuery.Action)
			assert.Equal(t, []interface{}([]interface{}{"withscores"}), parsedQuery.Options)

			r, _ := json.Marshal([]string{"bar", "rab"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.ZrevRangeByLex("foo", "-", "(g", nil)

	assert.Equal(t, []string{"bar", "rab"}, res)
}

func TestZrevRangeByLexWithLimits(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrevrangebylex", parsedQuery.Action)
			assert.Equal(t, []int{0, 1}, options.GetLimit())
			assert.Equal(t, []interface{}([]interface{}{"withscores"}), parsedQuery.Options)

			r, _ := json.Marshal([]string{"bar", "rab"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetLimit([]int{0, 1})
	res, _ := k.MemoryStorage.ZrevRangeByLex("foo", "-", "(g", qo)

	assert.Equal(t, []string{"bar", "rab"}, res)
}

func ExampleMs_ZrevRangeByLex() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetLimit([]int{0, 1})
	res, err := k.MemoryStorage.ZrevRangeByLex("foo", "-", "(g", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
