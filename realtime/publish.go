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

package realtime

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// Publish sends a real-time message to Kuzzle
func (r *Realtime) Publish(index string, collection string, body json.RawMessage, options types.QueryOptions) error {
	if (index == "" || collection == "") || body == nil {
		return types.NewError("Realtime.Publish: index, collection and body required", 400)
	}

	query := &types.KuzzleRequest{
		Controller: "realtime",
		Action:     "publish",
		Index:      index,
		Collection: collection,
		Body:       body,
	}

	result := make(chan *types.KuzzleResponse)

	go r.k.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return res.Error
	}

	return nil
}
