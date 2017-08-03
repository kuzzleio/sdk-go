package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Delete all keys from the database
*/
func (ms Ms) Geoadd(key string, points []types.GeoPoint, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Geoadd: key required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Points []types.GeoPoint `json:"points"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "geoadd",
		Id:         key,
		Body:       &body{Points: points},
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
