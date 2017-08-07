package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the probabilistic cardinality of a HyperLogLog data structure, or of the merged HyperLogLog structures if more than 1 is provided (see pfadd).
*/
func (ms Ms) Pfcount(keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, errors.New("Ms.Pfcount: please provide at least one key")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Elements []string `json:"elements"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "pfcount",
		Keys:       keys,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
