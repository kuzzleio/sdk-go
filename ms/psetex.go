package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Sets a key with the provided value, and an expiration delay expressed in milliseconds. If the key does not exist, it is created beforehand.
*/
func (ms Ms) Psetex(key string, value string, ttl int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Psetex: key required")
	}
	if value == "" {
		return "", errors.New("Ms.Psetex: value required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Value        string `json:"value"`
		Milliseconds int    `json:"milliseconds"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "psetex",
		Id:         key,
		Body:       &body{Value: value, Milliseconds: ttl},
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
