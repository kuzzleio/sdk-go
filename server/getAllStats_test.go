package server_test

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

func TestGetAllStatsQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	_, err := k.Server.GetAllStats(nil)
	assert.NotNil(t, err)
}

func TestGetAllStats(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "server", request.Controller)
			assert.Equal(t, "getAllStats", request.Action)

			type hits struct {
				Hits []*types.Statistics `json:"hits"`
			}

			m := map[string]int{}
			m["websocket"] = 42

			stats := types.Statistics{
				CompletedRequests: m,
			}

			hitsArray := []*types.Statistics{&stats}
			toMarshal := hits{hitsArray}

			h, err := json.Marshal(toMarshal)
			if err != nil {
				log.Fatal(err)
			}

			return &types.KuzzleResponse{Result: h}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.Server.GetAllStats(nil)

	assert.Equal(t, 42, res[0].CompletedRequests["websocket"])
}

func ExampleKuzzle_GetAllStats() {
	conn := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(conn, nil)

	res, err := k.Server.GetAllStats(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, stat := range res {
		fmt.Println(stat.Timestamp, stat.FailedRequests, stat.Connections, stat.CompletedRequests, stat.OngoingRequests)
	}
}
