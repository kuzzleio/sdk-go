package server

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//GetConfig returns the current Kuzzle configuration.
func (s *Server) GetConfig(options types.QueryOptions) (json.RawMessage, error) {
	result := make(chan *types.KuzzleResponse)
	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "getConfig",
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
