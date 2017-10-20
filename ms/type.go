package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Type returns the type of the value held by a key.
func (ms Ms) Type(key string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", types.NewError("Ms.Type: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "type",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return "", res.Error
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
