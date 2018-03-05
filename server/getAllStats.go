package server

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//GetAllStats get all Kuzzle usage statistics frames
func (s *Server) GetAllStats(options types.QueryOptions) (json.RawMessage, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "getAllStats",
	}

	go s.k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	return res.Result, nil
}
