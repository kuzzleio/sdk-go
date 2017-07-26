package kuzzle_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetAllStatisticsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)
			return types.KuzzleResponse{Error: types.MessageError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.GetAllStatistics(nil)
	assert.NotNil(t, err)
}

func TestGetAllStatistics(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)

			type hits struct {
				Hits []types.Statistics `json:"hits"`
			}

			m := make(map[string]int)
			m["websocket"] = 42

			stats := types.Statistics{
				CompletedRequests: m,
			}

			hitsArray := make([]types.Statistics, 0)
			hitsArray = append(hitsArray, stats)
			toMarshal := hits{hitsArray}

			h, err := json.Marshal(toMarshal)
			if err != nil {
				log.Fatal(err)
			}

			return types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.GetAllStatistics(nil)

	assert.Equal(t, 42, res[0].CompletedRequests["websocket"])
}
