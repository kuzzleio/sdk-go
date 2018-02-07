package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
)

//export kuzzle_collection_count
func kuzzle_collection_count(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.int_result {
	res, err := (*collection.Collection)(c.instance).Count(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCIntResult(res, err)
}

//export kuzzle_collection_create_document
func kuzzle_collection_create_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := (*collection.Collection)(c.instance).CreateDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_collection_delete_document
func kuzzle_collection_delete_document(c *C.collection, id *C.char, options *C.query_options) *C.string_result {
	res, err := (*collection.Collection)(c.instance).DeleteDocument(C.GoString(id), SetQueryOptions(options))
	return goToCStringResult(&res, err)
}

//export kuzzle_collection_document_exists
func kuzzle_collection_document_exists(c *C.collection, id *C.char, options *C.query_options) *C.bool_result {
	res, err := (*collection.Collection)(c.instance).DocumentExists(C.GoString(id), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_fetch_document
func kuzzle_collection_fetch_document(c *C.collection, id *C.char, options *C.query_options) *C.document_result {
	res, err := (*collection.Collection)(c.instance).FetchDocument(C.GoString(id), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_collection_replace_document
func kuzzle_collection_replace_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := (*collection.Collection)(c.instance).ReplaceDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_collection_update_document
func kuzzle_collection_update_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := (*collection.Collection)(c.instance).UpdateDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_collection_scroll
func kuzzle_collection_scroll(c *C.collection, scrollId *C.char, options *C.query_options) *C.search_result {
	res, err := (*collection.Collection)(c.instance).Scroll(C.GoString(scrollId), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_collection_search
func kuzzle_collection_search(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.search_result {
	res, err := (*collection.Collection)(c.instance).Search(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_collection_m_create_document
func kuzzle_collection_m_create_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.document_array_result {
	res, err := (*collection.Collection)(c.instance).MCreateDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCDocumentArrayResult(c, res, err)
}

//export kuzzle_collection_m_create_or_replace_document
func kuzzle_collection_m_create_or_replace_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.document_array_result {
	res, err := (*collection.Collection)(c.instance).MCreateOrReplaceDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCDocumentArrayResult(c, res, err)
}

//export kuzzle_collection_m_delete_document
func kuzzle_collection_m_delete_document(c *C.collection, ids **C.char, idsCount C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*collection.Collection)(c.instance).MDeleteDocument(cToGoStrings(ids, idsCount), SetQueryOptions(options))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_collection_m_get_document
func kuzzle_collection_m_get_document(c *C.collection, ids **C.char, idsCount C.size_t, options *C.query_options) *C.document_array_result {
	res, err := (*collection.Collection)(c.instance).MGetDocument(cToGoStrings(ids, idsCount), SetQueryOptions(options))
	return goToCDocumentArrayResult(c, res, err)
}

//export kuzzle_collection_m_replace_document
func kuzzle_collection_m_replace_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.document_array_result {
	res, err := (*collection.Collection)(c.instance).MReplaceDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCDocumentArrayResult(c, res, err)
}

//export kuzzle_collection_m_update_document
func kuzzle_collection_m_update_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.document_array_result {
	res, err := (*collection.Collection)(c.instance).MUpdateDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCDocumentArrayResult(c, res, err)
}
