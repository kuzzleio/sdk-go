package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
*/
import "C"
import "github.com/kuzzleio/sdk-go/collection"

//export kuzzle_collection_delete_specifications
func kuzzle_collection_delete_specifications(c *C.collection, options *C.query_options) *C.bool_result {
	res, err := (*collection.Collection)(c.instance).DeleteSpecifications(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_get_specifications
func kuzzle_collection_get_specifications(c *C.collection, options *C.query_options) *C.specification_result {
	res, err := (*collection.Collection)(c.instance).GetSpecifications(SetQueryOptions(options))
	return goToCSpecificationResult(res.Validation, err)
}

//export kuzzle_collection_scroll_specifications
func kuzzle_collection_scroll_specifications(c *C.collection, scrollId *C.char, options *C.query_options) *C.specification_search_result {
	res, err := (*collection.Collection)(c.instance).ScrollSpecifications(C.GoString(scrollId), SetQueryOptions(options))
	return goToCSpecificationSearchResult(res, err)
}

//export kuzzle_collection_search_specifications
func kuzzle_collection_search_specifications(c *C.collection, searchFilters *C.search_filters, options *C.query_options) *C.specification_search_result {
	res, err := (*collection.Collection)(c.instance).SearchSpecifications(cToGoSearchFilters(searchFilters), SetQueryOptions(options))
	return goToCSpecificationSearchResult(res, err)
}

//export kuzzle_collection_update_specifications
func kuzzle_collection_update_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.specification_result {
	res, err := (*collection.Collection)(c.instance).UpdateSpecifications(cToGoSpecification(specification), SetQueryOptions(options))
	return goToCSpecificationResult(res, err)
}

//export kuzzle_collection_validate_specifications
func kuzzle_collection_validate_specifications(c *C.collection, specification *C.specification, options *C.query_options) *C.bool_result {
	res, err := (*collection.Collection)(c.instance).ValidateSpecifications(cToGoSpecification(specification), SetQueryOptions(options))
	return goToCBoolResult(res.Valid, err)
}
