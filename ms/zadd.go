package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Adds the specified elements to the sorted set stored at key. If the key does not exist, it is created, holding an empty sorted set. If it already exists and does not hold a sorted set, an error is returned.

  Scores are expressed as floating point numbers.

  If a member to insert is already in the sorted set, its score is updated and the member is reinserted at the right position in the set.
*/
func (ms Ms) Zadd(key string, elements []types.MSSortedSet, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Zadd: key required")
	}
	if len(elements) == 0 {
		return 0, errors.New("Ms.Zadd: please provide at least one element")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Elements []types.MSSortedSet `json:"elements"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zadd",
		Id:         key,
		Body:       &body{Elements: elements},
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
