package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Ttl returns the remaining time to live of a key, in seconds,
// or a negative value if the key does not exist or if it is persistent.
func (ms Ms) Ttl(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Ttl: key required")
	}

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

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
