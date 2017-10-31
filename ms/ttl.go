package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Ttl returns the remaining time to live of a key, in seconds,
// or a negative value if the key does not exist or if it is persistent.
func (ms Ms) Ttl(key string, options types.QueryOptions) (int, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys []string `json:"keys"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "ttl",
		Id:         key,
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
