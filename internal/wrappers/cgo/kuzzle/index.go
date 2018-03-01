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
func unregisterIndex(i *C.index) {
	delete(indexInstances, (*indexPkg.Index)(i.instance))
}

// Allocates memory
//export kuzzle_new_index
func kuzzle_new_index(i *C.index, k *C.kuzzle) {
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
func kuzzle_index_create(i *C.index, index string) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Create(index)
	return goToCVoidResult(err)
}

//export kuzzle_index_delete
func kuzzle_index_delete(i *C.index, index string) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Delete(index)
	return goToCVoidResult(err)
}

//export kuzzle_index_mdelete
func kuzzle_index_mdelete(i *C.index, indexes []string) *C.string_array_result {
	res, err := (*indexPkg.Index)(i.instance).MDelete(indexes)
	return goToCStringArrayResult(res, err)
}

//export kuzzle_index_exists
func kuzzle_index_exists(i *C.index, index string) *C.bool_result {
	res, err := (*indexPkg.Index)(i.instance).Exists(index)
	return goToCBoolResult(res, err)
}

//export kuzzle_index_refresh
func kuzzle_index_refresh(i *C.index, index string) *C.void_result {
	err := (*indexPkg.Index)(i.instance).Refresh(index)
	return goToCVoidResult(err)
}

//export kuzzle_index_refresh_internal
func kuzzle_index_refresh_internal(i *C.index, index string) *C.void_result {
	err := (*indexPkg.Index)(i.instance).RefreshInternal()
	return goToCVoidResult(err)
}

//export kuzzle_index_set_auto_refresh
func kuzzle_index_set_auto_refresh(i *C.index, index string, autoRefresh bool) *C.void_result {
	err := (*indexPkg.Index)(i.instance).SetAutoRefresh(index, autoRefresh)
	return goToCVoidResult(err)
}
