package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Get the distance between two geospatial members of a key (see geoadd)
*/
func (ms Ms) Geodist(key string, member1 string, member2 string, options types.QueryOptions) (float64, error) {
	if key == "" {
		return 0, errors.New("Ms.Geodist: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "geodist",
		Id:         key,
		Member1:    member1,
		Member2:    member2,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult float64
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
