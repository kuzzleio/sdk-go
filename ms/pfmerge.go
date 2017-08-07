package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Merges multiple HyperLogLog data structures into an unique HyperLogLog structure stored at key, approximating the cardinality of the union of the source structures.
*/
func (ms Ms) Pfmerge(key string, sources []string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Pfmerge: key required")
	}
	if len(sources) == 0 {
		return "", errors.New("Ms.Pfmerge: please provide at least one source")
	}

	result := make(chan types.KuzzleResponse)

	type body struct {
		Sources []string `json:"sources"`
	}

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "pfmerge",
		Id:         key,
		Body:       &body{Sources: sources},
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
