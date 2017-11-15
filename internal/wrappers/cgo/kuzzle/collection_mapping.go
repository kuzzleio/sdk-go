package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_new_mapping
func kuzzle_wrapper_new_mapping(c *C.collection) *C.mapping {
	cm := (*C.mapping)(C.calloc(1, C.sizeof_mapping))
	cm.mapping = C.json_object_new_object()
	cm.collection = c

	return cm
}

//export kuzzle_wrapper_collection_get_mapping
func kuzzle_wrapper_collection_get_mapping(c *C.collection, options *C.query_options) *C.mapping_result {
	res, err := cToGoCollection(c).GetMapping(SetQueryOptions(options))
	return goToCMappingResult(c, res, err)
}

//export kuzzle_wrapper_mapping_apply
func kuzzle_wrapper_mapping_apply(cm *C.mapping, options *C.query_options) *C.bool_result {
	_, err := cToGoMapping(cm).Apply(SetQueryOptions(options))
	return goToCBoolResult(true, err)
}

//export kuzzle_wrapper_mapping_refresh
func kuzzle_wrapper_mapping_refresh(cm *C.mapping, options *C.query_options) *C.bool_result {
	_, err := cToGoMapping(cm).Refresh(SetQueryOptions(options))
	return goToCBoolResult(true, err)
}

//export kuzzle_wrapper_mapping_set
func kuzzle_wrapper_mapping_set(cm *C.mapping, jMap *C.json_object) {
	var mappings types.MappingFields

	if JsonCType(jMap) == C.json_type_object {
		jsonString := []byte(C.GoString(C.json_object_to_json_string(jMap)))
		json.Unmarshal(jsonString, &mappings)
	}

	cToGoMapping(cm).Set(&mappings)

	return
}

//export kuzzle_wrapper_mapping_set_headers
func kuzzle_wrapper_mapping_set_headers(cm *C.mapping, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		cToGoMapping(cm).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}
