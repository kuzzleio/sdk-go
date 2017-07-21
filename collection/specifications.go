package collection

import (
  "errors"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
)

/*
  Retrieves the current specifications of the collection.
*/
func (dc Collection) GetSpecifications(options *types.Options) (types.KuzzleSpecificationsResult, error) {
  ch := make(chan types.KuzzleResponse)

  query := types.KuzzleRequest{
    Collection: dc.collection,
    Index: dc.index,
    Controller: "collection",
    Action: "getSpecifications",
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
    Index: dc.index,
    Controller: "collection",
    Action: "validateSpecifications",
    Body: specificationsData,
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
    Index: dc.index,
    Controller: "collection",
    Action: "updateSpecifications",
    Body: specificationsData,
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
    Index: dc.index,
    Controller: "collection",
    Action: "deleteSpecifications",
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
