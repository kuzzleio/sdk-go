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

func TestMUpdateIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.MUpdate("", "collection", "body", opts)
	assert.NotNil(t, err)
}

func TestMUpdateCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.MUpdate("index", "", "body", opts)
	assert.NotNil(t, err)
}

func TestMUpdateBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.MUpdate("index", "collection", "", opts)
	assert.NotNil(t, err)
}

func TestMUpdateDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.MUpdate("index", "collection", "body", opts)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestMUpdateDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mUpdate", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			return &types.KuzzleResponse{Result: []byte(`
			{
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
	opts := &document.DocumentOptions{WaitFor: true, Volatile: ""}
	_, err := d.MUpdate("index", "collection", "body", opts)
	assert.Nil(t, err)
}
