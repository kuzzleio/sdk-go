// Copyright 2015-2018 Kuzzle
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
	#include <stdlib.h>
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"time"

	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_new_options
func kuzzle_new_options() *C.options {
	copts := (*C.options)(C.calloc(1, C.sizeof_options))
	opts := types.NewOptions()

	copts.queue_ttl = C.uint(opts.QueueTTL())
	copts.queue_max_size = C.ulong(opts.QueueMaxSize())
	copts.auto_queue = C.bool(opts.AutoQueue())
	copts.auto_reconnect = C.bool(opts.AutoReconnect())
	copts.auto_replay = C.bool(opts.AutoReplay())
	copts.auto_resubscribe = C.bool(opts.AutoResubscribe())
	copts.reconnection_delay = C.ulong(opts.ReconnectionDelay())
	copts.replay_interval = C.ulong(opts.ReplayInterval())

	if opts.OfflineMode() == 1 {
		copts.offline_mode = C.MANUAL
	} else {
		copts.offline_mode = C.AUTO
	}

	refresh := opts.Refresh()
	if len(refresh) > 0 {
		copts.refresh = C.CString(refresh)
	}

	return copts
}

func SetQueryOptions(options *C.query_options) (opts types.QueryOptions) {
	if options == nil {
		return
	}

	opts = types.NewQueryOptions()

	opts.SetQueuable(bool(options.queuable))
	opts.SetWithdist(bool(options.withdist))
	opts.SetWithcoord(bool(options.withcoord))
	opts.SetFrom(int(options.from))
	opts.SetSize(int(options.size))
	opts.SetScroll(C.GoString(options.scroll))
	opts.SetScrollId(C.GoString(options.scroll_id))
	opts.SetRefresh(C.GoString(options.refresh))
	opts.SetIfExist(C.GoString(options.if_exist))
	opts.SetRetryOnConflict(int(options.retry_on_conflict))
	volatiles := types.VolatileData(C.GoString(options.volatiles))
	if volatiles != nil {
		opts.SetVolatile(volatiles)
	}

	return
}

func SetOptions(options *C.options) (opts types.Options) {
	if options == nil {
		return
	}

	opts = types.NewOptions()

	opts.SetQueueTTL(time.Duration(uint16(options.queue_ttl)))
	opts.SetQueueMaxSize(int(options.queue_max_size))
	opts.SetOfflineMode(int(options.offline_mode))

	opts.SetAutoQueue(bool(options.auto_queue))
	opts.SetAutoReconnect(bool(options.auto_reconnect))
	opts.SetAutoReplay(bool(options.auto_replay))
	opts.SetAutoResubscribe(bool(options.auto_resubscribe))
	opts.SetReconnectionDelay(time.Duration(int(options.reconnection_delay)))
	opts.SetReplayInterval(time.Duration(int(options.replay_interval)))
	opts.SetRefresh(C.GoString(options.refresh))

	return
}

func SetRoomOptions(options *C.room_options) (opts types.RoomOptions) {
	opts = types.NewRoomOptions()

	if options != nil {
		opts.SetScope(C.GoString(options.scope))
		opts.SetState(C.GoString(options.state))
		opts.SetUsers(C.GoString(options.user))

		opts.SetSubscribeToSelf(bool(options.subscribe_to_self))

		if options.volatiles != nil {
			opts.SetVolatile(types.VolatileData(C.GoString(options.volatiles)))
		}
	}
	return
}
