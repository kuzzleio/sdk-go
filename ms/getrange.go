package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// Getrange returns a substring of a key's value (index starts at position 0).
func (ms Ms) Getrange(key string, start int, end int, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Getrange: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "getrange",
		Id:         key,
		Start:      &start,
		End:        end,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
