package document_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/document"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("", "collection", "id1", "body", opts)
	assert.NotNil(t, err)
}

func TestCreateCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("index", "", "id1", "body", opts)
	assert.NotNil(t, err)
}

func TestCreateIdNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("index", "collection", "", "body", opts)
	assert.NotNil(t, err)
}

func TestCreateBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("index", "collection", "id1", "", opts)
	assert.NotNil(t, err)
}

func TestCreateDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("index", "collection", "id1", "body", opts)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestCreateDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "create", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`
				{
					"_id": "myId",
					"_version": 1,
				}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.Create("index", "collection", id, "body", opts)
	assert.Nil(t, err)
}
