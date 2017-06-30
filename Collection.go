package core

import ("github.com/kuzzleio/sdk-core/types"
  "encoding/json"
)

type Collection struct {
  kuzzle *Kuzzle
  index, Collection string
  subscribeCallback interface{}
}

func NewCollection(kuzzle *Kuzzle, collection, index string) *Collection {
  return &Collection{
    index: index,
    Collection: collection,
    kuzzle: kuzzle,
  }
}

func (dc *Collection) Count(filters interface{}, resultChan chan<- int) {
  type countResult struct {
    Count int `json:"count"`
  }

  type body struct {
    Filters []byte
  }

  result := make(chan types.KuzzleResponse)

  dc.kuzzle.Query(dc.makeQuery("document", "count", filters), result, nil)

  res := <-result

  count := countResult{}
  json.Unmarshal(res.Result, &count)
  resultChan <- count.Count
}

func (dc *Collection) Subscribe(filters interface{}, subChan chan<- types.KuzzleNotification, result chan<- types.KuzzleResponse) {
  MyDocument := types.KuzzleRequest{
    Controller: "realtime",
    Action: "subscribe",
    Index: dc.index,
    Collection: dc.Collection,
    Body: filters,
  }

  dc.kuzzle.Query(MyDocument, result, subChan)
}

func (dc *Collection) makeQuery(controller string, action string, body interface{}) types.KuzzleRequest {
  return types.KuzzleRequest{
    Controller: controller,
    Action: action,
    Index: dc.index,
    Collection: dc.Collection,
    Body: body,
  }
}