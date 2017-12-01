package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Geopos returns the longitude/latitude values for the provided key's members
func (ms *Ms) Geopos(key string, members []string, options types.QueryOptions) ([]*types.GeoPoint, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "geopos",
		Id:         key,
		Members:    members,
	}
	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}
	var stringResults [][]string
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]*types.GeoPoint, len(stringResults))

	for i := 0; i < len(stringResults); i++ {
		returnedResults[i] = &types.GeoPoint{}
		tmp, err := strconv.ParseFloat(stringResults[i][0], 64)
		if err != nil {
			return nil, types.NewError(err.Error())
		}
		returnedResults[i].Lon = tmp

		tmp, err = strconv.ParseFloat(stringResults[i][1], 64)
		if err != nil {
			return nil, types.NewError(err.Error())
		}
		returnedResults[i].Lat = tmp

		returnedResults[i].Name = members[i]
	}

	return returnedResults, nil
}
