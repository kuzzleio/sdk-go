package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsetJwt(t *testing.T) {
	var k *Kuzzle
	renewcalled := false

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "login", request.Action)
			assert.Equal(t, 0, request.ExpiresIn)

			type loginResult struct {
				Jwt string `json:"jwt"`
			}

			loginRes := loginResult{"token"}
			marsh, _ := json.Marshal(loginRes)

			return types.KuzzleResponse{Result: marsh}
		},
		MockGetRooms: func() map[string]map[string]types.IRoom {
			rooms := make(map[string]map[string]types.IRoom)

			room := make(map[string]types.IRoom)
			newRoom := internal.MockedRoom{
				MockedRenew: func() {
					renewcalled = true
				},
			}

			room["id"] = newRoom
			rooms["roomId"] = room
			return rooms
		},
	}

	k, _ = NewKuzzle(c, nil)

	res, _ := k.Login("local", nil, nil)
	assert.Equal(t, "token", res)
	assert.Equal(t, "token", k.jwt)
	k.UnsetJwt()
	assert.Equal(t, "", k.jwt)
	assert.Equal(t, true, renewcalled)
}
