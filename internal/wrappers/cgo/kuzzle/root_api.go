package main

/*
  #cgo CFLAGS: -I../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include "kuzzlesdk.h"
  #include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"strconv"
	"time"
	"unsafe"
)

//export kuzzle_wrapper_list_collections
func kuzzle_wrapper_list_collections(k *C.kuzzle, index *C.char, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ListCollections(
		C.GoString(index),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_list_indexes
func kuzzle_wrapper_list_indexes(k *C.kuzzle, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ListIndexes(SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_wrapper_create_index
func kuzzle_wrapper_create_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).CreateIndex(
		C.GoString(index),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_refresh_index
func kuzzle_wrapper_refresh_index(k *C.kuzzle, index *C.char, options *C.query_options) *C.shards_result {
	result := (*C.shards_result)(C.calloc(1, C.sizeof_shards_result))
	opts := SetQueryOptions(options)

	shards, err := (*kuzzle.Kuzzle)(k.instance).RefreshIndex(C.GoString(index), opts)
	if err != nil {
		Set_shards_result_error(result, err)
		return result
	}

	result.result = goToCShards(shards)

	return result
}

//export kuzzle_wrapper_set_auto_refresh
func kuzzle_wrapper_set_auto_refresh(k *C.kuzzle, index *C.char, auto_refresh C.bool, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).SetAutoRefresh(
		C.GoString(index),
		bool(auto_refresh),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_get_auto_refresh
func kuzzle_wrapper_get_auto_refresh(k *C.kuzzle, index *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetAutoRefresh(
		C.GoString(index),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_get_server_info
func kuzzle_wrapper_get_server_info(k *C.kuzzle, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetServerInfo(SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_get_all_statistics
func kuzzle_wrapper_get_all_statistics(k *C.kuzzle, options *C.query_options) *C.all_statistics_result {
	result := (*C.all_statistics_result)(C.calloc(1, C.sizeof_statistics_result))
	opts := SetQueryOptions(options)

	stats, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)

	if err != nil {
		Set_all_statistics_error(result, err)
		return result
	}

	result.res = (*C.statistics)(C.calloc(C.size_t(len(stats)), C.sizeof_statistics_ptr))
	result.res_size = C.int(len(stats) - 1)
	statistics := (*[1<<30 - 1]*C.statistics)(unsafe.Pointer(result.res))[:len(stats)]

	for i, stat := range stats {
		statistics[i] = goToCStatistics(stat, err)
	}

	return result
}

//export kuzzle_wrapper_get_statistics
func kuzzle_wrapper_get_statistics(k *C.kuzzle, timestamp C.time_t, options *C.query_options) *C.statistics_result {
	result := (*C.statistics_result)(C.calloc(1, C.sizeof_statistics_result))
	opts := SetQueryOptions(options)

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&timestamp)), 10, 64)
	tm := time.Unix(t, 0)

	res, err := (*kuzzle.Kuzzle)(k.instance).GetStatistics(&tm, opts)

	result.result = goToCStatistics(res, err)

	return result
}

//export kuzzle_wrapper_now
func kuzzle_wrapper_now(k *C.kuzzle, options *C.query_options) *C.int_result {
	time, err := (*kuzzle.Kuzzle)(k.instance).Now(SetQueryOptions(options))

	return goToCIntResult(time, err)
}
