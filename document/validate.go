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

// Validate validates data against existing validation rules.
func (d *Document) Validate(index string, collection string, body json.RawMessage, options types.QueryOptions) (bool, error) {
	if index == "" {
		return false, types.NewError("Document.Validate: index required", 400)
	}

	if collection == "" {
		return false, types.NewError("Document.Validate: collection required", 400)
	}

	if body == nil {
		return false, types.NewError("Document.Validate: body required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: collection,
		Index:      index,
		Controller: "document",
		Action:     "validate",
		Body:       body,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Error() != "" {
		return false, res.Error
	}

	var valid struct {
		Valid bool `json:"valid"`
	}
	json.Unmarshal(res.Result, &valid)

	return valid.Valid, nil
}
