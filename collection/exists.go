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

package collection

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// Exists checks if a collection exists.
func (dc *Collection) Exists(index string, collection string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Collection.Exists: index required", 400)
	}

	if collection == "" {
		return false, types.NewError("Collection.Exists: collection required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "collection",
		Action:     "exists",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var exists bool

	err := json.Unmarshal(res.Result, &exists)

	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return exists, nil
}
