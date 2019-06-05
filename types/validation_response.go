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

package types

import (
	"encoding/json"
	"fmt"
)

// ValidationResponse contains a validation response
type ValidationResponse struct {
	Valid       bool     `json:"valid"`
	Details     []string `json:"details"`
	Description string   `json:"description"`
}

func NewValidationResponse(validationResponse json.RawMessage) (*ValidationResponse, error) {
	vr := &ValidationResponse{}

	err := json.Unmarshal(validationResponse, vr)

	if err != nil {
		return nil, NewError(fmt.Sprintf("Unable to parse response: %s\n%s", err.Error(), validationResponse), 500)
	}

	return vr, nil
}
