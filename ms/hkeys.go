package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Hkeys returns all the field names contained in a hash
func (ms Ms) Hkeys(key string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, types.NewError("Ms.Hkeys: key required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hkeys",
		Id:         key,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult []string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
