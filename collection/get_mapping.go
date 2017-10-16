package collection

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// GetMapping retrieves the current mapping of the collection.
func (dc *Collection) GetMapping(options types.QueryOptions) (*Mapping, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "getMapping",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return &Mapping{}, errors.New(res.Error.Message)
	}

	type mappingResult map[string]struct {
		Mappings map[string]struct {
			Properties *types.KuzzleFieldsMapping `json:"properties"`
		} `json:"mappings"`
	}

	result := mappingResult{}
	json.Unmarshal(res.Result, &result)

	if _, ok := result[dc.index]; ok {
		indexMappings := result[dc.index].Mappings

		if _, ok := indexMappings[dc.collection]; ok {
			cm := NewMapping(dc)
			cm.Set(indexMappings[dc.collection].Properties)

			return cm, nil
		} else {
			return &Mapping{}, errors.New("No mapping found for collection " + dc.collection)
		}
	} else {
		return &Mapping{}, errors.New("No mapping found for index " + dc.index)
	}
}
