package collection_test

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/collection"

	"testing"

	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateSpecificationsBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	nc := collection.NewCollection(k)
	_, err := nc.ValidateSpecifications(nil, nil)
	assert.NotNil(t, err)
}

func TestValidateSpecificationsError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	_, err := nc.ValidateSpecifications(json.RawMessage("body"), nil)
	assert.NotNil(t, err)
}

func TestValidateSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Result: []byte(`{
					"valid": true,
					"details": [],
					"description": "Some description text"
				}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	res, err := nc.ValidateSpecifications(json.RawMessage("body"), nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res)
}

func ExampleCollection_ValidateSpecifications() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	nc := collection.NewCollection(k)
	res, err := nc.ValidateSpecifications(json.RawMessage("body"), nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
