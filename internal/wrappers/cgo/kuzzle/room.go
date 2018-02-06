package main

import (
	"unsafe"

	"github.com/kuzzleio/sdk-go/collection"
)

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"

// map which stores instances to keep references in case the gc passes
var roomInstances map[interface{}]bool

// register new instance to the instances map
func registerRoom(instance interface{}) {
	if roomInstances == nil {
		roomInstances = make(map[interface{}]bool)
	}
	roomInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterRoom
func unregisterRoom(r *C.room) {
	delete(roomInstances, (*collection.Room)(r.instance))
}

//export room_new_room
func room_new_room(room *C.room, col *C.collection, filters *C.json_object, options *C.room_options) {
	opts := SetRoomOptions(options)

	r := collection.NewRoom((*collection.Collection)(col.instance), JsonCConvert(filters), opts)

	registerRoom(room)
	room.instance = unsafe.Pointer(r)
	room.filters = filters
	room.options = options
}

//export room_count
func room_count(room *C.room) *C.int_result {
	res, err := (*collection.Room)(room.instance).Count()
	return goToCIntResult(res, err)
}
