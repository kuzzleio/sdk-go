package auth_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGetMyRightsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getMyRights", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Auth.GetMyRights(nil)
	assert.NotNil(t, err)
}

func TestGetMyRights(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "getMyRights", request.Action)

			type hits struct {
				Hits []*types.UserRights `json:"hits"`
			}

			m := make(map[string]int)
			m["websocket"] = 42

			rights := types.UserRights{
				Controller: "controller",
				Action:     "action",
				Index:      "index",
				Collection: "collection",
				Value:      "allowed",
			}

			hitsArray := []*types.UserRights{&rights}
			toMarshal := hits{hitsArray}

			h, err := json.Marshal(toMarshal)
			if err != nil {
				log.Fatal(err)
			}

			return &types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Auth.GetMyRights(nil)

	assert.Equal(t, "controller", res[0].Controller)
	assert.Equal(t, "action", res[0].Action)
	assert.Equal(t, "index", res[0].Index)
	assert.Equal(t, "collection", res[0].Collection)
	assert.Equal(t, "allowed", res[0].Value)
}

func ExampleKuzzle_GetMyRights() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	myCredentials := credentials{"foo", "bar"}

	_, err := k.Login("local", myCredentials, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res, err := k.Auth.GetMyRights(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
