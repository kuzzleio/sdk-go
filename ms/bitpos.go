package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Bitpos returns the position of the first bit set to 1 or 0 in a string, or in a substring
func (ms Ms) Bitpos(key string, bit int, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Bitpos: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
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

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
