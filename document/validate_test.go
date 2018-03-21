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

func TestValidateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Validate("", "collection", "body", nil)
	assert.NotNil(t, err)
}

func TestValidateCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Validate("index", "", "body", nil)
	assert.NotNil(t, err)
}

func TestValidateBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	_, err := d.Validate("index", "collection", "", nil)
	assert.NotNil(t, err)
}

func TestValidateDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	_, err := d.Validate("index", "collection", "body", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestValidateDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "validate", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`true`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	_, err := d.Validate("index", "collection", "body", nil)
	assert.Nil(t, err)
}
