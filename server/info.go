package server

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//Info retrieves information about Kuzzle, its plugins and active services.
func (s *Server) Info(options types.QueryOptions) (json.RawMessage, error) {
	result := make(chan *types.KuzzleResponse)
	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "info",
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
