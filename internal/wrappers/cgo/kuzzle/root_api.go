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
	"strconv"
	"time"
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_list_collections
func kuzzle_list_collections(k *C.kuzzle, index *C.char, options *C.query_options) *C.collection_entry_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ListCollections(
		C.GoString(index),
		SetQueryOptions(options))

	return goToCCollectionListResult(res, err)
}

//export kuzzle_list_indexes
func kuzzle_list_indexes(k *C.kuzzle, options *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ListIndexes(SetQueryOptions(options))

	return goToCStringArrayResult(res, err)
}

//export kuzzle_get_auto_refresh
func kuzzle_get_auto_refresh(k *C.kuzzle, index *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetAutoRefresh(
		C.GoString(index),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_get_server_info
func kuzzle_get_server_info(k *C.kuzzle, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetServerInfo(SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_get_all_statistics
func kuzzle_get_all_statistics(k *C.kuzzle, options *C.query_options) *C.all_statistics_result {
	result := (*C.all_statistics_result)(C.calloc(1, C.sizeof_statistics_result))
	opts := SetQueryOptions(options)

	stats, err := (*kuzzle.Kuzzle)(k.instance).GetAllStatistics(opts)

	if err != nil {
		Set_all_statistics_error(result, err)
		return result
	}

	result.result = (*C.statistics)(C.calloc(C.size_t(len(stats)), C.sizeof_statistics))
	result.result_length = C.size_t(len(stats))
	statistics := (*[1<<30 - 1]C.statistics)(unsafe.Pointer(result.result))[:len(stats):len(stats)]

	for i, stat := range stats {
		fillStatistics(stat, &statistics[i])
	}

	return result
}

//export kuzzle_get_statistics
func kuzzle_get_statistics(k *C.kuzzle, start_time C.time_t, stop_time C.time_t, options *C.query_options) *C.statistics_result {
	result := (*C.statistics_result)(C.calloc(1, C.sizeof_statistics_result))
	opts := SetQueryOptions(options)

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&start_time)), 10, 64)
	start := time.Unix(t, 0)
	t, _ = strconv.ParseInt(C.GoString(C.ctime(&stop_time)), 10, 64)
	stop := time.Unix(t, 0)

	res, err := (*kuzzle.Kuzzle)(k.instance).GetStatistics(&start, &stop, opts)

	if err != nil {
		Set_statistics_error(result, err)
		return result
	}

	fillStatistics(res, result.result)

	return result
}

//export kuzzle_now
func kuzzle_now(k *C.kuzzle, options *C.query_options) *C.date_result {
	time, err := (*kuzzle.Kuzzle)(k.instance).Now(SetQueryOptions(options))

	return goToCDateResult(time, err)
}
