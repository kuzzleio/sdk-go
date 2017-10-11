package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Linsert inserts a value in a list, either before or after the reference pivot value.
func (ms Ms) Linsert(key string, position string, pivot string, value string, options types.QueryOptions) (int, error) {
	if key == "" {
		return 0, types.NewError("Ms.Linsert: key required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Position string `json:"position"`
		Pivot    string `json:"pivot"`
		Value    string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "linsert",
		Id:         key,
		Body:       &body{Position: position, Pivot: pivot, Value: value},
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
