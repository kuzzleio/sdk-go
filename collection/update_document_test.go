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

func TestUpdateDocumentError(t *testing.T) {
  type Document struct {
    Name string
    Function string
  }

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      return types.KuzzleResponse{Error: types.MessageError{Message: "Unit test error"}}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  _, err := collection.NewCollection(k, "collection", "index").UpdateDocument("id", Document{Name: "Obi Wan", Function: "Legend"}, nil)
  assert.NotNil(t, err)
}

func TestUpdateDocument(t *testing.T) {
  id := "myId"

  type InitialContent struct {
    Name string
    Function string
  }
  initialContent := InitialContent{
    Name: "Anakin",
    Function: "Padawan",
  }

  type NewContent struct {
    Function string
  }
  updatePart := NewContent{"Jedi Knight"}

  c := &internal.MockedConnection{
    MockSend: func(query []byte, options *types.Options) types.KuzzleResponse {
      parsedQuery := &types.KuzzleRequest{}
      json.Unmarshal(query, parsedQuery)

      assert.Equal(t, "document", parsedQuery.Controller)
      assert.Equal(t, "update", parsedQuery.Action)
      assert.Equal(t, "index", parsedQuery.Index)
      assert.Equal(t, "collection", parsedQuery.Collection)
      assert.Equal(t, id, parsedQuery.Id)

      assert.Equal(t, "Jedi Knight", parsedQuery.Body.(map[string]interface{})["Function"])

      res := types.Document{Id: id, Source: []byte(`{"Name":"Anakin","Function":"Jedi Knight"}`)}
      r, _ := json.Marshal(res)
      return types.KuzzleResponse{Result: r}
    },
  }
  k, _ := kuzzle.NewKuzzle(c, nil)

  res, _ := collection.NewCollection(k, "collection", "index").UpdateDocument(id, updatePart, nil)

  assert.Equal(t, id, res.Id)

  var result InitialContent
  json.Unmarshal(res.Source, &result)

  assert.Equal(t, initialContent.Name, result.Name)
  assert.NotEqual(t, initialContent.Function, result.Name)
  assert.Equal(t, updatePart.Function, result.Function)
}
