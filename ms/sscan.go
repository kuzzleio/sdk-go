package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Identical to scan, except that sscan iterates the members held by a set of unique values.
*/
func (ms Ms) Sscan(key string, cursor *int, options types.QueryOptions) (types.MSScanResponse, error) {
	if key == "" {
		return types.MSScanResponse{}, errors.New("Ms.Sscan: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "sscan",
		Id:         key,
		Cursor: 		cursor,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return types.MSScanResponse{}, errors.New(res.Error.Message)
	}

	var sscanResponse = types.MSScanResponse{}
	json.Unmarshal(res.Result, &sscanResponse)

	return sscanResponse, nil
}
