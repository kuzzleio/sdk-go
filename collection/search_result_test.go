package collection_test

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchResultSetup(c *internal.MockedConnection) *collection.SearchResult {
	var _c *internal.MockedConnection

	if c == nil {
		_c = &internal.MockedConnection{}
	} else {
		_c = c
	}

	kuzzle, _ := kuzzle.NewKuzzle(_c, nil)
	cl := collection.NewCollection(kuzzle, "collection", "index")

	filters := &types.SearchFilters{}
	options := types.NewQueryOptions()
	options.SetSize(0)
	options.SetFrom(0)
	response := &types.KuzzleResponse{}

	response.Result, _ = json.Marshal(map[string]interface{}{
		"total": 4,
		"hits": []map[string]interface{}{
			{
				"_id": "Bateman",
				"_source": map[string]string{
					"firstname": "Patrick",
					"killcount": "20",
				},
			},
			{
				"_id": "Morgan",
				"_source": map[string]string{
					"firstname": "Dexter",
					"killcount": "135?",
				},
			},
		},
		"aggregations": map[string]interface{}{"foo": "bar"},
	})

	return collection.NewSearchResult(cl, filters, options, response)
}

func TestConstructor(t *testing.T) {
	kuzzle, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)
	cl := collection.NewCollection(kuzzle, "collection", "index")

	filters := &types.SearchFilters{}
	options := types.NewQueryOptions()
	options.SetSize(42)
	options.SetFrom(13)
	response := &types.KuzzleResponse{}

	response.Result, _ = json.Marshal(map[string]interface{}{
		"total":      4,
		"_scroll_id": "fooscrollid",
		"hits": []map[string]string{
			{"_id": "foo", "_source": "bar"},
			{"_id": "Bateman", "_source": "Patrick"},
		},
		"aggregations": map[string]interface{}{"foo": "bar"},
	})

	searchResult := collection.NewSearchResult(cl, filters, options, response)

	assert.Equal(t, filters, searchResult.Filters)
	assert.Equal(t, cl, searchResult.Collection)
	assert.Equal(t, "fooscrollid", searchResult.Options.ScrollId())
	assert.Equal(t, 13, searchResult.Options.From())
	assert.Equal(t, 42, searchResult.Options.Size())
	assert.Equal(t, 4, searchResult.Total)
	assert.Equal(t, 2, searchResult.Fetched)
	assert.Equal(t, "foo", searchResult.Documents[0].Id)
	assert.Equal(t, "Bateman", searchResult.Documents[1].Id)
	assert.Equal(t, "bar", searchResult.Aggregations["foo"])
}

func TestFetchNextNoResultLeft(t *testing.T) {
	searchResult := searchResultSetup(nil)

	searchResult.Fetched = 4

	res, err := searchResult.FetchNext()

	assert.Nil(t, res)
	assert.Nil(t, err)
}

func TestFetchNextBadOptions(t *testing.T) {
	searchResult := searchResultSetup(nil)

	searchResult.Options.SetSize(0)
	searchResult.Options.SetScrollId("")

	res, err := searchResult.FetchNext()

	assert.Nil(t, res)
	assert.Equal(t, &types.KuzzleError{Message: "SearchResult.FetchNext: Unable to retrieve results: missing scrollId or from/size parameters", Stack: "", Status: 400}, err)
}

func TestFetchNextWithScroll(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "scroll", parsedQuery.Action)
			assert.Equal(t, "foobar", parsedQuery.ScrollId)

			fakeResult, _ := json.Marshal(map[string]interface{}{
				"total":      4,
				"_scroll_id": "nextscrollid",
				"hits": []map[string]string{
					{"_id": "foo2", "_source": "bar"},
					{"_id": "Bateman2", "_source": "Patrick"},
				},
			})

			return &types.KuzzleResponse{Result: fakeResult}
		},
	}

	searchResult := searchResultSetup(c)

	assert.Equal(t, 2, searchResult.Fetched)

	searchResult.Options.SetSize(0)
	searchResult.Options.SetScrollId("foobar")
	next, err := searchResult.FetchNext()

	assert.Nil(t, err)
	assert.Equal(t, "nextscrollid", next.Options.ScrollId())
	assert.Equal(t, 4, next.Fetched)
	assert.Equal(t, 4, next.Total)
	assert.Equal(t, "foo2", next.Documents[0].Id)
	assert.Equal(t, "Bateman2", next.Documents[1].Id)
	assert.Nil(t, next.Aggregations)
}

func TestFetchNextWithPage(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)
			assert.Equal(t, 2, parsedQuery.From)
			assert.Equal(t, 2, parsedQuery.Size)

			fakeResult, _ := json.Marshal(map[string]interface{}{
				"total": 4,
				"hits": []map[string]string{
					{"_id": "foo2", "_source": "bar"},
					{"_id": "Bateman2", "_source": "Patrick"},
				},
			})

			return &types.KuzzleResponse{Result: fakeResult}
		},
	}

	searchResult := searchResultSetup(c)

	assert.Equal(t, 2, searchResult.Fetched)

	searchResult.Options.SetSize(2)
	next, err := searchResult.FetchNext()

	assert.Nil(t, err)
	assert.Equal(t, 4, next.Fetched)
	assert.Equal(t, 4, next.Total)
	assert.Equal(t, "foo2", next.Documents[0].Id)
	assert.Equal(t, "Bateman2", next.Documents[1].Id)
	assert.Nil(t, next.Aggregations)
}

func TestFetchNextDoNotFetchAfterLastPage(t *testing.T) {
	searchResult := searchResultSetup(nil)

	assert.Equal(t, 2, searchResult.Fetched)

	searchResult.Options.SetFrom(2)
	searchResult.Options.SetSize(2)
	next, err := searchResult.FetchNext()

	assert.Nil(t, err)
	assert.Nil(t, next)
}

func TestFetchNextWithSearchAfter(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "search", parsedQuery.Action)

			filters := &types.SearchFilters{}
			rawBody, _ := json.Marshal(parsedQuery.Body)
			json.Unmarshal(rawBody, filters)

			assert.Equal(t, "Dexter", filters.SearchAfter[0])
			assert.Equal(t, "135?", filters.SearchAfter[1])

			fakeResult, _ := json.Marshal(map[string]interface{}{
				"total": 4,
				"hits": []map[string]string{
					{"_id": "foo2", "_source": "bar"},
					{"_id": "Bateman2", "_source": "Patrick"},
				},
			})

			return &types.KuzzleResponse{Result: fakeResult}
		},
	}

	searchResult := searchResultSetup(c)

	assert.Equal(t, 2, searchResult.Fetched)

	searchResult.Filters.Sort = append(searchResult.Filters.Sort, "firstname")
	searchResult.Filters.Sort = append(searchResult.Filters.Sort, map[string]interface{}{
		"killcount": map[string]string{"order": "asc"},
	})

	searchResult.Options.SetSize(2)
	next, err := searchResult.FetchNext()

	assert.Nil(t, err)
	assert.Equal(t, 4, next.Fetched)
	assert.Equal(t, 4, next.Total)
	assert.Equal(t, "foo2", next.Documents[0].Id)
	assert.Equal(t, "Bateman2", next.Documents[1].Id)
	assert.Nil(t, next.Aggregations)
}
