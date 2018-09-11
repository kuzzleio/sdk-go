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

package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MDelete deletes multiple documents at once
func (d *Document) MDelete(index string, collection string, ids []string, options types.QueryOptions) ([]string, error) {
	if index == "" {
		return nil, types.NewError("Document.MDelete: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.MDelete: collection required", 400)
	}

	if len(ids) == 0 {
		return nil, types.NewError("Document.MDelete: ids filled array required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	type body struct {
		Ids []string `json:"ids"`
	}

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mDelete",
		Body:       &body{ids},
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return nil, res.Error
	}

	var mDeleted []string
	json.Unmarshal(res.Result, &mDeleted)

	return mDeleted, nil
}
