package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Returns the position of the first bit set to 1 or 0 in a string, or in a substring
*/
func (ms Ms) Bitpos(key string, bit int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Bitpos: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "bitpos",
		Id:         key,
		Bit:        bit,
	}

	if options != nil {
		if options.GetStart() != 0 {
			query.Start = options.GetStart()
		}

		if options.GetEnd() != 0 {
			query.End = options.GetEnd()
		}
	}

	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return 0, errors.New(res.Error.Message)
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
