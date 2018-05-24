// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func TestGeoradiusbymemberError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", nil)

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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42)

	res, _ := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "some"}, {Name: "results"}}, res)
}

func ExampleMs_Georadiusbymember() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42)

	res, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	_, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	_, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	res, _ := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusbymemberWithCoord() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithcoord(true)

	res, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lat, res[0].Lon)
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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	_, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	res, _ := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Dist: 125}}, res)
}

func ExampleMs_GeoradiusbymemberWithDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithdist(true)

	res, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Dist)
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
	qo := types.NewQueryOptions()
	qo.SetSort("ASC").SetCount(42).SetWithdist(true).SetWithcoord(true)

	res, _ := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	assert.Equal(t, []*types.Georadius{{Name: "Montpellier", Dist: 125, Lon: 43.6075274, Lat: 3.9128795}}, res)
}

func ExampleMs_GeoradiusbymemberWithCoordAndDist() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetSort("ASC").SetCount(42).SetWithcoord(true).SetWithdist(true)

	res, err := k.MemoryStorage.Georadiusbymember("foo", "member", float64(200), "km", qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Name, res[0].Lat, res[0].Lon, res[0].Dist)
}
