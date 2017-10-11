package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	MemoryStorage "github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoradiusbymemberEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Georadiusbymember("", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Georadiusbymember: key required", fmt.Sprint(err))
}

func TestGeoradiusbymemberError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymember(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)

			var opts []interface{}
			opts = append(opts, "count")
			opts = append(opts, float64(42))
			opts = append(opts, "ASC")

			assert.Equal(t, opts, parsedQuery.Options)

			r, _ := json.Marshal([]string{"some", "results"})
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []string{"some", "results"}, res)
}

func ExampleMs_Georadiusbymember() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := memoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestGeoradiusbymemberWithCoordEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithCoord("", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusbymemberWithCoord: key required", fmt.Sprint(err))
}

func TestGeoradiusbymemberWithCoordError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithCoord("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoordLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithCoord("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoordLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithCoord("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoord(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusbymemberWithCoord("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.GeoradiusPointWithCoord{{Name: "Montpellier", Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusbymemberWithCoord() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := memoryStorage.GeoradiusbymemberWithCoord("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lat, res[0].Lon)
}

func TestGeoradiusbymemberWithDistEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithDist("", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusbymemberWithDist: key required", fmt.Sprint(err))
}

func TestGeoradiusbymemberWithDistError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithDistDistConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusbymemberWithDist("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.GeoradiusPointWithDist{{Name: "Montpellier", Dist: 125}}, res)
}

func ExampleMs_GeoradiusbymemberWithDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := memoryStorage.GeoradiusbymemberWithDist("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Dist)
}

func TestGeoradiusbymemberWithCoordAndDistEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.GeoradiusbymemberWithCoordAndDist: key required", fmt.Sprint(err))
}

func TestGeoradiusbymemberWithCoordAndDistError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}
func TestGeoradiusbymemberWithCoordAndDistDistConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoordAndDistLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoordAndDistLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	_, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusbymemberWithCoordAndDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "georadiusbymember", parsedQuery.Action)
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.GeoradiusPointWithCoordAndDist{{Name: "Montpellier", Dist: 125, Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusbymemberWithCoordAndDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := memoryStorage.GeoradiusbymemberWithCoordAndDist("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lat, res[0].Lon, res[0].Dist)
}
