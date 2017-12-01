package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Sscan is identical to scan, except that sscan iterates the members held by a set of unique values.
func (ms *Ms) Sscan(key string, cursor int, options types.QueryOptions) (*types.MSScanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sscan",
		Id:         key,
		Cursor:     cursor,
	}

	if options != nil {
		if options.Count() != 0 {
			query.Count = options.Count()
		}

		if options.Match() != "" {
			query.Match = options.Match()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var sscanResponse = &types.MSScanResponse{}
	json.Unmarshal(res.Result, sscanResponse)

	return sscanResponse, nil
}
