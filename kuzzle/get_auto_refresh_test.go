package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAutoRefreshDefaultIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	_, err := k.GetAutoRefresh("", nil)
	assert.NotNil(t, err)
}

func TestGetAutoRefreshIndexNull(t *testing.T) {
	opts := types.NewOptions()
	opts.SetDefaultIndex("myIndex")

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Result: json.RawMessage("true")}
		},
	}

	k, _ := kuzzle.NewKuzzle(c, opts)
	_, err := k.GetAutoRefresh("", nil)
	assert.Nil(t, err)
}

func TestGetAutoRefreshQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.GetAutoRefresh("index", nil)
	assert.NotNil(t, err)
}

func TestGetAutoRefresh(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Result: json.RawMessage("true")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	type ackResult struct {
		Acknowledged       bool
		ShardsAcknowledged bool
	}
	res, _ := k.GetAutoRefresh("index", nil)
	assert.Equal(t, true, res)
}
