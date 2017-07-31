package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

func (ms Ms) Bitcount(key string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, errors.New("Ms.Bitcount: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "bitcount",
		Id:         key,
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
	var countResult int
	json.Unmarshal(res.Result, &countResult)

	return countResult, nil
}
