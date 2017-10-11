package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// RpoplPush removes the last element of the list at source and pushes it
// back at the start of the list at destination.
func (ms Ms) RpoplPush(source string, destination string, options types.QueryOptions) (interface{}, error) {
	if source == "" {
		return "", types.NewError("Ms.RpoplPush: source required")
	}
	if destination == "" {
		return "", types.NewError("Ms.RpoplPush: destination required")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "rpoplpush",
		Body:       &body{Source: source, Destination: destination},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult interface{}
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
