package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type ICollectionMapping interface {
	Apply()
	Refresh()
	Set()
	SetHeaders()
}

type CollectionMapping struct {
	Mapping    *types.KuzzleFieldMapping
	Collection *Collection
}

// Apply applies the new mapping to the data collection.
func (cm *CollectionMapping) Apply(options types.QueryOptions) (*CollectionMapping, error) {
	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Properties *types.KuzzleFieldMapping `json:"properties"`
	}

	query := &types.KuzzleRequest{
		Collection: cm.Collection.collection,
		Index:      cm.Collection.index,
		Controller: "collection",
		Action:     "updateMapping",
		Body:       &body{Properties: cm.Mapping},
	}

	go cm.Collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return cm, errors.New(res.Error.Message)
	}

	return cm.Refresh(nil)
}

// Refresh Replaces the current content with the mapping stored in Kuzzle.
// Calling this function will discard any uncommitted changes. You can commit changes by calling the “apply” function
func (cm *CollectionMapping) Refresh(options types.QueryOptions) (*CollectionMapping, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: cm.Collection.collection,
		Index:      cm.Collection.index,
		Controller: "collection",
		Action:     "getMapping",
	}
	go cm.Collection.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return cm, errors.New(res.Error.Message)
	}

	type mappingResult map[string]struct {
		Mappings map[string]struct {
			Properties *types.KuzzleFieldMapping `json:"properties"`
		} `json:"mappings"`
	}

	result := mappingResult{}
	json.Unmarshal(res.Result, &result)

	if _, ok := result[cm.Collection.index]; ok {
		indexMappings := result[cm.Collection.index].Mappings

		if _, ok := indexMappings[cm.Collection.collection]; ok {
			cm.Mapping = indexMappings[cm.Collection.collection].Properties

			return cm, nil
		} else {
			return cm, errors.New("No mapping found for collection " + cm.Collection.collection)
		}
	} else {
		return cm, errors.New("No mapping found for index " + cm.Collection.index)
	}
}

/*
  Adds or updates a field mapping.

  Changes made by this function won’t be applied until you call the apply method
*/
func (cm *CollectionMapping) Set(mappings *types.KuzzleFieldMapping) *CollectionMapping {
	for field, mapping := range *mappings {
		(*cm.Mapping)[field] = mapping
	}

	return cm
}

// SetHeaders is is a helper function returning itself, allowing to easily chain calls.
// If the replace argument is set to true, replace the current headers with the provided content.
// Otherwise, it appends the content to the current headers, only replacing already existing values
func (cm *CollectionMapping) SetHeaders(content map[string]interface{}, replace bool) *CollectionMapping {
	cm.Collection.SetHeaders(content, replace)

	return cm
}
