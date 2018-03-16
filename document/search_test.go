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

func TestSearchIndexNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	from := 2
	size := 4
	opts := types.SearchOptions{Type: "all", From: &from, Size: &size, Scroll: "1m"}
	_, err := d.Search("", "collection", "body", &opts)
	assert.NotNil(t, err)
}

func TestSearchCollectionNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	from := 2
	size := 4
	opts := types.SearchOptions{Type: "all", From: &from, Size: &size, Scroll: "1m"}
	_, err := d.Search("index", "", "body", &opts)
	assert.NotNil(t, err)
}

func TestSearchBodyNull(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	d := document.NewDocument(k)
	from := 2
	size := 4
	opts := types.SearchOptions{Type: "all", From: &from, Size: &size, Scroll: "1m"}
	_, err := d.Search("index", "collection", "", &opts)
	assert.NotNil(t, err)
}

func TestSearchDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	d := document.NewDocument(k)
	from := 2
	size := 4
	opts := types.SearchOptions{Type: "all", From: &from, Size: &size, Scroll: "1m"}
	_, err := d.Search("index", "collection", "body", &opts)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestSearchDocument(t *testing.T) {

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)
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
	from := 2
	size := 4
	opts := types.SearchOptions{Type: "all", From: &from, Size: &size, Scroll: "1m"}
	_, err := d.Search("index", "collection", "body", &opts)
	assert.Nil(t, err)
}
