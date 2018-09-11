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

package index

import (
	"encoding/json"
	"fmt"

	"github.com/kuzzleio/sdk-go/types"
)

// GetAutoRefresh returns the current autoRefresh status for the given index.
func (i *Index) GetAutoRefresh(index string, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Index.GetAutoRefresh: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Controller: "index",
		Action:     "getAutoRefresh",
	}

	go i.kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var autoR bool

	err := json.Unmarshal(res.Result, &autoR)
	if err != nil {
		return false, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}
	return autoR, nil

}
