package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

type ZScanResponse struct {
	Cursor int
	Values []string
}

/*
  Identical to scan, except that zscan iterates the members held by a sorted set.
*/
func (ms Ms) Zscan(key string, cursor *int, options types.QueryOptions) (types.MSScanResponse, error) {
	if key == "" {
		return types.MSScanResponse{}, errors.New("Ms.Zscan: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "zscan",
		Id:         key,
		Cursor: 		cursor,
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return types.MSScanResponse{}, errors.New(res.Error.Message)
	}

	var scanResponse []interface{}
	json.Unmarshal(res.Result, &scanResponse)

	return formatZscanResponse(scanResponse), nil
}

func formatZscanResponse(response []interface{}) types.MSScanResponse {
	formatedResponse := types.MSScanResponse{}

	for _, element := range response {
		switch vf := element.(type) {
		case string:
			formatedResponse.Cursor, _ = strconv.Atoi(vf)
		case []interface{}:
			values := []string{}

			for _, v := range vf {
				switch vv := v.(type) {
				case string:
					values = append(values, vv)
				}
			}

			formatedResponse.Values = values
		}
	}

	return formatedResponse
}
