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

// MDelete deletes all matching indices at once
func (i *Index) MDelete(indexes []string, options types.QueryOptions) ([]string, error) {
	if len(indexes) == 0 {
		return nil, types.NewError("Index.MDelete: at least one index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Indexes []string `json:"indexes"`
	}

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "mDelete",
		Body:       &body{indexes},
	}
	go i.kuzzle.Query(query, options, result)

	res := <-result

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var deletedIndexes struct {
		Deleted []string
	}

	err := json.Unmarshal(res.Result, &deletedIndexes)
	if err != nil {
		return nil, types.NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), res.Result), 500)
	}

	return deletedIndexes.Deleted, nil
}
