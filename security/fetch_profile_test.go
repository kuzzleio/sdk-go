package security_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestFetchProfileEmptyId(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchProfile("", nil)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*types.KuzzleError).Status)
}

func TestFetchProfileError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{
				Error: types.NewError("Test error"),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Security.FetchProfile("profileId", nil)
	assert.NotNil(t, err)
}

func TestFetchProfile(t *testing.T) {
	id := "profileId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "security", parsedQuery.Controller)
			assert.Equal(t, "getProfile", parsedQuery.Action)
			assert.Equal(t, id, parsedQuery.Id)

			p := &security.Profile{
				Id: id,
				Policies: []*types.Policy{
					{RoleId: "admin"},
					{RoleId: "other"},
				},
			}
			r, _ := security.ProfileToJson(p)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, _ := k.Security.FetchProfile(id, nil)

	assert.Equal(t, id, res.Id)
	assert.Equal(t, []*types.Policy{{RoleId: "admin"}, {RoleId: "other"}}, res.Policies)
}

func ExampleFetchProfile() {
	id := "profileId"
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	res, err := k.Security.FetchProfile(id, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Policies)
}
