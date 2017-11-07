package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Psetex sets a key with the provided value, and an expiration delay expressed in milliseconds.
// If the key does not exist, it is created beforehand.
func (ms Ms) Psetex(key string, value string, ttl int, options types.QueryOptions) (string, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value        string `json:"value"`
		Milliseconds int    `json:"milliseconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "psetex",
		Id:         key,
		Body:       &body{Value: value, Milliseconds: ttl},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
