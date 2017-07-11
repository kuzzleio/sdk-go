package kuzzle

import (
  "github.com/kuzzleio/sdk-go/internal"
  "testing"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/stretchr/testify/assert"
)

func TestCheckTokenTokenNull(t *testing.T) {
  _, err := kuzzle.CheckToken(nil, "")
  assert.NotNil(t, err)
}

func TestCheckTokenQueryError(t *testing.T) {
  k := &internal.MockedKuzzle{
    MockQuery: func() types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
    },
  }

  _, err := kuzzle.CheckToken(k, "token")
  assert.NotNil(t, err)
}

func TestCheckToken(t *testing.T) {
  k := &internal.MockedKuzzle{
    MockQuery: func() types.KuzzleResponse {
      tokenValidity := kuzzle.TokenValidity{Valid: true}
      r, _ := json.Marshal(tokenValidity)
      return types.KuzzleResponse{Result: r}
    },
  }

  type ackResult struct {
    Acknowledged bool
    ShardsAcknowledged bool
  }
  res, _ := kuzzle.CheckToken(k, "token")
  assert.Equal(t, true, res.Valid)
}

