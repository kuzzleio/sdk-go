package kuzzle

import (
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "errors"
)

type Collection struct {
  kuzzle            *Kuzzle
  index, collection string
  subscribeCallback interface{}
}

func NewCollection(kuzzle *Kuzzle, collection, index string) *Collection {
  return &Collection{
    index:      index,
    collection: collection,
    kuzzle:     kuzzle,
  }
}

/*
  Returns the number of documents matching the provided set of filters.

  There is a small delay between documents creation and their existence in our advanced search layer,
  usually a couple of seconds.
  That means that a document that was just been created won’t be returned by this function
*/
func (dc *Collection) Count(filters interface{}, options *types.Options) (*int, error) {
  type countResult struct {
    Count int `json:"count"`
  }

  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "document", "count", filters), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }
  result := &countResult{}
  json.Unmarshal(res.Result, result)

  return &result.Count, nil
}

/*
  Create a new empty data collection, with no associated mapping.
*/
func (dc *Collection) Create(options *types.Options) (*types.AckResponse, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "collection", "create", nil), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  ack := &types.AckResponse{}
  json.Unmarshal(res.Result, &ack)

  return ack, nil
}

/*
  Create a new document in Kuzzle.
  Takes an optional argument object with the following properties:
     - volatile (object, default: null):
         Additional information passed to notifications to other users
     - ifExist (string, allowed values: "error" (default), "replace"):
         If the same document already exists:
           - resolves with an error if set to "error".
           - replaces the existing document if set to "replace"
*/
func (dc *Collection) CreateDocument(id string, document interface{}, options *types.Options) (*types.Document, error) {
  ch := make(chan types.KuzzleResponse)

  go dc.kuzzle.Query(buildQuery(dc.collection, dc.index, "document", "create", document), options, ch)

  res := <-ch

  if res.Error.Message != "" {
    return nil, errors.New(res.Error.Message)
  }

  documentResponse := &types.Document{}
  json.Unmarshal(res.Result, documentResponse)

  return documentResponse, nil
}
