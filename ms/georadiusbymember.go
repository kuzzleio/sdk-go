package ms

import (
	"github.com/kuzzleio/sdk-go/types"
)

// Georadiusbymember returns the geospatial members of a key inside the provided radius
func (ms *Ms) Georadiusbymember(key string, member string, distance float64, unit string, options types.QueryOptions) ([]*types.Georadius, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "georadiusbymember",
		Id:         key,
		Member:     member,
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
