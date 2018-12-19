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

package kuzzle

import (
	"encoding/json"
	"time"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
)

// FlushQueue empties the offline queue without replaying it.
func (k *Kuzzle) FlushQueue() {
	k.offlineQueue = nil
}

// PlayQueue replays the requests queued during offline mode.
func (k *Kuzzle) PlayQueue() {
	if k.protocol.IsReady() {
		k.cleanQueue()
		k.dequeue()
	}
}

// StartQueuing start the requests queuing.
func (k *Kuzzle) StartQueuing() {
	k.queuing = true
}

// StopQueuing stop the requests queuing.
func (k *Kuzzle) StopQueuing() {
	k.queuing = false
}

// Clean up the queue, ensuring the queryTTL and queryMaxSize properties are respected
func (k *Kuzzle) cleanQueue() {
	now := time.Now()
	now = now.Add(-k.queueTTL * time.Millisecond)

	// Clean queue of timed out query
	if k.queueTTL > 0 {
		var query *types.QueryObject
		for _, query = range k.offlineQueue {
			if query.Timestamp.Before(now) {
				k.offlineQueue = k.offlineQueue[1:]
			} else {
				break
			}
		}
	}

	if k.queueMaxSize > 0 && len(k.offlineQueue) > k.queueMaxSize {
		for len(k.offlineQueue) > k.queueMaxSize {
			eventListener := k.eventListeners[event.OfflineQueuePop]
			for c := range eventListener {
				json, _ := json.Marshal(k.offlineQueue[0])
				c <- json
			}

			eventListener = k.eventListenersOnce[event.OfflineQueuePop]
			for c := range eventListener {
				json, _ := json.Marshal(k.offlineQueue[0])
				c <- json
				delete(k.eventListenersOnce[event.OfflineQueuePop], c)
			}

			k.offlineQueue = k.offlineQueue[1:]
		}
	}
}

func (k *Kuzzle) mergeOfflineQueueWithLoader() error {
	type query struct {
		requestId  string `json:"requestId"`
		controller string `json:"controller"`
		action     string `json:"action"`
	}

	additionalOfflineQueue := k.offlineQueueLoader.Load()

	for _, additionalQuery := range additionalOfflineQueue {
		for _, offlineQuery := range k.offlineQueue {
			q := query{}
			json.Unmarshal(additionalQuery.Query, &q)
			if q.requestId != "" || q.action != "" || q.controller != "" {
				offlineQ := query{}
				json.Unmarshal(offlineQuery.Query, &offlineQ)
				if q.requestId != offlineQ.requestId {
					k.offlineQueue = append(k.offlineQueue, additionalQuery)
				} else {
					additionalOfflineQueue = additionalOfflineQueue[:1]
				}
			} else {
				return types.NewError("Invalid offline queue request. One or more missing properties: requestId, action, controller.")
			}
		}
	}
	return nil
}

func (k *Kuzzle) dequeue() error {
	if k.offlineQueueLoader != nil {
		err := k.mergeOfflineQueueWithLoader()
		if err != nil {
			return err
		}
	}
	// Example from sdk where we have a good use of _
	if len(k.offlineQueue) > 0 {
		for _, query := range k.offlineQueue {
			k.protocol.Send(query.Query, query.Options, query.ResChan, query.RequestId)
			k.offlineQueue = k.offlineQueue[:1]
			k.EmitEvent(event.OfflineQueuePop, query)
			time.Sleep(k.replayInterval * time.Millisecond)
			k.offlineQueue = k.offlineQueue[:1]
		}
	} else {
		k.queuing = false
	}
	return nil
}
