package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// RefreshIndex forces the provided data index to refresh on each modification
func (k Kuzzle) RefreshIndex(index string, options types.QueryOptions) (*types.Shards, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "refresh",
		Index:      index,
	}
	go k.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	type s struct {
		Shards *types.Shards `json:"_shards"`
	}

	shards := s{}

	json.Unmarshal(res.Result, &shards)

	return shards.Shards, nil
}
