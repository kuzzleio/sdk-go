package index_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestListError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	i := index.NewIndex(k)
	_, err := i.List()
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestList(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "index", parsedQuery.Controller)
			assert.Equal(t, "list", parsedQuery.Action)

			res := types.KuzzleResponse{Result: []byte(`
				{
					"total": 2,
					"hits": [
						"index_1",
						"index_2"
					]
  			}`),
			}

			r, _ := json.Marshal(res.Result)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	i := index.NewIndex(k)
	res, err := i.List()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func ExampleIndex_List() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	i := index.NewIndex(k)
	res, err := i.List()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
