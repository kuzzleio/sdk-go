package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryOptions(t *testing.T) {
	var k *Kuzzle

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)

			assert.Equal(t, k.version, request.Volatile["sdkVersion"])
			assert.Equal(t, 0, request.From)
			assert.Equal(t, 10, request.Size)
			assert.Equal(t, "1m", request.Scroll)
			assert.Equal(t, "", request.ScrollId)

			return types.KuzzleResponse{}
		},
	}
	k, _ = NewKuzzle(c, nil)

	ch := make(chan types.KuzzleResponse)
	options := types.DefaultOptions()
	go k.Query(types.KuzzleRequest{}, options, ch)
	<-ch
}
