package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
*/
import "C"

//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_collection_count(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.int_result {
	res, err := cToGoCollection(c).Count(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCIntResult(res, err)
}

//export kuzzle_wrapper_collection_create_document
func kuzzle_wrapper_collection_create_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := cToGoCollection(c).CreateDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_wrapper_collection_delete_document
func kuzzle_wrapper_collection_delete_document(c *C.collection, id *C.char, options *C.query_options) *C.string_result {
	res, err := cToGoCollection(c).DeleteDocument(C.GoString(id), SetQueryOptions(options))
	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_collection_document_exists
func kuzzle_wrapper_collection_document_exists(c *C.collection, id *C.char, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).DocumentExists(C.GoString(id), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_collection_fetch_document
func kuzzle_wrapper_collection_fetch_document(c *C.collection, id *C.char, options *C.query_options) *C.document_result {
	res, err := cToGoCollection(c).FetchDocument(C.GoString(id), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_wrapper_collection_replace_document
func kuzzle_wrapper_collection_replace_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := cToGoCollection(c).ReplaceDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_wrapper_collection_update_document
func kuzzle_wrapper_collection_update_document(c *C.collection, id *C.char, document *C.document, options *C.query_options) *C.document_result {
	res, err := cToGoCollection(c).UpdateDocument(C.GoString(id), cToGoDocument(c, document), SetQueryOptions(options))
	return goToCDocumentResult(c, res, err)
}

//export kuzzle_wrapper_collection_scroll
func kuzzle_wrapper_collection_scroll(c *C.collection, scrollId *C.char, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).Scroll(C.GoString(scrollId), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_search
func kuzzle_wrapper_collection_search(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).Search(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_create_document
func kuzzle_wrapper_collection_m_create_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).MCreateDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_create_or_replace_document
func kuzzle_wrapper_collection_m_create_or_replace_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).MCreateOrReplaceDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_delete_document
func kuzzle_wrapper_collection_m_delete_document(c *C.collection, ids **C.char, idsCount C.uint, options *C.query_options) *C.string_array_result {
	res, err := cToGoCollection(c).MDeleteDocument(cToGoStrings(ids, idsCount), SetQueryOptions(options))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_collection_m_get_document
func kuzzle_wrapper_collection_m_get_document(c *C.collection, ids **C.char, idsCount C.uint, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).MGetDocument(cToGoStrings(ids, idsCount), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_replace_document
func kuzzle_wrapper_collection_m_replace_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).MReplaceDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}

//export kuzzle_wrapper_collection_m_update_document
func kuzzle_wrapper_collection_m_update_document(c *C.collection, documents **C.document, docCount C.uint, options *C.query_options) *C.search_result {
	res, err := cToGoCollection(c).MUpdateDocument(cToGoDocuments(c, documents, docCount), SetQueryOptions(options))
	return goToCSearchResult(c, res, err)
}
