package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMappingError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").GetMapping(nil)
	assert.NotNil(t, err)
}

func TestGetMappingUnknownIndex(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "getMapping", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.KuzzleResponse{Result: []byte(`{"index":{"mappings":{"collection":{"properties":{"foo":{"type":"text","properties":{"type":"keyword","ignore_above":255}}}}}}}`)}
			r, _ := json.Marshal(res.Result)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")
	fields := make(map[string]interface{})
	fields["type"] = interface{}("keyword")
	fields["ignore_above"] = interface{}(255.0)

	res, _ := cl.GetMapping(nil)
	assert.Equal(t, collection.CollectionMapping{
		Mapping: types.KuzzleFieldsMapping{
			"foo": {
				Type:       "text",
				Properties: fields,
			},
		},
		Collection: cl,
	}, res)
}

func TestGetMappingUnknownCollection(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "getMapping", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "wrong-collection", parsedQuery.Collection)

			res := types.KuzzleResponse{Result: []byte(`{"index":{"mappings":{"collection":{"properties":{"foo":{"type":"text","properties":{"type":"keyword","ignore_above":255}}}}}}}`)}
			r, _ := json.Marshal(res.Result)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "wrong-collection", "index")

	_, err := cl.GetMapping(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "No mapping found for collection wrong-collection", fmt.Sprint(err))
}

func TestGetMapping(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "getMapping", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.KuzzleResponse{Result: []byte(`{"index":{"mappings":{"collection":{"properties":{"foo":{"type":"text","properties":{"type":"keyword","ignore_above":255}}}}}}}`)}
			r, _ := json.Marshal(res.Result)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")
	fields := make(map[string]interface{})
	fields["type"] = interface{}("keyword")
	fields["ignore_above"] = interface{}(255.0)

	res, _ := cl.GetMapping(nil)
	assert.Equal(t, collection.CollectionMapping{
		Mapping: types.KuzzleFieldsMapping{
			"foo": {
				Type:       "text",
				Properties: fields,
			},
		},
		Collection: cl,
	}, res)
}

func ExampleCollection_GetMapping() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	cl := collection.NewCollection(k, "collection", "index")

	res, err := cl.GetMapping(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Collection, res.Mapping)
}
