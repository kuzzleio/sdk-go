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

func TestReplaceIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("", "collection", "id1", "body", nil)
	assert.NotNil(t, err)
}

func TestReplaceCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "", "id1", "body", nil)
	assert.NotNil(t, err)
}

func TestReplaceIdNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "", "body", nil)
	assert.NotNil(t, err)
}

func TestReplaceBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "id1", "", nil)
	assert.NotNil(t, err)
}

func TestReplaceDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", "id1", "body", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestReplaceDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "replace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`
				{
					"_id": "<documentId>",
					"_version": "<number>",// The new version number of this document
					"created": false
				}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.Replace("index", "collection", id, "body", nil)
	assert.Nil(t, err)
}
