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

import "encoding/json"

type KuzzleRequest struct {
	RequestId    string        `json:"requestId,omitempty"`
	Controller   string        `json:"controller,omitempty"`
	Action       string        `json:"action,omitempty"`
	Index        string        `json:"index,omitempty"`
	Collection   string        `json:"collection,omitempty"`
	Body         interface{}   `json:"body"`
	Id           string        `json:"_id,omitempty"`
	From         int           `json:"from"`
	Size         int           `json:"size"`
	Scroll       string        `json:"scroll,omitempty"`
	ScrollId     string        `json:"scrollId,omitempty"`
	Strategy     string        `json:"strategy,omitempty"`
	ExpiresIn    int           `json:"expiresIn,omitempty"`
	Volatile     VolatileData  `json:"volatile"`
	Scope        string        `json:"scope"`
	State        string        `json:"state"`
	Users        string        `json:"users"`
	Start        int           `json:"start,omitempty"`
	Stop         int           `json:"stop,omitempty"`
	End          int           `json:"end,omitempty"`
	Bit          int           `json:"bit,omitempty"`
	Member       string        `json:"member,omitempty"`
	Member1      string        `json:"member1,omitempty"`
	Member2      string        `json:"member2,omitempty"`
	Members      []string      `json:"members,omitempty"`
	Lon          float64       `json:"lon,omitempty"`
	Lat          float64       `json:"lat,omitempty"`
	Distance     float64       `json:"distance,omitempty"`
	Unit         string        `json:"unit,omitempty"`
	Options      []interface{} `json:"options,omitempty"`
	Keys         []string      `json:"keys,omitempty"`
	Cursor       int           `json:"cursor,omitempty"`
	Offset       int           `json:"offset,omitempty"`
	Field        string        `json:"field,omitempty"`
	Fields       []string      `json:"fields,omitempty"`
	Subcommand   string        `json:"subcommand,omitempty"`
	Pattern      string        `json:"pattern,omitempty"`
	Idx          int           `json:"idx, omitempty"`
	Min          string        `json:"min,omitempty"`
	Max          string        `json:"max,omitempty"`
	Limit        string        `json:"limit,omitempty"`
	Count        int           `json:"count,omitempty"`
	Match        string        `json:"match,omitempty"`
	Reset        bool          `json:"reset,omitempty"`
	IncludeTrash bool          `json:"includeTrash,omitempty"`
}

type SubscribeQuery struct {
	Scope string      `json:"scope"`
	State string      `json:"state"`
	User  string      `json:"user"`
	Body  interface{} `json:"body"`
}

type VolatileData = json.RawMessage

type UserData struct {
	ProfileIds []string               `json:"profileIds"`
	Content    map[string]interface{} `json:"content"`
}

type PolicyRestriction struct {
	Index       string   `json:"index"`
	Collections []string `json:"collections,omitempty"`
}

type Policy struct {
	RoleId       string               `json:"roleId"`
	RestrictedTo []*PolicyRestriction `json:"restrictedTo,omitempty"`
}

type Policies struct {
	Policies []*Policy `json:"policies"`
}

type GeoPoint struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name string  `json:"name"`
}

type MsHashField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type MSKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MSSortedSet struct {
	Score  float64 `json:"score"`
	Member string  `json:"member"`
}

type SearchFilters struct {
	Query        json.RawMessage `json:"query,omitempty"`
	Sort         json.RawMessage `json:"sort,omitempty"`
	Aggregations json.RawMessage `json:"aggregations,omitempty"`
	SearchAfter  json.RawMessage `json:"search_after,omitempty"`
}
