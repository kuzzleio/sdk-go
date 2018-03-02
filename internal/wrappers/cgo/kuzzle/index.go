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
func kuzzle_index_create(i *C.kuzzle_index, index *C.char) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Create(C.GoString(index))
	return goToCVoidResult(err)
}

//export kuzzle_index_delete
func kuzzle_index_delete(i *C.kuzzle_index, index *C.char) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Delete(C.GoString(index))
	return goToCVoidResult(err)
}

//export kuzzle_index_mdelete
func kuzzle_index_mdelete(i *C.kuzzle_index, indexes **C.char, l C.size_t) *C.string_array_result {
	res, err := (*indexPkg.Index)(i.instance).MDelete(cToGoStrings(indexes, l))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_index_exists
func kuzzle_index_exists(i *C.kuzzle_index, index *C.char) *C.bool_result {
	res, err := (*indexPkg.Index)(i.instance).Exists(C.GoString(index))
	return goToCBoolResult(res, err)
}

//export kuzzle_index_refresh
func kuzzle_index_refresh(i *C.kuzzle_index, index *C.char) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Refresh(C.GoString(index))
	return goToCVoidResult(err)
}

//export kuzzle_index_refresh_internal
func kuzzle_index_refresh_internal(i *C.kuzzle_index) *C.void_result {
	err := (*indexPkg.Index)(i.instance).RefreshInternal()
	return goToCVoidResult(err)
}

//export kuzzle_index_set_auto_refresh
func kuzzle_index_set_auto_refresh(i *C.kuzzle_index, index *C.char, autoRefresh C.bool) *C.void_result {
	err := (*indexPkg.Index)(i.instance).SetAutoRefresh(C.GoString(index), bool(autoRefresh))
	return goToCVoidResult(err)
}
