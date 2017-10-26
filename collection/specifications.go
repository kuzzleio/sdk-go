package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetSpecifications retrieves the current specifications of the collection.
func (dc *Collection) GetSpecifications(options types.QueryOptions) (*types.SpecificationsResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "getSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specifications := &types.SpecificationsResult{}
	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}

// SearchSpecifications searches specifications across indexes/collections according to the provided filters.
func (dc *Collection) SearchSpecifications(filters interface{}, options types.QueryOptions) (*types.SpecificationSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
		Body: struct {
			Query interface{} `json:"query"`
		}{Query: filters},
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()
		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specifications := &types.SpecificationSearchResult{}
	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}

// ScrollSpecifications retrieves next result of a specification search with scroll query.
func (dc *Collection) ScrollSpecifications(scrollId string, options types.QueryOptions) (*types.SpecificationSearchResult, error) {
	if scrollId == "" {
		return nil, types.NewError("Collection.ScrollSpecifications: scroll id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "scrollSpecifications",
		ScrollId:   scrollId,
	}

	if options != nil {
		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specifications := &types.SpecificationSearchResult{}
	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}

// ValidateSpecifications validates the provided specifications.
func (dc *Collection) ValidateSpecifications(specifications *types.Specification, options types.QueryOptions) (*types.ValidResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	type Specifications map[string]map[string]*types.Specification

	specificationsData := Specifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "validateSpecifications",
		Body:       specificationsData,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	response := &types.ValidResponse{}
	json.Unmarshal(res.Result, response)

	return response, nil
}

// UpdateSpecifications updates the current specifications of this collection.
func (dc *Collection) UpdateSpecifications(specifications *types.Specification, options types.QueryOptions) (*types.Specification, error) {
	ch := make(chan *types.KuzzleResponse)

	type Specifications map[string]map[string]*types.Specification

	specificationsData := &Specifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "updateSpecifications",
		Body:       specificationsData,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	specification := &Specifications{}
	json.Unmarshal(res.Result, specification)

	result := (*specification)[dc.index][dc.collection]

	return result, nil
}

// DeleteSpecifications deletes the current specifications of this collection.
func (dc *Collection) DeleteSpecifications(options types.QueryOptions) (*types.AckResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "deleteSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	response := &types.AckResponse{}
	json.Unmarshal(res.Result, response)

	return response, nil
}
