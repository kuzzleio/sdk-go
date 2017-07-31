package collection_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").GetSpecifications(nil)
	assert.NotNil(t, err)
}

func TestGetSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "getSpecifications", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			validation := types.KuzzleValidation{
				Strict: false,
				Fields: types.KuzzleValidationFields{
					"foo": {
						Mandatory:    false,
						Type:         "bool",
						DefaultValue: "Boring value",
					},
				},
			}

			res := types.KuzzleSpecificationsResult{
				Index:      parsedQuery.Index,
				Collection: parsedQuery.Collection,
				Validation: validation,
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").GetSpecifications(nil)
	assert.Equal(t, "index", res.Index)
	assert.Equal(t, "collection", res.Collection)
	assert.Equal(t, false, res.Validation.Strict)
	assert.Equal(t, 1, len(res.Validation.Fields))
	assert.Equal(t, false, res.Validation.Fields["foo"].Mandatory)
	assert.Equal(t, "bool", res.Validation.Fields["foo"].Type)
	assert.Equal(t, "Boring value", res.Validation.Fields["foo"].DefaultValue)
}

func TestSearchSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").SearchSpecifications(nil, nil)
	assert.NotNil(t, err)
}

func TestSearchSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "searchSpecifications", parsedQuery.Action)

			res := types.KuzzleSpecificationSearchResult{
				ScrollId: "f00b4r",
				Total:    1,
				Hits: []struct {
					Source types.KuzzleSpecificationsResult `json:"_source"`
				}{{Source: types.KuzzleSpecificationsResult{
					Index:      "index",
					Collection: "collection",
					Validation: types.KuzzleValidation{
						Strict: false,
						Fields: types.KuzzleValidationFields{
							"foo": {
								Mandatory:    true,
								Type:         "string",
								DefaultValue: "Value found with search",
							},
						},
					},
				}}},
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")

	res, _ := collection.NewCollection(k, "collection", "index").SearchSpecifications(nil, opts)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, 1, res.Total)
	assert.Equal(t, "Value found with search", res.Hits[0].Source.Validation.Fields["foo"].DefaultValue)
}

func TestScrollSpecificationsEmptyScrollId(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Collection.ScrollSpecifications: scroll id required"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ScrollSpecifications("", nil)
	assert.NotNil(t, err)
}

func TestScrollSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ScrollSpecifications("f00b4r", nil)
	assert.NotNil(t, err)
}

func TestScrollSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "scrollSpecifications", parsedQuery.Action)

			res := types.KuzzleSpecificationSearchResult{
				ScrollId: "f00b4r",
				Total:    1,
				Hits: []struct {
					Source types.KuzzleSpecificationsResult `json:"_source"`
				}{{Source: types.KuzzleSpecificationsResult{
					Index:      "index",
					Collection: "collection",
					Validation: types.KuzzleValidation{
						Strict: false,
						Fields: types.KuzzleValidationFields{
							"bar": {
								Mandatory:    true,
								Type:         "string",
								DefaultValue: "Value found with scroll",
							},
						},
					},
				}}},
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetScroll("1m")

	res, _ := collection.NewCollection(k, "collection", "index").ScrollSpecifications("f00b4r", opts)
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, 1, res.Total)
	assert.Equal(t, "Value found with scroll", res.Hits[0].Source.Validation.Fields["bar"].DefaultValue)
}

func TestValidateSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").ValidateSpecifications(types.KuzzleValidation{}, nil)
	assert.NotNil(t, err)
}

func TestValidateSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "validateSpecifications", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.ValidResponse{Valid: true}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	specifications := types.KuzzleValidation{
		Strict: false,
		Fields: types.KuzzleValidationFields{
			"foo": {
				Mandatory:    true,
				Type:         "string",
				DefaultValue: "Exciting value",
			},
		},
	}

	res, _ := collection.NewCollection(k, "collection", "index").ValidateSpecifications(specifications, nil)
	assert.Equal(t, true, res.Valid)
}

func TestUpdateSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").UpdateSpecifications(types.KuzzleValidation{}, nil)
	assert.NotNil(t, err)
}

func TestUpdateSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "updateSpecifications", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.KuzzleSpecifications{
				"index": {
					"collection": types.KuzzleValidation{
						Strict: true,
						Fields: types.KuzzleValidationFields{
							"foo": {
								Mandatory:    true,
								Type:         "string",
								DefaultValue: "Exciting value",
							},
						},
					},
				},
			}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	specifications := types.KuzzleValidation{
		Strict: true,
		Fields: types.KuzzleValidationFields{
			"foo": {
				Mandatory:    true,
				Type:         "string",
				DefaultValue: "Exciting value",
			},
		},
	}

	res, _ := collection.NewCollection(k, "collection", "index").UpdateSpecifications(specifications, nil)
	assert.Equal(t, true, res["index"]["collection"].Strict)
	assert.Equal(t, 1, len(res["index"]["collection"].Fields))
	assert.Equal(t, true, res["index"]["collection"].Fields["foo"].Mandatory)
	assert.Equal(t, "string", res["index"]["collection"].Fields["foo"].Type)
	assert.Equal(t, "Exciting value", res["index"]["collection"].Fields["foo"].DefaultValue)
}

func TestDeleteSpecificationsError(t *testing.T) {
	type Document struct {
		Title string
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").DeleteSpecifications(nil)
	assert.Equal(t, false, res.Acknowledged)
	assert.NotNil(t, err)
}

func TestDeleteSpecifications(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "collection", parsedQuery.Controller)
			assert.Equal(t, "deleteSpecifications", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := types.AckResponse{Acknowledged: true}
			r, _ := json.Marshal(res)
			return types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").DeleteSpecifications(nil)
	assert.Equal(t, true, res.Acknowledged)
}
