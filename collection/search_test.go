package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").Search(&types.SearchFilters{}, nil)
	assert.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	hits := []*collection.Document{
		&collection.Document{Id: "doc42", Content: json.RawMessage(`{"foo":"bar"}`)},
	}

	results := collection.SearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := collection.SearchResult{Total: results.Total, Hits: results.Hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").Search(&types.SearchFilters{}, nil)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, len(hits), len(res.Hits))
	assert.Equal(t, res.Hits[0].Id, hits[0].Id)
	assert.Equal(t, res.Hits[0].Content, hits[0].Content)
}

func ExampleCollection_Search() {
	hits := make([]*collection.Document, 1)
	hits[0] = &collection.Document{Id: "doc42", Content: json.RawMessage(`{"foo":"bar"}`)}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").Search(&types.SearchFilters{}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestSearchWithScroll(t *testing.T) {
	hits := []*collection.Document{
		&collection.Document{Id: "doc42", Content: json.RawMessage(`{"foo":"bar"}`)},
	}
	results := collection.SearchResult{Total: 42, Hits: hits}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, 2, parsedQuery.From)
			assert.Equal(t, 4, parsedQuery.Size)
			assert.Equal(t, "1m", parsedQuery.Scroll)

			res := collection.SearchResult{Total: results.Total, Hits: results.Hits, ScrollId: "f00b4r"}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	opts := types.NewQueryOptions()
	opts.SetFrom(2)
	opts.SetSize(4)
	opts.SetScroll("1m")
	res, _ := collection.NewCollection(k, "collection", "index").Search(&types.SearchFilters{}, opts)
	assert.Equal(t, results.Total, res.Total)
	assert.Equal(t, len(hits), len(res.Hits))
	assert.Equal(t, "f00b4r", res.ScrollId)
	assert.Equal(t, res.Hits[0].Id, hits[0].Id)
	assert.Equal(t, res.Hits[0].Content, hits[0].Content)
}
