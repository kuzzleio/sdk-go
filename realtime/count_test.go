package realtime_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/realtime"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCountError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	realTime := realtime.NewRealtime(k)

	_, err := realTime.Count("index", "collection", "42")
	assert.NotNil(t, err)
}

func TestCount(t *testing.T) {
	type result struct {
		Count int `json:"count"`
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			res := result{Count: 10}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	realTime := realtime.NewRealtime(k)

	res, _ := realTime.Count("index", "collection", "42")
	assert.Equal(t, 10, res)
}

func Example_Count() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	realTime := realtime.NewRealtime(k)

	res, err := realTime.Count("index", "collection", "42")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
