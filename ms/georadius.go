package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

func assignGeoradiusOptions(query *types.KuzzleRequest, options types.QueryOptions, withCoord bool, withDist bool) {
	opts := make([]interface{}, 0, 5)

	if options != nil {
		if options.GetCount() != 0 {
			opts = append(opts, "count")
			opts = append(opts, options.GetCount())
		}

		if options.GetSort() != "" {
			opts = append(opts, options.GetSort())
		}
	}

	if withCoord {
		opts = append(opts, "withcoord")
	}

	if withDist {
		opts = append(opts, "withdist")
	}

	query.Options = []interface{}(opts)
}

// Georadius returns the geospatial members of a key inside the provided radius
func (ms Ms) Georadius(key string, lon float64, lat float64, distance float64, unit string, options types.QueryOptions) ([]string, error) {
	if key == "" {
		return nil, errors.New("Ms.Georadius: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadius",
		Id:         key,
		Lon:        lon,
		Lat:        lat,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(&query, options, false, false)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var returnedResults []string
	json.Unmarshal(res.Result, &returnedResults)

	return returnedResults, nil
}

// GeoradiusWithCoord returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusWithCoord(key string, lon float64, lat float64, distance float64, unit string, options types.QueryOptions) ([]types.GeoradiusPointWithCoord, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusWithCoord: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadius",
		Id:         key,
		Lon:        lon,
		Lat:        lat,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(&query, options, true, false)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]types.GeoradiusPointWithCoord, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i].Name = value[0].(string)

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

// GeoradiusWithDist returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusWithDist(key string, lon float64, lat float64, distance float64, unit string, options types.QueryOptions) ([]types.GeoradiusPointWithDist, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusWithDist: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadius",
		Id:         key,
		Lon:        lon,
		Lat:        lat,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(&query, options, false, true)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]types.GeoradiusPointWithDist, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i].Name = value[0].(string)

		tmpF, err := strconv.ParseFloat(value[1].(string), 64)
		if err != nil {
			return nil, err
		}

		returnedResults[i].Dist = tmpF
	}

	return returnedResults, nil
}

// GeoradiusWithCoordAndDist returns the geospatial members of a key inside the provided radius
func (ms Ms) GeoradiusWithCoordAndDist(key string, lon float64, lat float64, distance float64, unit string, options types.QueryOptions) ([]types.GeoradiusPointWithCoordAndDist, error) {
	if key == "" {
		return nil, errors.New("Ms.GeoradiusWithCoordAndDist: key required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadius",
		Id:         key,
		Lon:        lon,
		Lat:        lat,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(&query, options, true, true)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}
	var stringResults [][]interface{}
	json.Unmarshal(res.Result, &stringResults)

	returnedResults := make([]types.GeoradiusPointWithCoordAndDist, len(stringResults))

	for i, value := range stringResults {
		returnedResults[i].Name = value[0].(string)

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
