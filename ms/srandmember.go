package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns one or more members of a set of unique values, at random.
  If count is provided and is positive, the returned values are unique. If count is negative, a set member can be returned multiple times.
*/
func (ms Ms) SrandMember(key string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return []string{}, errors.New("Ms.SrandMember: key required")
	}

	if options == nil || options.GetCount() == 0 {
		options.SetCount(1)
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "srandmember",
		Id:         key,
	}

	if options != nil {
		if options.GetCount() != 0 {
			query.Count = options.GetCount()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return []string{}, errors.New(res.Error.Message)
	}

	if options.GetCount() == 1 {
		var returnedResult string
		json.Unmarshal(res.Result, &returnedResult)

		return []string{returnedResult}, nil
	} else {
		var returnedResult []string
		json.Unmarshal(res.Result, &returnedResult)

		return returnedResult, nil
	}
}
