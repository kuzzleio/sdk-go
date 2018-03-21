package collection_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSearchSpecificationsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	_, err := nc.SearchSpecifications(nil)
	assert.NotNil(t, err)
}

func TestSearchSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "searchSpecifications", parsedQuery.Action)

			res := types.SpecificationSearchResult{
				ScrollId: "f00b4r",
				Total:    1,
				Hits:     make([]types.SpecificationSearchResultHit, 1),
			}
			res.Hits[0] = types.SpecificationSearchResultHit{
				Source: types.SpecificationEntry{
					Index:      "index",
					Collection: "collection",
					Validation: &types.Specification{
						Strict: false,
						Fields: types.SpecificationFields{
							"foo": types.SpecificationField{
								Mandatory:    true,
								Type:         "string",
								DefaultValue: "Value found with search",
							},
						},
					},
				},
			}

			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	res, err := nc.SearchSpecifications(nil)
	assert.Equal(t, 1, res.Total)
	assert.Nil(t, err)
}

func ExampleCollection_SearchSpecifications() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)

	res, err := nc.SearchSpecifications(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
