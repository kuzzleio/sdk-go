package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Pexpire sets a timeout (in milliseconds) on a key.
// After the timeout has expired, the key will automatically be deleted.
func (ms Ms) Pexpire(key string, ttl int, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Milliseconds int `json:"milliseconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "pexpire",
		Id:         key,
		Body:       &body{Milliseconds: ttl},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
