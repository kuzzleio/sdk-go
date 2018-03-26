package security_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetProfileEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetProfile("", nil)

	assert.NotNil(t, err)
}

func TestGetProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.GetProfile("profileId", nil)
	assert.NotNil(t, err)
}

func TestGetProfile(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`{
			    "_id": "profileId",
			    "index": "%kuzzle",
			    "collection": "profiles",
			    "action": "getProfile",
			    "controller": "security",
			    "requestId": "<unique request identifier>"
			  }`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.GetProfile(id, nil)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, id, res.Id)

}

func ExampleGetProfile() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.GetProfile(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Policies)
}
