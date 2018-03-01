package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
    #include "sdk_wrappers_internal.h"
*/
import "C"

import (
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/realtime"
)

// map which stores instances to keep references in case the gc passes
var realtimeInstances map[interface{}]bool

// register new instance to the instances map
func registerRealtime(instance interface{}) {
	realtimeInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterRealtime
func unregisterRealtime(rt *C.realtime) {
	delete(realtimeInstances, (*realtime.Realtime)(rt.instance))
}

// Allocates memory
//export kuzzle_new_realtime
func kuzzle_new_realtime(rt *C.realtime, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	gort := realtime.NewRealtime(kuz)

	if realtimeInstances == nil {
		realtimeInstances = make(map[interface{}]bool)
	}

	rt.instance = unsafe.Pointer(gort)
	rt.kuzzle = k

	registerRealtime(rt)
}

//export kuzzle_realtime_count
func kuzzle_realtime_count(rt *C.realtime, index, collection, roomId *C.char) *C.int_result {
	res, err := (*realtime.Realtime)(rt.instance).Count(C.GoString(index), C.GoString(collection), C.GoString(roomId))
	return goToCIntResult(res, err)
}
