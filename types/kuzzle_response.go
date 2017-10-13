package types

import (
	"encoding/json"
)

type (
	MessageError struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

	KuzzleMeta struct {
		Author    string `json:"author"`
		CreatedAt int    `json:"createdAt"`
		UpdatedAt int    `json:"updatedAt"`
		Updater   string `json:"updater"`
		Active    bool   `json:"active"`
		DeletedAt int    `json:"deletedAt"`
	}

	KuzzleResult struct {
		Id         string          `json:"_id"`
		Meta       *KuzzleMeta     `json:"_meta"`
		Content    json.RawMessage `json:"_source"`
		Version    int             `json:"_version"`
		Collection string          `json:"collection"`
	}

	KuzzleNotification struct {
		RequestId string        `json:"requestId"`
		Result    *KuzzleResult `json:"result"`
		RoomId    string        `json:"room"`
		Error     *MessageError `json:"error"`
	}

	KuzzleResponse struct {
		RequestId string          `json:"requestId"`
		Result    json.RawMessage `json:"result"`
		RoomId    string          `json:"room"`
		Channel   string          `json:"channel"`
		Error     *MessageError   `json:"error"`
	}

	KuzzleValidationFields map[string]*struct {
		Type        string   `json:"type,omitempty"`
		Path        []string `json:"path,omitempty"`
		Depth       int      `json:"depth,omitempty"`
		Mandatory   bool     `json:"mandatory,omitempty"`
		Description string   `json:"description,omitempty"`
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

	KuzzleValidation struct {
		Strict     bool                    `json:"strict,omitempty"`
		Fields     *KuzzleValidationFields `json:"fields,omitempty"`
		Validators json.RawMessage         `json:"validators,omitempty"`
	}

	KuzzleFieldMapping struct {
		Type       string                 `json:"type,omitempty"`
		Properties map[string]interface{} `json:"properties,omitempty"`
	}

	KuzzleFieldsMapping map[string]*KuzzleFieldMapping

	KuzzleSpecifications map[string]map[string]*KuzzleValidation

	KuzzleSpecificationsResult struct {
		Validation *KuzzleValidation `json:"validation"`
		Index      string            `json:"index"`
		Collection string            `json:"collection"`
	}

	KuzzleSpecificationSearchResult struct {
		Hits []*struct {
			Source *KuzzleSpecificationsResult `json:"_source"`
		} `json:"hits"`
		Total    int    `json:"total"`
		ScrollId string `json:"scrollId"`
	}

	ValidResponse struct {
		Valid bool `json:"valid"`
	}

	RealtimeResponse struct {
		Published bool `json:"published"`
	}

	AckResponse struct {
		Acknowledged       bool `json:"acknowledged"`
		ShardsAcknowledged bool `json:"shardsAcknowledged"`
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
		CompletedRequests map[string]int `json:"completedRequests"`
		Connections       map[string]int `json:"connections"`
		FailedRequests    map[string]int `json:"failedRequests"`
		OngoingRequests   map[string]int `json:"ongoingRequests"`
		Timestamp         int            `json:"timestamp"`
	}

	Rights struct {
		Controller string `json:"controller"`
		Action     string `json:"action"`
		Index      string `json:"index"`
		Collection string `json:"collection"`
		Value      string `json:"value"`
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
		Controllers map[string]Controller `json:"controllers"`
	}

	SecurityDocument struct {
		Id         string          `json:"_id"`
		Source     json.RawMessage `json:"_source"`
		Meta       *KuzzleMeta     `json:"_meta"`
		Strategies []string        `json:"strategies"`
	}

	Profile SecurityDocument
	Role    SecurityDocument

	CredentialStrategyFields []string
	CredentialFields         map[string]CredentialStrategyFields

	Credentials map[string]interface{}

	UserRights struct {
		Controller string `json:"controller"`
		Action     string `json:"action"`
		Index      string `json:"index"`
		Collection string `json:"collection"`
		Value      string `json:"value"`
	}

	User struct {
		Id         string          `json:"_id"`
		Source     json.RawMessage `json:"_source"`
		Meta       *KuzzleMeta     `json:"_meta"`
		Strategies []string        `json:"strategies"`
	}

	GeoradiusPointWithCoord struct {
		Name string
		Lon  float64
		Lat  float64
	}

	GeoradiusPointWithDist struct {
		Name string
		Dist float64
	}

	GeoradiusPointWithCoordAndDist struct {
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
