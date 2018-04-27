// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	"sync"
	"unsafe"

	indexPkg "github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

// map which stores instances to keep references in case the gc passes
var indexInstances sync.Map

//register new instance of index
func registerIndex(instance interface{}) {
	indexInstances.Store(instance, true)
}

// unregister an instance from the instances map
//export unregisterIndex
func unregisterIndex(i *C.kuzzle_index) {
	indexInstances.Delete(i)
}

// Allocates memory
//export kuzzle_new_index
func kuzzle_new_index(i *C.kuzzle_index, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	index := indexPkg.NewIndex(kuz)

	i.instance = unsafe.Pointer(index)
	i.kuzzle = k

	registerIndex(i)
}

//export kuzzle_index_create
func kuzzle_index_create(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.error_result {
	err := (*indexPkg.Index)(i.instance).Create(C.GoString(index), SetQueryOptions(options))
	return goToCErrorResult(err)
}

//export kuzzle_index_delete
func kuzzle_index_delete(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.error_result {
	err := (*indexPkg.Index)(i.instance).Delete(C.GoString(index), SetQueryOptions(options))
	return goToCErrorResult(err)
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
func kuzzle_index_refresh(i *C.kuzzle_index, index *C.char, options *C.query_options) *C.error_result {
	err := (*indexPkg.Index)(i.instance).Refresh(C.GoString(index), SetQueryOptions(options))
	return goToCErrorResult(err)
}

//export kuzzle_index_refresh_internal
func kuzzle_index_refresh_internal(i *C.kuzzle_index, options *C.query_options) *C.error_result {
	err := (*indexPkg.Index)(i.instance).RefreshInternal(SetQueryOptions(options))
	return goToCErrorResult(err)
}

//export kuzzle_index_set_auto_refresh
func kuzzle_index_set_auto_refresh(i *C.kuzzle_index, index *C.char, autoRefresh C.bool, options *C.query_options) *C.error_result {
	err := (*indexPkg.Index)(i.instance).SetAutoRefresh(C.GoString(index), bool(autoRefresh), SetQueryOptions(options))
	return goToCErrorResult(err)
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
