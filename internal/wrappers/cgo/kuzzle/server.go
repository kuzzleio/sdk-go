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
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/server"
)

// map which stores instances to keep references in case the gc passes
var serverInstances sync.Map

//register new instance of server
func registerServer(instance interface{}) {
	serverInstances.Store(instance, true)
}

// unregister an instance from the instances map
//export unregisterServer
func unregisterServer(s *C.server) {
	serverInstances.Delete(s)
}

// Allocates memory
//export kuzzle_new_server
func kuzzle_new_server(s *C.server, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	server := server.NewServer(kuz)

	s.instance = unsafe.Pointer(server)
	s.kuzzle = k

	registerServer(s)
}

//export kuzzle_admin_exists
func kuzzle_admin_exists(s *C.server, options *C.query_options) *C.bool_result {
	opts := SetQueryOptions(options)

	res, err := (*server.Server)(s.instance).AdminExists(opts)
	return goToCBoolResult(res, err)
}

//export kuzzle_get_all_stats
func kuzzle_get_all_stats(s *C.server, options *C.query_options) *C.string_result {
	opts := SetQueryOptions(options)

	stats, err := (*server.Server)(s.instance).GetAllStats(opts)

	str := string(stats)
	return goToCStringResult(&str, err)
}

//export kuzzle_get_stats
func kuzzle_get_stats(s *C.server, start_time C.time_t, stop_time C.time_t, options *C.query_options) *C.string_result {
	opts := SetQueryOptions(options)

	t, _ := strconv.ParseInt(C.GoString(C.ctime(&start_time)), 10, 64)
	start := time.Unix(t, 0)
	t, _ = strconv.ParseInt(C.GoString(C.ctime(&stop_time)), 10, 64)
	stop := time.Unix(t, 0)

	res, err := (*server.Server)(s.instance).GetStats(&start, &stop, opts)

	str := string(res)
	return goToCStringResult(&str, err)
}

//export kuzzle_get_last_stats
func kuzzle_get_last_stats(s *C.server, options *C.query_options) *C.string_result {
	opts := SetQueryOptions(options)

	res, err := (*server.Server)(s.instance).GetLastStats(opts)

	str := string(res)
	return goToCStringResult(&str, err)
}

//export kuzzle_get_config
func kuzzle_get_config(s *C.server, options *C.query_options) *C.string_result {
	res, err := (*server.Server)(s.instance).GetConfig(SetQueryOptions(options))

	str := string(res)
	return goToCStringResult(&str, err)
}

//export kuzzle_info
func kuzzle_info(s *C.server, options *C.query_options) *C.string_result {
	res, err := (*server.Server)(s.instance).Info(SetQueryOptions(options))

	str := string(res)
	return goToCStringResult(&str, err)
}

//export kuzzle_now
func kuzzle_now(s *C.server, options *C.query_options) *C.date_result {
	time, err := (*server.Server)(s.instance).Now(SetQueryOptions(options))

	return goToCDateResult(time, err)
}
