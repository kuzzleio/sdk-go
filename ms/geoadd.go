package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Geoadd deletes all keys from the database
func (ms Ms) Geoadd(key string, points []*types.GeoPoint, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Geoadd: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Points []*types.GeoPoint `json:"points"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "geoadd",
		Id:         key,
		Body:       &body{Points: points},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
