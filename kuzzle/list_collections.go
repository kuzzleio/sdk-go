package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
 * List data collections
 */
func (k Kuzzle) ListCollections(index string, options *types.Options) ([]types.CollectionsList, error) {
	if index == "" {
		return nil, errors.New("Kuzzle.ListCollections: index required")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "collection",
		Action:     "list",
		Index:      index,
	}

	type collections struct {
		Collections []types.CollectionsList `json:"collections"`
	}

	go k.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	s := collections{}
	json.Unmarshal(res.Result, &s)

	return s.Collections, nil
}
