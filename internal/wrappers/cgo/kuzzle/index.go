package main

/*
	#cgo CFLAGS: -I../../headers
	#include <errno.h>
	#include <stdlib.h>
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"unsafe"

	indexPkg "github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

// map which stores instances to keep references in case the gc passes
var indexInstances map[interface{}]bool

//register new instance of index
func registerIndex(instance interface{}) {
	indexInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterIndex
func unregisterIndex(i *C.kuzzle_index) {
	delete(indexInstances, (*indexPkg.Index)(i.instance))
}

// Allocates memory
//export kuzzle_new_index
func kuzzle_new_index(i *C.kuzzle_index, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	index := indexPkg.NewIndex(kuz)

	if indexInstances == nil {
		indexInstances = make(map[interface{}]bool)
	}

	i.instance = unsafe.Pointer(index)
	i.kuzzle = k

	registerIndex(i)
}

//export kuzzle_index_create
func kuzzle_index_create(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Create(C.GoString(index), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_index_delete
func kuzzle_index_delete(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Delete(C.GoString(index), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_index_mdelete
func kuzzle_index_mdelete(i *C.kuzzle_index, indexes **C.char, l C.size_t, options *C.query_options) *C.string_array_result {
	res, err := (*indexPkg.Index)(i.instance).MDelete(cToGoStrings(indexes, l), SetQueryOptions(options))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_index_exists
func kuzzle_index_exists(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.bool_result {
	res, err := (*indexPkg.Index)(i.instance).Exists(C.GoString(index), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_index_refresh
func kuzzle_index_refresh(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Refresh(C.GoString(index), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_index_refresh_internal
func kuzzle_index_refresh_internal(i *C.kuzzle_index, options *C.query_options) *C.void_result {
	err := (*indexPkg.Index)(i.instance).RefreshInternal(SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_index_set_auto_refresh
func kuzzle_index_set_auto_refresh(i *C.kuzzle_index, index *C.char, autoRefresh C.bool, options *C.query_options) *C.void_result {
	err := (*indexPkg.Index)(i.instance).SetAutoRefresh(C.GoString(index), bool(autoRefresh), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_index_get_auto_refresh
func kuzzle_index_get_auto_refresh(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.bool_result {
	res, err := (*indexPkg.Index)(i.instance).GetAutoRefresh(C.GoString(index), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_index_list
func kuzzle_index_list(i *C.kuzzle_index, options *C.query_options) *C.string_array_result {
	res, err := (*indexPkg.Index)(i.instance).List(SetQueryOptions(options))
	return goToCStringArrayResult(res, err)
}
