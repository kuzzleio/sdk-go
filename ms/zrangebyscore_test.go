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

func TestZrangeByScoreEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZrangeByScore("", 1, 6, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.ZrangeByScore: key required", fmt.Sprint(err))
}

func TestZrangeByScoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.ZrangeByScore("foo", 1, 6, qo)

	assert.NotNil(t, err)
}

func TestZrangeByScore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrangebyscore", parsedQuery.Action)

			r, _ := json.Marshal([]string{"bar", "5", "foo", "1.377"})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.ZrangeByScore("foo", 1, 6, qo)

	expectedResult := []types.MSSortedSet{}
	expectedResult = append(expectedResult, types.MSSortedSet{Member: "bar", Score: 5})
	expectedResult = append(expectedResult, types.MSSortedSet{Member: "foo", Score: 1.377})

	assert.Equal(t, expectedResult, res)
}

func TestZrangeByScoreWithLimits(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zrangebyscore", parsedQuery.Action)
			assert.Equal(t, []interface{}([]interface{}{"withscores"}), parsedQuery.Options)
			assert.Equal(t, "0,1", parsedQuery.Limit)

			r, _ := json.Marshal([]string{"bar", "5"})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetLimit([]int{0, 1})
	res, _ := memoryStorage.ZrangeByScore("foo", 1, 6, qo)

	expectedResult := []types.MSSortedSet{}
	expectedResult = append(expectedResult, types.MSSortedSet{Member: "bar", Score: 5})

	assert.Equal(t, expectedResult, res)
}

func ExampleMs_ZrangeByScore() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, err := memoryStorage.ZrangeByScore("foo", 1, 6, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Member, res[0].Score)
}
