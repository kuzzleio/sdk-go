package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

type HscanResponse struct {
	Cursor int
	Values []string
}

// Hscan is identical to scan, except that hscan iterates the fields contained in a hash.
func (ms Ms) Hscan(key string, cursor *int, options types.QueryOptions) (HscanResponse, error) {
	if key == "" {
		return HscanResponse{}, errors.New("Ms.Hscan: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "hscan",
		Id:         key,
		Cursor:     cursor,
	}

	if options != nil {
		if options.GetCount() != 0 {
			query.Count = options.GetCount()
		}

		if options.GetMatch() != "" {
			query.Match = options.GetMatch()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return HscanResponse{}, errors.New(res.Error.Message)
	}

	var stringResult []interface{}
	json.Unmarshal(res.Result, &stringResult)

	returnedResult := HscanResponse{}

	tmp, err := strconv.ParseInt(stringResult[0].(string), 10, 0)
	if err != nil {
		return HscanResponse{}, err
	}
	returnedResult.Cursor = int(tmp)

	tmpS := stringResult[1].([]interface{})

	for _, value := range tmpS {
		returnedResult.Values = append(returnedResult.Values, value.(string))
	}

	return returnedResult, nil
}
