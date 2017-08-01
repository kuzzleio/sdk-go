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
func (ms Ms) Geopos(key string, members []string, options types.QueryOptions) ([]types.GeoPoint, error) {
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

	returnedResults := make([]types.GeoPoint, len(stringResults))

	for i := 0; i < len(stringResults); i++ {
		returnedResults[i] = types.GeoPoint{}
		tmp, err := strconv.ParseFloat(stringResults[i][0], 64)
		if err != nil {
			return nil, err
		}
		returnedResults[i].Lon = tmp

		tmp, err = strconv.ParseFloat(stringResults[i][1], 64)
		if err != nil {
			return nil, err
		}
		returnedResults[i].Lat = tmp

		returnedResults[i].Name = members[i]
	}

	return returnedResults, nil
}