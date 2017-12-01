package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
)

// Allocates memory
//export kuzzle_new_collection
func kuzzle_new_collection(k *C.kuzzle, colName *C.char, index *C.char) *C.collection {
	col := (*C.collection)(C.calloc(1, C.sizeof_collection))
	col.index = C.CString(C.GoString(index))
	col.collection = C.CString(C.GoString(colName))
	col.kuzzle = k

	return col
}

//export kuzzle_collection_create
func kuzzle_collection_create(c *C.collection, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).Create(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_publish_message
func kuzzle_collection_publish_message(c *C.collection, message *C.json_object, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).PublishMessage(JsonCConvert(message).(map[string]interface{}), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_truncate
func kuzzle_collection_truncate(c *C.collection, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).Truncate(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_subscribe
// TODO loop and close on Unsubscribe
func kuzzle_collection_subscribe(col *C.collection, filters *C.search_filters, options *C.room_options, cb C.kuzzle_notification_listener) {
	c := make(chan *types.KuzzleNotification)
	cToGoCollection(col).Subscribe(cToGoSearchFilters(filters), SetRoomOptions(options), c)

	go func() {
		res := <-c
		C.kuzzle_notify(cb, goToCNotificationResult(res))
	}()
}
