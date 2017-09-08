package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryDefaultOptions(t *testing.T) {
	var k *Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, k.version, request.Volatile["sdkVersion"])
			assert.Equal(t, 0, request.From)
			assert.Equal(t, 10, request.Size)
			assert.Equal(t, "", request.Scroll)
			assert.Equal(t, "", request.ScrollId)

			return types.KuzzleResponse{}
		},
	}
	k, _ = NewKuzzle(c, nil)

	ch := make(chan types.KuzzleResponse)
	options := types.NewQueryOptions()
	go k.Query(types.KuzzleRequest{}, options, ch)
	<-ch
}

func TestQueryWithOptions(t *testing.T) {
	var k *Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, k.version, request.Volatile["sdkVersion"])
			assert.Equal(t, 42, request.From)
			assert.Equal(t, 24, request.Size)
			assert.Equal(t, "5m", request.Scroll)
			assert.Equal(t, "f00b4r", request.ScrollId)

			rawRequest := map[string]interface{}{}
			json.Unmarshal(query, &rawRequest)

			assert.Equal(t, "wait_for", rawRequest["refresh"])
			assert.Equal(t, "wait_for", rawRequest["refresh"])
			assert.Equal(t, 7.0, rawRequest["retryOnConflict"])

			return types.KuzzleResponse{}
		},
	}
	k, _ = NewKuzzle(c, nil)

	ch := make(chan types.KuzzleResponse)
	options := types.NewQueryOptions()

	options.SetFrom(42)
	options.SetSize(24)
	options.SetScroll("5m")
	options.SetScrollId("f00b4r")
	options.SetRefresh("wait_for")
	options.SetRetryOnConflict(7)

	k.headers = map[string]interface{}{"random": "header"}

	go k.Query(types.KuzzleRequest{}, options, ch)
	<-ch
}

func TestQueryWithVolatile(t *testing.T) {
	var k *Kuzzle
	var volatileData = types.VolatileData{
		"modifiedBy": "awesome me",
		"reason": "it needed to be modified",
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, volatileData, request.Volatile)
			assert.Equal(t, k.version, request.Volatile["sdkVersion"])

			return types.KuzzleResponse{}
		},
	}
	k, _ = NewKuzzle(c, nil)

	ch := make(chan types.KuzzleResponse)
	options := types.NewQueryOptions()
	options.SetVolatile(volatileData)
	go k.Query(types.KuzzleRequest{}, options, ch)
	<-ch
}
