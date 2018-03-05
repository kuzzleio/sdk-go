package server

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//GetLastStats get Kuzzle usage statistics
func (s *Server) GetLastStats(options types.QueryOptions) (json.RawMessage, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "getLastStats",
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
