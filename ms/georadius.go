package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"strconv"
)

func assignGeoradiusOptions(query *types.KuzzleRequest, options types.QueryOptions) {
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

	if options.GetWithcoord() {
		opts = append(opts, "withcoord")
	}

	if options.GetWithdist() {
		opts = append(opts, "withdist")
	}

	query.Options = []interface{}(opts)
}

func responseToGeoradius(response *types.KuzzleResponse, options types.QueryOptions) ([]*types.Georadius, error) {
	var stringResults []interface{}

	json.Unmarshal(response.Result, &stringResults)
	returnedResults := make([]*types.Georadius, len(stringResults))

	for i, value := range stringResults {
		var err error

		// if none of the 2 options below are provided, then we have
		// a simple array of strings and not an array of arrays
		if !options.GetWithdist() && !options.GetWithcoord() {
			returnedResults[i] = &types.Georadius{Name: value.(string)}
		} else {
			returnedResults[i] = &types.Georadius{Name: value.([]interface{})[0].(string)}
		}

		if options.GetWithdist() {
			returnedResults[i].Dist, err = strconv.ParseFloat(value.([]interface{})[1].(string), 64)
			if err != nil {
				return nil, types.NewError(err.Error())
			}
		}

		if options.GetWithcoord() {
			coordstart := 1

			if (options.GetWithdist()) {
				coordstart++
			}

			tmp := value.([]interface{})[coordstart].([]interface{})[0].(string)
			returnedResults[i].Lon, err = strconv.ParseFloat(tmp, 64)

			if err != nil {
				return nil, types.NewError(err.Error())
			}

			tmp = value.([]interface{})[coordstart].([]interface{})[1].(string)
			returnedResults[i].Lat, err = strconv.ParseFloat(tmp, 64)

			if err != nil {
				return nil, types.NewError(err.Error())
			}
		}
	}

	return returnedResults, nil
}

// Georadius returns the geospatial members of a key inside the provided radius
func (ms Ms) Georadius(key string, lon float64, lat float64, distance float64, unit string, options types.QueryOptions) ([]*types.Georadius, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadius",
		Id:         key,
		Lon:        lon,
		Lat:        lat,
		Distance:   distance,
		Unit:       unit,
	}

	assignGeoradiusOptions(query, options)

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return responseToGeoradius(res, options)	
}
