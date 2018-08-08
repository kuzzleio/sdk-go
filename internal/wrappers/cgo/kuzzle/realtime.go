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
	#include "kuzzlesdk.h"
    #include "sdk_wrappers_internal.h"
*/
import "C"

import (
	"encoding/json"
	"sync"
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/realtime"
	"github.com/kuzzleio/sdk-go/types"
)

// map which stores instances to keep references in case the gc passes
var realtimeInstances sync.Map

// register new instance to the instances map
func registerRealtime(instance interface{}, ptr unsafe.Pointer) {
	realtimeInstances.Store(instance, ptr)
}

// unregister an instance from the instances map
//export unregisterRealtime
func unregisterRealtime(rt *C.realtime) {
	realtimeInstances.Delete(rt)
}

// Allocates memory
//export kuzzle_new_realtime
func kuzzle_new_realtime(rt *C.realtime, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	gort := realtime.NewRealtime(kuz)

	ptr := unsafe.Pointer(gort)
	rt.instance = ptr
	rt.k = k

	registerRealtime(rt, ptr)
}

//export kuzzle_realtime_count
func kuzzle_realtime_count(rt *C.realtime, index, collection, roomId *C.char, options *C.query_options) *C.int_result {
	res, err := (*realtime.Realtime)(rt.instance).Count(C.GoString(index), C.GoString(collection), C.GoString(roomId), SetQueryOptions(options))
	return goToCIntResult(res, err)
}

//export kuzzle_realtime_list
func kuzzle_realtime_list(rt *C.realtime, index, collection *C.char, options *C.query_options) *C.string_result {
	res, err := (*realtime.Realtime)(rt.instance).List(C.GoString(index), C.GoString(collection), SetQueryOptions(options))
	var stringResult string
	json.Unmarshal(res, &stringResult)
	return goToCStringResult(&stringResult, err)
}

//export kuzzle_realtime_publish
func kuzzle_realtime_publish(rt *C.realtime, index, collection, body *C.char, options *C.query_options) *C.error_result {
	err := (*realtime.Realtime)(rt.instance).Publish(C.GoString(index), C.GoString(collection), json.RawMessage(C.GoString(body)), SetQueryOptions(options))
	return goToCErrorResult(err)
}

//export kuzzle_realtime_unsubscribe
func kuzzle_realtime_unsubscribe(rt *C.realtime, roomId *C.char, options *C.query_options) *C.error_result {
	err := (*realtime.Realtime)(rt.instance).Unsubscribe(C.GoString(roomId), SetQueryOptions(options))
	return goToCErrorResult(err)
}

//export kuzzle_realtime_subscribe
func kuzzle_realtime_subscribe(rt *C.realtime, index, collection, body *C.char, callback C.kuzzle_notification_listener, data unsafe.Pointer, options *C.room_options) *C.subscribe_result {
	c := make(chan types.KuzzleNotification)
	subRes, err := (*realtime.Realtime)(rt.instance).Subscribe(C.GoString(index), C.GoString(collection), json.RawMessage(C.GoString(body)), c, SetRoomOptions(options))

	if err != nil {
		return goToCSubscribeResult(subRes, err)
	}

	go func() {
		for {
			res, ok := <-c
			if ok == false {
				break
			}
			C.kuzzle_notify(callback, goToCNotificationResult(&res), data)
		}
	}()

	return goToCSubscribeResult(subRes, err)
}

//export kuzzle_realtime_validate
func kuzzle_realtime_validate(rt *C.realtime, index, collection, body *C.char, options *C.query_options) *C.bool_result {
	res, err := (*realtime.Realtime)(rt.instance).Validate(C.GoString(index), C.GoString(collection), C.GoString(body), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}
