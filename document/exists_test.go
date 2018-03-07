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

func TestExistsIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Exists("", "collection", "id1")
	assert.NotNil(t, err)
}

func TestExistsCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Exists("index", "", "id1")
	assert.NotNil(t, err)
}

func TestExistsIdNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Exists("index", "collection", "")
	assert.NotNil(t, err)
}

func TestExistsDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	_, err := d.Exists("index", "collection", "id1")
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestExistsDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "exists", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			return &types.KuzzleResponse{Result: []byte(`true`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	res, err := d.Exists("index", "collection", id)
	assert.Nil(t, err)
	assert.Equal(t, true, res)
}
