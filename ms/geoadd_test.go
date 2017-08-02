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

func TestGeoaddEmptyKey(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Geoadd("", []types.GeoPoint{}, qo)

	assert.NotNil(t, err)
	assert.Equal(t, "Ms.Geoadd: key required", fmt.Sprint(err))
}

func TestGeoaddError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	_, err := memoryStorage.Geoadd("foo", []types.GeoPoint{}, qo)

	assert.NotNil(t, err)
}

func TestGeoadd(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "geoadd", parsedQuery.Action)
			assert.Equal(t, "foo", parsedQuery.Id)
			assert.Equal(t, "Montpellier", parsedQuery.Body.(map[string]interface {})["points"].([]interface{})[0].(map[string]interface{})["name"].(string))
			assert.Equal(t, float64(43.6075274), parsedQuery.Body.(map[string]interface {})["points"].([]interface{})[0].(map[string]interface{})["lon"].(float64))
			assert.Equal(t, float64(3.9128795), parsedQuery.Body.(map[string]interface {})["points"].([]interface{})[0].(map[string]interface{})["lat"].(float64))

			r, _ := json.Marshal(1)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	memoryStorage := MemoryStorage.NewMs(k)
	qo := types.NewQueryOptions()

	res, _ := memoryStorage.Geoadd("foo", []types.GeoPoint{{float64(43.6075274), float64(3.9128795), "Montpellier"}}, qo)

	assert.Equal(t, 1, res)
}
