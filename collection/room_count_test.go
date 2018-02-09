package collection

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestRoomCountError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := NewRoom(NewCollection(k, "collection", "index"), nil, nil).Count()
	assert.NotNil(t, err)
}

func TestRoomCount(t *testing.T) {
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

	r := NewRoom(NewCollection(k, "collection", "index"), nil, nil)
	r.internalState = active
	res, _ := r.Count()
	assert.Equal(t, 10, res)
}

func ExampleRoom_Count() {
	type result struct {
		Count int `json:"count"`
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := NewRoom(NewCollection(k, "collection", "index"), nil, nil).Count()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
