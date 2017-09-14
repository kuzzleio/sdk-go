package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Geohash returns the geohash values for the provided key's members
func (ms Ms) Geohash(key string, members []string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, errors.New("Ms.Geohash: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "geohash",
		Id:         key,
		Members:    members,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
