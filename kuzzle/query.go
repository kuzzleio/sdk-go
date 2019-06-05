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
	"fmt"
	"time"

	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

// Query this is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query *types.KuzzleRequest, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse) {
	u, _ := uuid.NewV4()
	requestId := u.String()

	if query.RequestId == "" {
		query.RequestId = requestId
	}

	type body struct{}

	if query.Body == nil {
		query.Body = json.RawMessage("{}")
	}

	if options == nil {
		options = types.NewQueryOptions()
	}

	if options.Volatile() != nil {
		query.Volatile = options.Volatile()

		mapped := make(map[string]interface{})
		json.Unmarshal(query.Volatile, &mapped)

		mapped["sdkVersion"] = version

		query.Volatile, _ = json.Marshal(mapped)

	} else {
		vol := fmt.Sprintf(`{"sdkVersion": "%s"}`, version)
		query.Volatile = types.VolatileData(vol)
	}

	jsonRequest, _ := json.Marshal(query)

	out := map[string]interface{}{}
	err := json.Unmarshal(jsonRequest, &out)

	if err != nil {
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Error: types.NewError(err.Error())}
		}
		return
	}

	refresh := options.Refresh()
	if refresh != "" {
		out["refresh"] = refresh
	}

	out["from"] = options.From()
	out["size"] = options.Size()

	scroll := options.Scroll()
	if scroll != "" {
		out["scroll"] = scroll
	}

	scrollId := options.ScrollId()
	if scrollId != "" {
		out["scrollId"] = scrollId
	}

	retryOnConflict := options.RetryOnConflict()
	if retryOnConflict > 0 {
		out["retryOnConflict"] = retryOnConflict
	}

	jwt := k.Jwt()
	if jwt != "" {
		out["jwt"] = jwt
	}

	finalRequest, err := json.Marshal(out)

	if err != nil {
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Error: types.NewError(err.Error())}
		}
		return
	}

	queuable := options == nil || options.Queuable()
	queuable = queuable && k.queueFilter(finalRequest)

	if k.queuing {
		if queuable {
			k.cleanQueue()
			qo := &types.QueryObject{
				Timestamp: time.Now(),
				ResChan:   responseChannel,
				Query:     finalRequest,
				RequestId: requestId,
				Options:   options,
			}
			k.offlineQueue = append(k.offlineQueue, qo)
			k.EmitEvent(event.OfflineQueuePush, qo)
			return
		}

		k.EmitEvent(event.Discarded, finalRequest)
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Status: 400, Error: types.NewError("Unable to execute request: not connected to a Kuzzle server.\nDiscarded request: "+string(finalRequest), 400)}
		}
		return
	}

	err = k.protocol.Send(finalRequest, options, responseChannel, requestId)

	if err != nil {
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Error: types.NewError(err.Error())}
		}
		return
	}
}
