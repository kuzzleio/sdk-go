package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include <string.h>
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"encoding/json"
	"time"
	"unsafe"

	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

// map which stores instances to keep references in case the gc passes
var instances map[interface{}]bool

// map which stores channel and function's pointers adresses for listeners
var listeners_list map[uintptr]chan<- interface{}

// register new instance to the instances map
func registerKuzzle(instance interface{}) {
	instances[instance] = true
}

// unregister an instance from the instances map
//export unregisterKuzzle
func unregisterKuzzle(k *C.kuzzle) {
	delete(instances, (*kuzzle.Kuzzle)(k.instance))
}

//export kuzzle_new_kuzzle
func kuzzle_new_kuzzle(k *C.kuzzle, host, protocol *C.char, options *C.options) {
	var c connection.Connection

	if instances == nil {
		instances = make(map[interface{}]bool)
	}

	if listeners_list == nil {
		listeners_list = make(map[uintptr]chan<- interface{})
	}

	opts := SetOptions(options)

	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), opts)
	}

	inst, err := kuzzle.NewKuzzle(c, opts)

	if err != nil {
		panic(err.Error())
	}

	registerKuzzle(inst)

	k.instance = unsafe.Pointer(inst)
	k.loader = nil
}

//export kuzzle_get_document_controller
func kuzzle_get_document_controller(k *C.kuzzle) *C.document {
	return (*C.document)(unsafe.Pointer((*kuzzle.Kuzzle)(k.instance).Document))
}

//export kuzzle_get_auth_controller
func kuzzle_get_auth_controller(k *C.kuzzle) *C.auth {
	return (*C.auth)(unsafe.Pointer((*kuzzle.Kuzzle)(k.instance).Auth))
}

//export kuzzle_get_index_controller
func kuzzle_get_index_controller(k *C.kuzzle) *C.kuzzle_index {
	return (*C.kuzzle_index)(unsafe.Pointer((*kuzzle.Kuzzle)(k.instance).Index))
}

//export kuzzle_get_server_controller
func kuzzle_get_server_controller(k *C.kuzzle) *C.server {
	return (*C.server)(unsafe.Pointer((*kuzzle.Kuzzle)(k.instance).Server))
}

// Allocates memory
//export kuzzle_connect
func kuzzle_connect(k *C.kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Connect()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_disconnect
func kuzzle_disconnect(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).Disconnect()
}

//export kuzzle_get_default_index
func kuzzle_get_default_index(k *C.kuzzle) *C.char {
	return C.CString((*kuzzle.Kuzzle)(k.instance).DefaultIndex())
}

//export kuzzle_set_default_index
func kuzzle_set_default_index(k *C.kuzzle, index *C.char) C.int {
	err := (*kuzzle.Kuzzle)(k.instance).SetDefaultIndex(C.GoString(index))
	if err != nil {
		return C.int(C.EINVAL)
	}

	return 0
}

//export kuzzle_get_offline_queue
func kuzzle_get_offline_queue(k *C.kuzzle) *C.offline_queue {
	result := (*C.offline_queue)(C.calloc(1, C.sizeof_offline_queue))

	offlineQueue := (*kuzzle.Kuzzle)(k.instance).OfflineQueue()
	result.queries_length = C.size_t(len(offlineQueue))

	result.queries = (**C.query_object)(C.calloc(result.queries_length, C.sizeof_query_object_ptr))
	queryObjects := (*[1<<30 - 1]*C.query_object)(unsafe.Pointer(result.queries))[:result.queries_length:result.queries_length]

	idx := 0
	for _, queryObject := range offlineQueue {
		queryObjects[idx] = (*C.query_object)(C.calloc(1, C.sizeof_query_object))
		queryObjects[idx].timestamp = C.ulonglong(queryObject.Timestamp.Unix())
		queryObjects[idx].request_id = C.CString(queryObject.RequestId)
		mquery, _ := json.Marshal(queryObject.Query)

		buffer := C.CString(string(mquery))
		queryObjects[idx].query = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))

		idx += 1
	}

	return result
}

//export kuzzle_flush_queue
func kuzzle_flush_queue(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).FlushQueue()
}

//export kuzzle_replay_queue
func kuzzle_replay_queue(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).ReplayQueue()
}

//export kuzzle_start_queuing
func kuzzle_start_queuing(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).StartQueuing()
}

//export kuzzle_stop_queuing
func kuzzle_stop_queuing(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).StopQueuing()
}

//export kuzzle_add_listener
// TODO loop and close on Unsubscribe
func kuzzle_add_listener(k *C.kuzzle, e C.int, cb C.kuzzle_event_listener, data unsafe.Pointer) {
	c := make(chan interface{})

	listeners_list[uintptr(unsafe.Pointer(cb))] = c
	(*kuzzle.Kuzzle)(k.instance).AddListener(int(e), c)
	go func() {
		res := <-c

		var jsonRes *C.json_object
		r, _ := json.Marshal(res)

		buffer := C.CString(string(r))
		jsonRes = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))

		C.kuzzle_trigger_event(e, cb, jsonRes, data)
	}()
}

//export kuzzle_once
func kuzzle_once(k *C.kuzzle, e C.int, cb C.kuzzle_event_listener, data unsafe.Pointer) {
	c := make(chan interface{})

	listeners_list[uintptr(unsafe.Pointer(cb))] = c
	(*kuzzle.Kuzzle)(k.instance).Once(int(e), c)
	go func() {
		res := <-c

		var jsonRes *C.json_object
		r, _ := json.Marshal(res)

		buffer := C.CString(string(r))
		jsonRes = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))

		C.kuzzle_trigger_event(e, cb, jsonRes, data)
	}()
}

//export kuzzle_listener_count
func kuzzle_listener_count(k *C.kuzzle, event C.int) int {
	return (*kuzzle.Kuzzle)(k.instance).ListenerCount(int(event))
}

//export kuzzle_remove_listener
func kuzzle_remove_listener(k *C.kuzzle, event C.int, cb unsafe.Pointer) {
	(*kuzzle.Kuzzle)(k.instance).RemoveListener(int(event), listeners_list[uintptr(cb)])
}

//export kuzzle_remove_all_listeners
func kuzzle_remove_all_listeners(k *C.kuzzle, event C.int) {
	(*kuzzle.Kuzzle)(k.instance).RemoveAllListeners(int(event))
}

//export kuzzle_get_auto_queue
func kuzzle_get_auto_queue(k *C.kuzzle) C.bool {
	return C.bool((*kuzzle.Kuzzle)(k.instance).AutoQueue())
}

//export kuzzle_set_auto_queue
func kuzzle_set_auto_queue(k *C.kuzzle, value C.bool) {
	(*kuzzle.Kuzzle)(k.instance).SetAutoQueue(bool(value))
}

//export kuzzle_get_auto_reconnect
func kuzzle_get_auto_reconnect(k *C.kuzzle) C.bool {
	return C.bool((*kuzzle.Kuzzle)(k.instance).AutoReconnect())
}

//export kuzzle_get_auto_resubscribe
func kuzzle_get_auto_resubscribe(k *C.kuzzle) C.bool {
	return C.bool((*kuzzle.Kuzzle)(k.instance).AutoResubscribe())
}

//export kuzzle_get_auto_replay
func kuzzle_get_auto_replay(k *C.kuzzle) C.bool {
	return C.bool((*kuzzle.Kuzzle)(k.instance).AutoReplay())
}

//export kuzzle_set_auto_replay
func kuzzle_set_auto_replay(k *C.kuzzle, value C.bool) {
	(*kuzzle.Kuzzle)(k.instance).SetAutoReplay(bool(value))
}

//export kuzzle_get_host
func kuzzle_get_host(k *C.kuzzle) *C.char {
	return C.CString((*kuzzle.Kuzzle)(k.instance).Host())
}

//export kuzzle_get_offline_queue_loader
func kuzzle_get_offline_queue_loader(k *C.kuzzle) C.kuzzle_offline_queue_loader {
	return k.loader
}

//export kuzzle_set_offline_queue_loader
func kuzzle_set_offline_queue_loader(k *C.kuzzle, loader C.kuzzle_offline_queue_loader) {
	k.loader = loader
}

//export kuzzle_get_port
func kuzzle_get_port(k *C.kuzzle) C.int {
	return C.int((*kuzzle.Kuzzle)(k.instance).Port())
}

//export kuzzle_get_queue_filter
func kuzzle_get_queue_filter(k *C.kuzzle) C.kuzzle_queue_filter {
	return k.filter
}

//export kuzzle_set_queue_filter
func kuzzle_set_queue_filter(k *C.kuzzle, f C.kuzzle_queue_filter) {
	k.filter = f

	if f != nil {
		filter := func(q []byte) bool {
			return bool(C.kuzzle_filter_query(f, (*C.char)(unsafe.Pointer(&q[0]))))
		}

		(*kuzzle.Kuzzle)(k.instance).SetQueueFilter(filter)
	} else {
		(*kuzzle.Kuzzle)(k.instance).SetQueueFilter(nil)
	}
}

//export kuzzle_get_queue_max_size
func kuzzle_get_queue_max_size(k *C.kuzzle) C.int {
	return C.int((*kuzzle.Kuzzle)(k.instance).QueueMaxSize())
}

//export kuzzle_set_queue_max_size
func kuzzle_set_queue_max_size(k *C.kuzzle, size C.int) {
	(*kuzzle.Kuzzle)(k.instance).SetQueueMaxSize(int(size))
}

//export kuzzle_get_queue_ttl
func kuzzle_get_queue_ttl(k *C.kuzzle) C.int {
	return C.int((*kuzzle.Kuzzle)(k.instance).QueueTTL())
}

//export kuzzle_set_queue_ttl
func kuzzle_set_queue_ttl(k *C.kuzzle, ttl C.int) {
	(*kuzzle.Kuzzle)(k.instance).SetQueueTTL(time.Duration(ttl))
}

//export kuzzle_get_replay_interval
func kuzzle_get_replay_interval(k *C.kuzzle) C.int {
	return C.int((*kuzzle.Kuzzle)(k.instance).ReplayInterval())
}

//export kuzzle_set_replay_interval
func kuzzle_set_replay_interval(k *C.kuzzle, interval C.int) {
	(*kuzzle.Kuzzle)(k.instance).SetReplayInterval(time.Duration(interval))
}

//export kuzzle_get_reconnection_delay
func kuzzle_get_reconnection_delay(k *C.kuzzle) C.int {
	return C.int((*kuzzle.Kuzzle)(k.instance).ReconnectionDelay())
}

//export kuzzle_get_ssl_connection
func kuzzle_get_ssl_connection(k *C.kuzzle) C.bool {
	return C.bool((*kuzzle.Kuzzle)(k.instance).SslConnection())
}

//export kuzzle_get_volatile
func kuzzle_get_volatile(k *C.kuzzle) *C.json_object {
	r, _ := goToCJson((*kuzzle.Kuzzle)(k.instance).Volatile())
	return r
}

//export kuzzle_set_volatile
func kuzzle_set_volatile(k *C.kuzzle, v *C.json_object) {
	(*kuzzle.Kuzzle)(k.instance).SetVolatile(JsonCConvert(v).(map[string]interface{}))
}

func main() {

}
