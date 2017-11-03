package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoradiusError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", nil)

	assert.NotNil(t, err)
}

func TestGeoradius(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "some"}, {Name: "results"}}, res)
}

func ExampleMs_Georadius() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestGeoradiusWithCoordLonConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	_, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoordLatConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	_, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithCoord(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	res, _ := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusWithCoord() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	res, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lon, res[0].Lat)
}

func TestGeoradiusWithDistDistConvError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	_, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.NotNil(t, err)
}

func TestGeoradiusWithDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	res, _ := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Dist: 125}}, res)
}

func ExampleMs_GeoradiusWithDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	res, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Dist)
}

func TestGeoradiusWithCoordAndDist(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
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
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true).SetWithcoord(true)

	res, _ := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Dist: 125, Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusWithCoordAndDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithcoord(true).SetWithdist(true)

	res, err := k.MemoryStorage.Georadius("foo", float64(43.6075274), float64(3.9128795), float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lat, res[0].Lon, res[0].Dist)
}
