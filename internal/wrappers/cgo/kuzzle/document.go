package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
    #include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"unsafe"

	col "github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
)

// map which stores instances to keep references in case the gc passes
var documentInstances map[interface{}]bool

// register new instance to the instances map
func registerDocument(instance interface{}) {
	documentInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterDocument
func unregisterDocument(d *C.document) {
	delete(documentInstances, (*col.Document)(d.instance))
}

//export kuzzle_new_document
func kuzzle_new_document(d *C.document, c *C.collection, id *C.char, content *C.json_object) {
	doc := col.NewDocument((*col.Collection)(c.instance), C.GoString(id))

	if content != nil {
		doc.SetContent(JsonCConvert(content).(map[string]interface{}), true)
	}
	if documentInstances == nil {
		documentInstances = make(map[interface{}]bool)
	}

	registerDocument(doc)
	d._collection = c
	d.instance = unsafe.Pointer(doc)
}

//export kuzzle_document_subscribe
// TODO close on Unsubscribe
func kuzzle_document_subscribe(d *C.document, options *C.room_options, cb C.kuzzle_notification_listener, data unsafe.Pointer) *C.room_result {
	c := make(chan types.KuzzleNotification)
	goroom, _ := cToGoDocument(d._collection, d).Subscribe(SetRoomOptions(options), c)

	go func() {
		for {
			res := <-c
			C.kuzzle_notify(cb, goToCNotificationResult(&res), data)
		}
	}()
	<-goroom.ResponseChannel()

	room := (*C.room)(C.calloc(1, C.sizeof_room))
	room.instance = unsafe.Pointer(goroom)
	registerRoom(room)
	result := (*C.room_result)(C.calloc(1, C.sizeof_room_result))

	result.result = room
	return result
}

// Does not re-allocate the document
//export kuzzle_document_save
func kuzzle_document_save(d *C.document, options *C.query_options) *C.document_result {
	_, err := cToGoDocument(d._collection, d).Save(SetQueryOptions(options))
	return currentDocumentResult(d, err)
}

//export kuzzle_document_refresh
func kuzzle_document_refresh(d *C.document, options *C.query_options) *C.document_result {
	res, err := cToGoDocument(d._collection, d).Refresh(SetQueryOptions(options))
	return goToCDocumentResult(d._collection, res, err)
}

//export kuzzle_document_publish
func kuzzle_document_publish(d *C.document, options *C.query_options) *C.bool_result {
	//res, err := cToGoDocument(d._collection, d).Publish(SetQueryOptions(options))

	res, err := (*col.Document)(d.instance).Publish(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_document_exists
func kuzzle_document_exists(d *C.document, options *C.query_options) *C.bool_result {
	res, err := cToGoDocument(d._collection, d).Exists(SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_document_delete
func kuzzle_document_delete(d *C.document, options *C.query_options) *C.string_result {
	res, err := (*col.Document)(d.instance).Delete(SetQueryOptions(options))
	return goToCStringResult(&res, err)
}

//export kuzzle_document_set_content
func kuzzle_document_set_content(d *C.document, content *C.json_object, replace C.bool) {
	(*col.Document)(d.instance).SetContent(JsonCConvert(content).(map[string]interface{}), bool(replace))
}

//export kuzzle_document_get_content
func kuzzle_document_get_content(d *C.document) *C.json_object {
	r, _ := goToCJson((*col.Document)(d.instance).Content)
	return r
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
