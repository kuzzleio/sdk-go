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

type (
	// KuzzleError is a custom Error type for Kuzzle
	KuzzleError struct {
		Message string `json:"message"`
		Stack   string `json:"stack"`
		Status  int    `json:"status"`
	}

	// Meta contains metadata
	Meta struct {
		Author    string `json:"author"`
		CreatedAt int    `json:"createdAt"`
		UpdatedAt int    `json:"updatedAt"`
		Updater   string `json:"updater"`
		Active    bool   `json:"active"`
		DeletedAt int    `json:"deletedAt"`
	}

	SubscribeResult struct {
		Room    string `json:"roomId"`
		Channel string `json:"channel"`
	}

	// NotificationContent contains
	NotificationContent struct {
		Id      string          `json:"_id"`
		Meta    *Meta           `json:"_meta"`
		Content json.RawMessage `json:"_source"`
		Count   int             `json:"count"`
	}

	// NotificationResult is a notification from Kuzzle
	NotificationResult struct {
		RequestId  string               `json:"requestId"`
		Result     *NotificationContent `json:"result"`
		Volatile   json.RawMessage      `json:"volatile"`
		Index      string               `json:"index"`
		Collection string               `json:"collection"`
		Controller string               `json:"controller"`
		Action     string               `json:"action"`
		Protocol   string               `json:"protocol"`
		Scope      string               `json:"scope"`
		State      string               `json:"state"`
		User       string               `json:"user"`
		Type       string               `json:"type"`
		RoomId     string               `json:"room"`
		Channel    string               `json:"channel"`
		Timestamp  int                  `json:"timestamp"`
		Status     int                  `json:"status"`
		Error      KuzzleError          `json:"error"`
	}

	// KuzzleResponse is a response to a KuzzleRequest
	KuzzleResponse struct {
		RequestId  string          `json:"requestId"`
		Result     json.RawMessage `json:"result"`
		Volatile   VolatileData    `json:"volatile"`
		Index      string          `json:"index"`
		Collection string          `json:"collection"`
		Controller string          `json:"controller"`
		Action     string          `json:"action"`
		RoomId     string          `json:"room"`
		Channel    string          `json:"channel"`
		Status     int             `json:"status"`
		Error      KuzzleError     `json:"error"`
	}

	SpecificationField struct {
		Type        string `json:"type,omitempty"`
		Depth       int    `json:"depth,omitempty"`
		Mandatory   bool   `json:"mandatory,omitempty"`
		Description string `json:"description,omitempty"`
		Multivalued struct {
			Value    bool `json:"value,omitempty"`
			MinCount int  `json:"minCount,omitempty"`
			MaxCount int  `json:"maxCount,omitempty"`
		} `json:"multivalued,omitempty"`
		DefaultValue interface{} `json:"defaultValue,omitempty"`
		TypeOptions  struct {
			Range struct {
				Min interface{} `json:"min,omitempty"`
				Max interface{} `json:"max,omitempty"`
			} `json:"range,omitempty"`
			Length struct {
				Min int         `json:"min,omitempty"`
				Max interface{} `json:"max,omitempty"`
			} `json:"length"`
			NotEmpty   bool     `json:"notEmpty,omitempty"`
			Formats    []string `json:"formats,omitempty"`
			Strict     bool     `json:"strict,omitempty"`
			Values     []string `json:"values,omitempty"`
			ShapeTypes []string `json:"shapeTypes,omitempty"`
		} `json:"typeOptions,omitempty"`
	}

	SpecificationFields map[string]SpecificationField

	Specification struct {
		Strict     bool                `json:"strict,omitempty"`
		Fields     SpecificationFields `json:"fields,omitempty"`
		Validators json.RawMessage     `json:"validators,omitempty"`
	}

	MappingField struct {
		Analyzer                 string      `json:"analyzer,omitempty"`
		Normalizer               interface{} `json:"normalizer,omitempty"`
		DocValues                bool        `json:"doc_values,omitempty"`
		Boost                    float64     `json:"boost,omitempty"`
		Coerce                   bool        `json:"coerce,omitempty"`
		Enabled                  bool        `json:"enabled,omitempty"`
		FieldData                bool        `json:"fielddata,omitempty"`
		FieldDataFrequencyFilter struct {
			Min            float64 `json:"min,omitempty"`
			Max            float64 `json:"max,omitempty"`
			MinSegmentSize int     `json:"min_segment_size,omitempty"`
		} `json:"fielddata_frequency_filter,omitempty"`
		Format               string                  `json:"format,omitempty"`
		IgnoreAbove          int                     `json:"ignore_above,omitempty"`
		IgnoreMalformed      bool                    `json:"ignore_malformed,omitempty"`
		IncludeInAll         bool                    `json:"include_in_all,omitempty"`
		Index                bool                    `json:"index,omitempty"`
		IndexOptions         bool                    `json:"index_options,omitempty"`
		Fields               map[string]MappingField `json:"fields,omitempty"`
		Norms                bool                    `json:"norms,omitempty"`
		NullValue            bool                    `json:"null_value,omitempty"`
		PositionIncrementGap bool                    `json:"position_increment_gap,omitempty"`
		Type                 string                  `json:"type,omitempty"`
		All                  *struct {
			Enabled bool   `json:"enabled,omitempty"`
			Format  string `json:"format, omitempty"`
		} `json:"_all,omitempty"`
		Properties               MappingFields          `json:"properties,omitempty"`
		SearchAnalyzer           string                 `json:"search_analyzer,omitempty"`
		Similarity               string                 `json:"similarity,omitempty"`
		Store                    bool                   `json:"store,omitempty"`
		TermVector               string                 `json:"term_vector,omitempty"`
		Tree                     string                 `json:"tree,omitempty"`
		Precision                string                 `json:"precision,omitempty"`
		TreeLevels               int                    `json:"tree_levels,omitempty"`
		Strategy                 string                 `json:"strategy,omitempty"`
		DistanceErrorPct         float64                `json:"distance_error_pct,omitempty"`
		Orientation              string                 `json:"orientation,omitempty"`
		PointsOnly               bool                   `json:"points_only,omitempty"`
		EagerGlobalOrdinals      bool                   `json:"eager_global_ordinals,omitempty"`
		Dynamic                  interface{}            `json:"dynamic,omitempty"`
		SearchQuoteAnalyzer      string                 `json:"search_quote_analyzer,omitempty"`
		EnablePositionIncrements bool                   `json:"enable_position_increments,omitempty"`
		Relations                map[string]interface{} `json:"relations,omitempty"`
	}

	MappingFields map[string]MappingField

	SpecificationEntry struct {
		Validation *Specification `json:"validation"`
		Index      string         `json:"index"`
		Collection string         `json:"collection"`
	}

	SpecificationSearchResultHit struct {
		Source SpecificationEntry `json:"_source"`
	}

	SpecificationSearchResult struct {
		Hits     []SpecificationSearchResultHit `json:"hits"`
		Total    int                            `json:"total"`
		ScrollId string                         `json:"scrollId"`
	}

	ValidResponse struct {
		Valid bool `json:"valid"`
	}

	RealtimeResponse struct {
		Published bool `json:"published"`
	}

	ShardResponse struct {
		Found   bool    `json:"found"`
		Index   string  `json:"_index"`
		Type    string  `json:"_type"`
		Id      string  `json:"_id"`
		Version int     `json:"_version"`
		Result  string  `json:"result"`
		Shards  *Shards `json:"_shards"`
	}

	Statistics struct {
		CompletedRequests json.RawMessage `json:"completedRequests"`
		Connections       json.RawMessage `json:"connections"`
		FailedRequests    json.RawMessage `json:"failedRequests"`
		OngoingRequests   json.RawMessage `json:"ongoingRequests"`
		Timestamp         int             `json:"timestamp"`
	}

	LoginAttempt struct {
		Success bool  `json:"success"`
		Error   error `json:"error"`
	}

	Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	}

	CollectionsList struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	Controller struct {
		Actions map[string]bool `json:"actions"`
	}

	Controllers struct {
		Controllers map[string]*Controller `json:"controllers"`
	}

	CredentialStrategyFields []string
	CredentialFields         map[string]CredentialStrategyFields

	Credentials = json.RawMessage

	UserRights struct {
		Controller string `json:"controller"`
		Action     string `json:"action"`
		Index      string `json:"index"`
		Collection string `json:"collection"`
		Value      string `json:"value"`
	}

	Georadius struct {
		Name string
		Lon  float64
		Lat  float64
		Dist float64
	}

	MSScanResponse struct {
		Cursor int      `json:"cursor"`
		Values []string `json:"values"`
	}
)

func (e KuzzleError) Error() string {
	msg := e.Message

	if len(e.Stack) > 0 {
		msg = fmt.Sprintf("%s\n%s", msg, e.Stack)
	}

	if e.Status > 0 {
		msg = fmt.Sprintf("[%d] %s", e.Status, msg)
	}

	return msg
}

// NewError instanciates a new KuzzleError
func NewError(msg string, status ...int) KuzzleError {
	err := KuzzleError{Message: msg}

	if len(status) == 1 {
		err.Status = status[0]
	}

	return err
}
