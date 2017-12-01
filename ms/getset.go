package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Getset sets a new value for a key and returns its previous value.
func (ms *Ms) Getset(key string, value string, options types.QueryOptions) (*string, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Value string `json:"value"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "getset",
		Id:         key,
		Body:       &body{Value: value},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var returnedResult *string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
