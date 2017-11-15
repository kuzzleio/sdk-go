package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
*/
import "C"

//export kuzzle_wrapper_collection_delete_specifications
func kuzzle_wrapper_collection_delete_specifications(c *C.collection, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).DeleteSpecifications(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_collection_get_specifications
func kuzzle_wrapper_collection_get_specifications(c *C.collection, options *C.query_options) *C.specification_result {
	res, err := cToGoCollection(c).GetSpecifications(SetQueryOptions(options))
	return goToCSpecificationResult(res.Validation, err)
}

//export kuzzle_wrapper_collection_scroll_specifications
func kuzzle_wrapper_collection_scroll_specifications(c *C.collection, scrollId *C.char, options *C.query_options) *C.specification_search_result {
	res, err := cToGoCollection(c).ScrollSpecifications(C.GoString(scrollId), SetQueryOptions(options))
	return goToCSpecificationSearchResult(res, err)
}

//export kuzzle_wrapper_collection_search_specifications
func kuzzle_wrapper_collection_search_specifications(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.specification_search_result {
	res, err := cToGoCollection(c).SearchSpecifications(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCSpecificationSearchResult(res, err)
}

//export kuzzle_wrapper_collection_update_specifications
func kuzzle_wrapper_collection_update_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.specification_result {
	res, err := cToGoCollection(c).UpdateSpecifications(cToGoSpecification(specification), SetQueryOptions(options))
	return goToCSpecificationResult(res, err)
}

//export kuzzle_wrapper_collection_validate_specifications
func kuzzle_wrapper_collection_validate_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).ValidateSpecifications(cToGoSpecification(specification), SetQueryOptions(options))
	return goToCBoolResult(res.Valid, err)
}
