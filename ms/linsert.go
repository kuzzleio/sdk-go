package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Linsert inserts a value in a list, either before or after the reference pivot value.
func (ms Ms) Linsert(key string, position string, pivot string, value string, options types.QueryOptions) (int, error) {
	if position != "before" && position != "after" {
		return -1, types.NewError("Ms.Linsert: invalid position argument (must be 'before' or 'after')", 400)
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
		return -1, res.Error
	}
	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
