package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Smove moves a member from a set of unique values to another.
func (ms Ms) Smove(key string, destination string, member string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Smove: key required")
	}
	if destination == "" {
		return 0, errors.New("Ms.Smove: destination required")
	}
	if member == "" {
		return 0, errors.New("Ms.Smove: member required")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Destination string `json:"destination"`
		Member      string `json:"member"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "smove",
		Id:         key,
		Body:       &body{Destination: destination, Member: member},
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
