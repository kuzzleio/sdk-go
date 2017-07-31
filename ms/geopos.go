package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

/*
  Return the longitude/latitude values for the provided key's members
*/
func (ms Ms) Geopos(key string, members []string, options types.QueryOptions) ([][]float64, error) {
	if key == "" {
		return nil, errors.New("Ms.Geopos: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "geopos",
		Id: 				key,
		Members:    members,
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]string
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([][]float64, len(stringResults))

	for i := 0; i < len(stringResults); i++ {
		returnedResults[i] = make([]float64, 2)
		for j := 0; j < 2; j++ {
			tmp, err := strconv.ParseFloat(stringResults[i][j], 64)
			if err != nil {
				return nil, err
			}
			returnedResults[i][j] = tmp
		}
	}

	return returnedResults, nil
}
