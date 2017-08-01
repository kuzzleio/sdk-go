package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type SscanResponse struct {
	Cursor int      `json:"cursor"`
	Values []string `json:"values"`
}

/*
  Identical to scan, except that sscan iterates the members held by a set of unique values.
*/
func (ms Ms) Sscan(key string, cursor *int, options types.QueryOptions) (SscanResponse, error) {
	if key == "" {
		return SscanResponse{}, errors.New("Ms.Sscan: key required")
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
		return SscanResponse{}, errors.New(res.Error.Message)
	}

	var sscanResponse = SscanResponse{}
	json.Unmarshal(res.Result, &sscanResponse)

	return sscanResponse, nil
}
