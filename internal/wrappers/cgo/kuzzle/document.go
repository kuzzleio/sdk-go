package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
    #include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_new_document
func kuzzle_wrapper_new_document(c *C.collection) *C.document {
	return goToCDocument(c, cToGoCollection(c).Document(), nil)
}

//export kuzzle_wrapper_document_subscribe
// TODO loop and close on Unsubscribe
func kuzzle_wrapper_document_subscribe(d *C.document, options *C.room_options, cb unsafe.Pointer) {
	c := make(chan *types.KuzzleNotification)
	cToGoDocument(d._collection, d).Subscribe(SetRoomOptions(options), c)

	go func() {
		res := <-c
		C.call_notification_result(cb, goToCNotificationResult(res))
	}()
}

// Does not re-allocate the document
// export kuzzle_wrapper_document_save
func kuzzle_wrapper_document_save(d *C.document, options *C.query_options) *C.document_result {
	_, err := cToGoDocument(d._collection, d).Save(SetQueryOptions(options))
	return currentDocumentResult(d, err)
}

// export kuzzle_wrapper_document_refresh
func kuzzle_wrapper_document_refresh(d *C.document, options *C.query_options) *C.document_result {
	res, err := cToGoDocument(d._collection, d).Refresh(SetQueryOptions(options))
	return goToCDocumentResult(d._collection, res, err)
}

//export kuzzle_wrapper_document_set_headers
func kuzzle_wrapper_document_set_headers(d *C.document, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		cToGoDocument(d._collection, d).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}

// export kuzzle_wrapper_document_publish
func kuzzle_wrapper_document_publish(d *C.document, options *C.query_options) *C.bool_result {
	res, err := cToGoDocument(d._collection, d).Publish(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

// export kuzzle_wrapper_document_exists
func kuzzle_wrapper_document_exists(d *C.document, options *C.query_options) *C.bool_result {
	res, err := cToGoDocument(d._collection, d).Exists(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

// export kuzzle_wrapper_document_delete
func kuzzle_wrapper_document_delete(d *C.document, options *C.query_options) *C.string_result {
	res, err := cToGoDocument(d._collection, d).Delete(SetQueryOptions(options))
	return goToCStringResult(&res, err)
}

// Allocates memory for result, not document
func currentDocumentResult(d *C.document, err error) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))

	if err != nil {
		Set_document_error(result, err)
	}

	result.result = d

	return result
}
