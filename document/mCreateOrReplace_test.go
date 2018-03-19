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

func TestMCreateOrReplaceIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.MCreateOrReplace("", "collection", "body", nil)
	assert.NotNil(t, err)
}

func TestMCreateOrReplaceCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.MCreateOrReplace("index", "", "body", nil)
	assert.NotNil(t, err)
}

func TestMCreateOrReplaceBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)

	_, err := d.MCreateOrReplace("index", "collection", "", nil)
	assert.NotNil(t, err)
}

func TestMCreateOrReplaceDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.MCreateOrReplace("index", "collection", "body", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestMCreateOrReplaceDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mCreateOrReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`{
				"hits": [
					{
						"_id": "id1",
						"_index": "index",
						"_shards": {
							"failed": 0,
							"successful": 1,
							"total": 2
						},
						"_source": {
							"document": "body"
						},
						"_meta": {
							"active": true,
							"author": "-1",
							"createdAt": 1484225532686,
							"deletedAt": null,
							"updatedAt": null,
							"updater": null
						},
						"_type": "collection",
						"_version": 1,
						"created": true,
						"result": "created"
					}
				],
				"total": "1"
			}`),
			}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)

	_, err := d.MCreateOrReplace("index", "collection", "body", nil)
	assert.Nil(t, err)
}