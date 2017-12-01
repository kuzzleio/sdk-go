package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

type HscanResponse struct {
	Cursor int
	Values []string
}

// Hscan is identical to scan, except that hscan iterates the fields contained in a hash.
func (ms *Ms) Hscan(key string, cursor int, options types.QueryOptions) (*HscanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "hscan",
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

	var stringResult []interface{}
	json.Unmarshal(res.Result, &stringResult)

	returnedResult := &HscanResponse{}

	tmp, err := strconv.ParseInt(stringResult[0].(string), 10, 0)
	if err != nil {
		return returnedResult, types.NewError(err.Error())
	}
	returnedResult.Cursor = int(tmp)

	tmpS := stringResult[1].([]interface{})

	for _, value := range tmpS {
		returnedResult.Values = append(returnedResult.Values, value.(string))
	}

	return returnedResult, nil
}
