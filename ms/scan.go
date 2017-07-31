package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type ScanResponse struct {
	Cursor int      `json:"cursor"`
	Values []string `json:"values"`
}

/*
  Iterates incrementally the set of keys in the database using a cursor.

	An iteration starts when the cursor is set to 0.
	To get the next page of results, simply re-send the identical request with the updated cursor position provided in the result set.
	The scan terminates when the next position cursor returned by the server is 0.
*/
func (ms Ms) Scan(cursor *int, options types.QueryOptions) (ScanResponse, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "scan",
		Cursor: 		cursor,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return ScanResponse{}, errors.New(res.Error.Message)
	}

	var scanResponse = ScanResponse{}
	json.Unmarshal(res.Result, &scanResponse)

	return scanResponse, nil
}
