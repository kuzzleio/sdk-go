package collection_test

import (
  "testing"
  "github.com/kuzzleio/sdk-go/internal"
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/collection"
)

func TestCreateDocumentError(t *testing.T) {
  type Document struct {
    Title string
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := collection.NewCollection(k, "collection", "index").CreateDocument("id", Document{Title: "yolo"}, nil)
  assert.NotNil(t, err)
}

func TestCreateDocument(t *testing.T) {
  type Document struct {
    Title string
  }

  id := "myId"

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "document", parsedQuery.Controller)
      assert.Equal(t, "create", parsedQuery.Action)
      assert.Equal(t, "index", parsedQuery.Index)
      assert.Equal(t, "collection", parsedQuery.Collection)
      assert.Equal(t, id, parsedQuery.Id)

      assert.Equal(t, "yolo", parsedQuery.Body.(map[string]interface{})["Title"])

      res := types.Document{Id: id, Source: []byte(`{"title": "yolo"}`)}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := collection.NewCollection(k, "collection", "index").CreateDocument(id, Document{Title: "yolo"}, nil)
  assert.Equal(t, id, res.Id)
}
