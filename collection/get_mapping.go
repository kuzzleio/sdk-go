package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetMapping retrieves the current mapping of the collection.
func (dc *Collection) GetMapping(options types.QueryOptions) (*CollectionMapping, error) {
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
		return nil, res.Error
	}

	type mappingResult map[string]struct {
		Mappings map[string]struct {
			Properties *types.KuzzleFieldMapping `json:"properties"`
		} `json:"mappings"`
	}

	result := mappingResult{}
	json.Unmarshal(res.Result, &result)

	if _, ok := result[dc.index]; ok {
		indexMappings := result[dc.index].Mappings

		if _, ok := indexMappings[dc.collection]; ok {
			return &CollectionMapping{Mapping: indexMappings[dc.collection].Properties, Collection: dc}, nil
		} else {
			return nil, types.NewError("No mapping found for collection " + dc.collection, 404)
		}
	} else {
		return nil, types.NewError("No mapping found for index " + dc.index, 404)
	}
}
