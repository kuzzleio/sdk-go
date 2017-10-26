package security_test

import (
	"fmt"
	"testing"

	"encoding/json"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestFetchEmptyKuid(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchUser("", nil)

	assert.NotNil(t, err)
}

func TestFetchUserError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchUser("userId", nil)
	assert.NotNil(t, err)
}

func TestFetchUser(t *testing.T) {
	id := "userId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getUser", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
				"_id": "userId",
				"_source": {
					"profileIds": ["admin", "other"],
					"name": "Luke",
					"function": "Jedi"
				}
			}`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.FetchUser(id, nil)

	assert.Equal(t, id, res.Id)

	assert.Equal(t, []string{"admin", "other"}, res.ProfileIds)

	assert.Equal(t, "Luke", res.Content["name"])
	assert.Equal(t, "Jedi", res.Content["function"])

	contentAsMap := make(map[string]interface{})
	contentAsMap["name"] = "Luke"
	contentAsMap["function"] = "Jedi"

	assert.Equal(t, contentAsMap, res.Content)
}

func ExampleUser_Fetch() {
	id := "userId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.FetchUser(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}
