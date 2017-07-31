package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
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

	fmt.Println(string(res.Result))

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]string
	json.Unmarshal(res.Result, &stringResults)

	var returnedResults [][]float64

	for i := 0; i < len(stringResults); i++ {
		for j := 0; j < 2; j++ {
			tmp, err := strconv.ParseFloat(stringResults[j][0], 64)
			if err != nil {
				return nil, err
			}
			returnedResults[j][0] = tmp
		}
	}

	return returnedResults, nil
}
