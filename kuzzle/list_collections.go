package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

// ListCollections retrieves the list of known data collections contained in a specified index.
func (k Kuzzle) ListCollections(index string, options types.QueryOptions) ([]types.CollectionsList, error) {
	if index == "" {
		if k.defaultIndex == "" {
			return nil, errors.New("Kuzzle.ListCollections: index required")
		}
		index = k.defaultIndex
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "collection",
		Action:     "list",
		Index:      index,
	}

	if options != nil {
		if options.GetType() != "" {
			type body struct {
				Type string `json:"type"`
			}

			query.Body = &body{Type: options.GetType()}
		}
	}

	type collections struct {
		Collections []types.CollectionsList `json:"collections"`
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	s := collections{}
	json.Unmarshal(res.Result, &s)

	return s.Collections, nil
}
