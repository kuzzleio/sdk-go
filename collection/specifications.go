package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Retrieves the current specifications of the collection.
*/
func (dc Collection) GetSpecifications(options *types.Options) (types.KuzzleSpecificationsResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "getSpecifications",
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSpecificationsResult{}, errors.New(res.Error.Message)
	}

	specification := types.KuzzleSpecificationsResult{}
	json.Unmarshal(res.Result, &specification)

	return specification, nil
}

/*
  Searches specifications across indexes/collections according to the provided filters.
*/
func (dc Collection) SearchSpecifications(filters interface{}, options *types.Options) (types.KuzzleSpecificationSearchResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
		Body: struct {
			Query interface{} `json:"query"`
		}{Query: filters},
	}

	if options != nil {
		query.From = options.From
		query.Size = options.Size
		if options.Scroll != "" {
			query.Scroll = options.Scroll
		}
	}

	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSpecificationSearchResult{}, errors.New(res.Error.Message)
	}

	specifications := types.KuzzleSpecificationSearchResult{}
	json.Unmarshal(res.Result, &specifications)

	return specifications, nil
}

/*
  Retrieves next result of a specification search with scroll query.
*/
func (dc Collection) ScrollSpecifications(scrollId string, options *types.Options) (types.KuzzleSpecificationSearchResult, error) {
	if scrollId == "" {
		return types.KuzzleSpecificationSearchResult{}, errors.New("Collection.ScrollSpecifications: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "collection",
		Action:     "scrollSpecifications",
		ScrollId:   scrollId,
	}

	if options != nil {
		if options.Scroll != "" {
			query.Scroll = options.Scroll
		}
	}

	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSpecificationSearchResult{}, errors.New(res.Error.Message)
	}

	specifications := types.KuzzleSpecificationSearchResult{}
	json.Unmarshal(res.Result, &specifications)

	return specifications, nil
}

/*
  Validates the provided specifications.
*/
func (dc Collection) ValidateSpecifications(specifications types.KuzzleValidation, options *types.Options) (types.ValidResponse, error) {
	ch := make(chan types.KuzzleResponse)

	specificationsData := types.KuzzleSpecifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "validateSpecifications",
		Body:       specificationsData,
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.ValidResponse{}, errors.New(res.Error.Message)
	}

	response := types.ValidResponse{}
	json.Unmarshal(res.Result, &response)

	return response, nil
}

/*
  Updates the current specifications of this collection.
*/
func (dc Collection) UpdateSpecifications(specifications types.KuzzleValidation, options *types.Options) (types.KuzzleSpecifications, error) {
	ch := make(chan types.KuzzleResponse)

	specificationsData := types.KuzzleSpecifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "updateSpecifications",
		Body:       specificationsData,
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSpecifications{}, errors.New(res.Error.Message)
	}

	specification := types.KuzzleSpecifications{}
	json.Unmarshal(res.Result, &specification)

	return specification, nil
}

/*
  Deletes the current specifications of this collection.
*/
func (dc Collection) DeleteSpecifications(options *types.Options) (types.AckResponse, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "deleteSpecifications",
	}
	go dc.kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.AckResponse{Acknowledged: false}, errors.New(res.Error.Message)
	}

	response := types.AckResponse{}
	json.Unmarshal(res.Result, &response)

	return response, nil
}
