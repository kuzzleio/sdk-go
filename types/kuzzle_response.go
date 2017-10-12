package types

import (
	"fmt"
	"encoding/json"
)

type (
	KuzzleError struct {
		Message string `json:"message"`
		Stack   string `json:"stack"`
		Status  uint   `json:"status"`
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
		Id         string           `json:"_id"`
		Meta       *KuzzleMeta      `json:"_meta"`
		Content    json.RawMessage  `json:"_source"`
		Version    int              `json:"_version"`
		Collection string           `json:"collection"`
	}

	KuzzleNotification struct {
		RequestId string        `json:"requestId"`
		Result    *KuzzleResult `json:"result"`
		RoomId    string        `json:"room"`
		Error     *KuzzleError `json:"error"`
	}

	KuzzleResponse struct {
		RequestId string           `json:"requestId"`
		Result    json.RawMessage  `json:"result"`
		RoomId    string           `json:"room"`
		Channel   string           `json:"channel"`
		Status		uint 		  			 `json:"status"`
		Error     *KuzzleError     `json:"error"`
	}

	KuzzleValidationFields map[string]*struct {
		Type         string `json:"type"`
		Mandatory    bool   `json:"mandatory"`
		DefaultValue string `json:"defaultValue"`
	}

	KuzzleValidation struct {
		Strict bool                    `json:"strict"`
		Fields *KuzzleValidationFields `json:"fields"`
	}

	KuzzleFieldMapping map[string]*struct {
		Type   string           `json:"type"`
		Fields json.RawMessage  `json:"fields"`
	}

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
		Id         string           `json:"_id"`
		Source     json.RawMessage  `json:"_source"`
		Meta       *KuzzleMeta      `json:"_meta"`
		Strategies []string         `json:"strategies"`
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
		Id         string           `json:"_id"`
		Source     json.RawMessage  `json:"_source"`
		Meta       *KuzzleMeta      `json:"_meta"`
		Strategies []string         `json:"strategies"`
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

func (e *KuzzleError) Error() string {
	msg := e.Message

  if len(e.Stack) > 0 {
    msg = fmt.Sprintf("%s\n%s", msg, e.Stack)
  }

  if e.Status > 0 {
  	msg = fmt.Sprintf("[%d] %s", msg)
  }

  return msg
}

func NewError(msg string, status ...int)  *KuzzleError {
	err := &KuzzleError{Message: msg}

	if len(status) == 1 {
		err.Status = status[0]
	} 

	return err
}
