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

func TestMCreateDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MCreateDocument(documents, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestMCreateDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mCreate", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := map[string]interface{}{"Total": 2, "Hits": documents}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MCreateDocument(documents, nil)
	assert.Equal(t, 2, len(res))

	for index, doc := range res {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MCreateDocument() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MCreateDocument([]*collection.Document{
		{Content: []byte(`{"title":"yolo"}`)},
		{Content: []byte(`{"title":"oloy"}`)},
	}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Content)
}

func TestMCreateOrReplaceDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MCreateOrReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMCreateOrReplaceDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mCreateOrReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := map[string]interface{}{"Total": 2, "Hits": documents}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MCreateOrReplaceDocument(documents, nil)
	assert.Equal(t, 2, len(res))

	for index, doc := range res {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MCreateOrReplaceDocument() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MCreateOrReplaceDocument([]*collection.Document{
		{Content: []byte(`{"title":"yolo"}`)},
		{Content: []byte(`{"title":"oloy"}`)},
	}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Content)
}

func TestMReplaceDocumentEmptyDocuments(t *testing.T) {
	documents := []*collection.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.MReplaceDocument: please provide at least one document to replace"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMReplaceDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mReplace", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := map[string]interface{}{"Total": 2, "Hits": documents}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)
	assert.Equal(t, 2, len(res))

	for index, doc := range res {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MReplaceDocument() {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MReplaceDocument(documents, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Content)
}

func TestMDeleteDocumentEmptyIds(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("should have failed before that line")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MDeleteDocument([]string{}, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Collection.MDeleteDocument: please provide at least one id of document to delete", err.(*types.KuzzleError).Message)
}

func TestMDeleteDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MDeleteDocument([]string{"foo", "bar"}, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestMDeleteDocument(t *testing.T) {
	ids := []string{"foo", "bar"}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mDelete", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, []interface{}{"foo", "bar"}, parsedQuery.Body.(map[string]interface{})["ids"])

			return &types.KuzzleResponse{Result: []byte(`["foo","bar"]`)}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MDeleteDocument(ids, nil)
	assert.Equal(t, []string{"foo", "bar"}, res)
}

func ExampleCollection_MDeleteDocument() {
	ids := []string{"foo", "bar"}
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MDeleteDocument(ids, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}

func TestMGetDocumentEmptyIds(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.MGetDocument: please provide at least one id of document to retrieve"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MGetDocument([]string{"foo", "bar"}, nil)
	assert.NotNil(t, err)
}

func TestMGetDocument(t *testing.T) {
	hits := []*collection.Document{
		{Id: "foo", Content: json.RawMessage(`{"title":"foo"}`)},
		{Id: "bar", Content: json.RawMessage(`{"title":"bar"}`)},
	}

	ids := []string{"foo", "bar"}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mGet", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, []interface{}{"foo", "bar"}, parsedQuery.Body.(map[string]interface{})["ids"])

			res := map[string]interface{}{"Total": 2, "Hits": hits}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MGetDocument(ids, nil)
	assert.Equal(t, 2, len(res))

	for i := range res {
		assert.Equal(t, hits[i].Id, res[i].Id)
		assert.Equal(t, hits[i].Content, res[i].Content)
	}
}

func ExampleCollection_MGetDocument() {
	ids := []string{"foo", "bar"}
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MGetDocument(ids, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Content)
}

func TestMUpdateDocumentEmptyDocuments(t *testing.T) {
	documents := []*collection.Document{}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Collection.MUpdateDocument: please provide at least one document to update"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocumentError(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.NotNil(t, err)
}

func TestMUpdateDocument(t *testing.T) {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "mUpdate", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)

			res := map[string]interface{}{"Total": 2, "Hits": documents}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)
	assert.Equal(t, 2, len(res))

	for index, doc := range res {
		assert.Equal(t, documents[index].Id, doc.Id)
		assert.Equal(t, documents[index].Content, doc.Content)
	}
}

func ExampleCollection_MUpdateDocument() {
	documents := []*collection.Document{
		{Id: "foo", Content: []byte(`{"title":"Foo"}`)},
		{Id: "bar", Content: []byte(`{"title":"Bar"}`)},
	}

	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, err := collection.NewCollection(k, "collection", "index").MUpdateDocument(documents, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res[0].Id, res[0].Content)
}
