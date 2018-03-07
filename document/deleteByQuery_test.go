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

func TestDeleteByQueryIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.DeleteByQuery("", "collection", "body", opts)
	assert.NotNil(t, err)
}

func TestDeleteByQueryCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.DeleteByQuery("index", "", "body", opts)
	assert.NotNil(t, err)
}

func TestDeleteByQueryBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.DeleteByQuery("index", "collection", "", opts)
	assert.NotNil(t, err)
}

func TestDeleteByQueryDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.DeleteByQuery("index", "collection", "body", opts)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestDeleteByQueryDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "deleteByQuery", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`
			{
				"hits": ["id1", "id2"]
			}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.DeleteByQuery("index", "collection", "body", opts)
	assert.Nil(t, err)
}
