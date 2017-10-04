package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// SetEx sets a key with the provided value, and an expiration delay expressed in seconds.
// If the key does not exist, it is created beforehand.
func (ms Ms) SetEx(key string, value interface{}, ttl int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.SetEx: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value interface{} `json:"value"`
		Ttk   int         `json:"seconds"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "setex",
		Id:         key,
		Body:       &body{Value: value, Ttk: ttl},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
