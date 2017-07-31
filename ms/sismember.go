package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Checks if member is a member of the set of unique values stored at key.
*/
func (ms Ms) SisMember(key string, member string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.SisMember: key required")
	}
	if member == "" {
		return 0, errors.New("Ms.SisMember: member required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sismember",
		Id:         key,
		Member:     member,
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
