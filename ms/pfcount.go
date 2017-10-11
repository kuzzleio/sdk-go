package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Pfcount returns the probabilistic cardinality of a HyperLogLog data structure,
// or of the merged HyperLogLog structures if more than 1 is provided (see pfadd).
func (ms Ms) Pfcount(keys []string, options types.QueryOptions) (int, error) {
	if len(keys) == 0 {
		return 0, types.NewError("Ms.Pfcount: please provide at least one key")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Elements []string `json:"elements"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pfcount",
		Keys:       keys,
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
