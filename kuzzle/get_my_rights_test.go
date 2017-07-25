package kuzzle_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "log"
)

func TestGetMyRightsQueryError(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getMyRights", request.Action)
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)
  _, err := k.GetMyRights(nil)
  assert.NotNil(t, err)
}

func TestGetMyRights(t *testing.T) {
  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      request := types.KuzzleRequest{}
      json.Unmarshal(query, &request)
      assert.Equal(t, "auth", request.Controller)
      assert.Equal(t, "getMyRights", request.Action)

      type hits struct {
        Hits []types.Rights `json:"hits"`
      }

      m := make(map[string]int)
      m["websocket"] = 42

      rights := types.Rights{
        Controller: "controller",
        Action: "action",
        Index: "index",
        Collection: "collection",
        Value: "allowed",
      }

      hitsArray := make([]types.Rights, 0)
      hitsArray = append(hitsArray, rights)
      toMarshal := hits{hitsArray}

      h, err := json.Marshal(toMarshal)
      if err != nil {
        log.Fatal(err)
      }

      return types.KuzzleResponse{Result: h}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := k.GetMyRights(nil)

  assert.Equal(t, "controller", res[0].Controller)
  assert.Equal(t, "action", res[0].Action)
  assert.Equal(t, "index", res[0].Index)
  assert.Equal(t, "collection", res[0].Collection)
  assert.Equal(t, "allowed", res[0].Value)
}
