package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Smove moves a member from a set of unique values to another.
func (ms *Ms) Smove(key string, destination string, member string, options types.QueryOptions) (bool, error) {
	result := make(chan *types.KuzzleResponse)

	type body struct {
		Destination string `json:"destination"`
		Member      string `json:"member"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "smove",
		Id:         key,
		Body:       &body{Destination: destination, Member: member},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult == 1, nil
}
