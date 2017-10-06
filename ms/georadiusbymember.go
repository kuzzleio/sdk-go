package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

// Georadiusbymember returns the geospatial members of a key inside the provided radius
func (ms Ms) Georadiusbymember(key string, member string, distance float64, unit string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, errors.New("Ms.Georadiusbymember: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options, false, false)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}
	var returnedResults []string
	json.Unmarshal(res.Result, &returnedResults)

	return returnedResults, nil
}

// GeoradiusbymemberWithCoord returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusbymemberWithCoord(key string, member string, distance float64, unit string, options types.QueryOptions) ([]*types.GeoradiusPointWithCoord, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusbymemberWithCoord: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options, true, false)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]*types.GeoradiusPointWithCoord, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i] = &types.GeoradiusPointWithCoord{Name: value[0].(string)}

		tmp := value[1].([]interface{})[0].(string)
		tmpF, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Lon = tmpF

		tmp = value[1].([]interface{})[1].(string)
		tmpF, err = strconv.ParseFloat(tmp, 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Lat = tmpF
	}

	return returnedResults, nil
}

// GeoradiusbymemberWithDist returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusbymemberWithDist(key string, member string, distance float64, unit string, options types.QueryOptions) ([]*types.GeoradiusPointWithDist, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusbymemberWithDist: key required")
	}

	result := make(chan *types.KuzzleResponse)

	query := *types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options, false, true)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]*types.GeoradiusPointWithDist, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i] = &types.GeoradiusPointWithDist{Name: value[0].(string)}

		tmpF, err := strconv.ParseFloat(value[1].(string), 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Dist = tmpF
	}

	return returnedResults, nil
}

// GeoradiusbymemberWithCoordAndDist returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusbymemberWithCoordAndDist(key string, member string, distance float64, unit string, options types.QueryOptions) ([]*types.GeoradiusPointWithCoordAndDist, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusbymemberWithCoordAndDist: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options, true, true)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]*types.GeoradiusPointWithCoordAndDist, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i] = &types.GeoradiusPointWithCoordAndDist{Name: value[0].(string)}

		tmpF, err := strconv.ParseFloat(value[1].(string), 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Dist = tmpF

		tmp := value[2].([]interface{})[0].(string)
		tmpF, err = strconv.ParseFloat(tmp, 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Lon = tmpF

		tmp = value[2].([]interface{})[1].(string)
		tmpF, err = strconv.ParseFloat(tmp, 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Lat = tmpF
	}

	return returnedResults, nil
}
