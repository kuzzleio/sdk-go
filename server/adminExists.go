package server

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

//AdminExists checks if an administrator account has been created, and return a boolean as a result.
func (s *Server) AdminExists(options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "server",
		Action:     "adminExists",
	}

	go s.k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	type exists struct {
		Exists bool `json:"exists"`
	}

	r := exists{}
	json.Unmarshal(res.Result, &r)
	return r.Exists, nil
}
