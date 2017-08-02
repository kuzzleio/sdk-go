package ms_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"testing"
	"fmt"
)

func TestGeoradiusEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Georadius("", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Georadius: key required", fmt.Sprint(err))
}

func TestGeoradiusError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradius(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")

			assert.Equal(t, opts, parsedQuery.Options)

			r, _ := json.Marshal([]string{"some", "results"})
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []string{"some", "results"}, res)
}

func TestGeoradiusWithCoordEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithCoord("", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusWithCoord: key required", fmt.Sprint(err))
}

func TestGeoradiusWithCoordError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithCoord("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoordLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075abc")
			point = append(point, "3.9128795")
			location = append(location, "Montpellier")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithCoord("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoordLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075274")
			point = append(point, "3.9128abc")
			location = append(location, "Montpellier")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithCoord("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}



func TestGeoradiusWithCoord(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075274")
			point = append(point, "3.9128795")
			location = append(location, "Montpellier")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusWithCoord("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []types.GeoradiusPointWithCoord{{Name:"Montpellier", Lon:43.6075274, Lat:3.9128795}}, res)
}

func TestGeoradiusWithDistEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithDist("", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusWithDist: key required", fmt.Sprint(err))
}

func TestGeoradiusWithDistError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithDistDistConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}

			location = append(location, "Montpellier")
			location = append(location, "125.23abc")
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}

			location = append(location, "Montpellier")
			location = append(location, "125")
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusWithDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []types.GeoradiusPointWithDist{{Name:"Montpellier", Dist: 125}}, res)
}


func TestGeoradiusWithCoordAndDistEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithCoordAndDist("", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusWithCoordAndDist: key required", fmt.Sprint(err))
}

func TestGeoradiusWithCoordAndDistError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusWithCoordAndDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}
func TestGeoradiusWithCoordAndDistDistConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075274")
			point = append(point, "3.9128795")
			location = append(location, "Montpellier")
			location = append(location, "125.23abc")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithCoordAndDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoordAndDistLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075abc")
			point = append(point, "3.9128795")
			location = append(location, "Montpellier")
			location = append(location, "125")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithCoordAndDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoordAndDistLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075274")
			point = append(point, "3.9128abc")
			location = append(location, "Montpellier")
			location = append(location, "125")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusWithCoordAndDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}


func TestGeoradiusWithCoordAndDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadius", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")
			opts = append(opts, "withcoord")
			opts = append(opts, "withdist")

			assert.Equal(t, opts, parsedQuery.Options)

			var response [][]interface{}
			var location []interface{}
			var point []interface{}

			point = append(point, "43.6075274")
			point = append(point, "3.9128795")
			location = append(location, "Montpellier")
			location = append(location, "125")
			location = append(location, point)
			response = append(response, location)

			r, _ := json.Marshal(response)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusWithCoordAndDist("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []types.GeoradiusPointWithCoordAndDist{{Name:"Montpellier", Dist: 125, Lon:43.6075274, Lat:3.9128795}}, res)
}
