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

	KuzzleNotification struct {
		RequestId string       `json:"requestId"`
		Result    KuzzleResult `json:"result"`
		RoomId    string       `json:"room"`
		Error     MessageError `json:"error"`
	}

	KuzzleResult struct {
		Id      string          `json:"_id"`
		Meta    KuzzleMeta      `json:"_meta"`
		Source  json.RawMessage `json:"_source"`
		Version int             `json:"_version"`
	}

	KuzzleResponse struct {
		RequestId string          `json:"requestId"`
		Result    json.RawMessage `json:"result"`
		RoomId    string          `json:"room"`
		Error     MessageError    `json:"error"`
	}

	KuzzleSearchResult struct {
		Hits     []KuzzleResult `json:"hits"`
		Total    int            `json:"total"`
		ScrollId string         `json:"_scroll_id"`
	}

	KuzzleValidationFields map[string]struct {
		Type         string `json:"type"`
		Mandatory    bool   `json:"mandatory"`
		DefaultValue string `json:"defaultValue"`
	}

	KuzzleValidation struct {
		Strict bool                   `json:"strict"`
		Fields KuzzleValidationFields `json:"fields"`
	}

	KuzzleFieldMapping map[string]struct {
		Type   string `json:"type"`
		Fields json.RawMessage `json:"fields"`
	}

	KuzzleSpecifications map[string]map[string]struct {
		Strict bool                   `json:"strict"`
		Fields KuzzleValidationFields `json:"fields"`
	}

	KuzzleSpecificationsResult struct {
		Validation KuzzleValidation `json:"validation"`
		Index      string           `json:"index"`
		Collection string           `json:"collection"`
	}

	KuzzleSpecificationSearchResult struct {
		Hits []struct {
			Source KuzzleSpecificationsResult `json:"_source"`
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

	Document KuzzleResult

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
		Successful int `json:successful`
		Failed     int `json:"failed"`
	}

	CollectionsList struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	SecurityDocument struct {
		Id     string          `json:"_id"`
		Source json.RawMessage `json:"_source"`
		Meta   KuzzleMeta      `json:"_meta"`
	}

	User        SecurityDocument
	Profile     SecurityDocument
	Role        SecurityDocument
	Credentials map[string]interface{}

	UserRights struct {
		Controller string `json:"controller"`
		Action     string `json:"action"`
		Index      string `json:"index"`
		Collection string `json:"collection"`
		Value      string `json:"value"`
	}
)

func (user User) ProfileIDs() []string {
	type profileIds struct {
		ProfileIds []string `json:"profileIds"`
	}

	var content = profileIds{}
	json.Unmarshal(user.Source, &content)

	return content.ProfileIds
}

func (user User) Content(key string) interface{} {
	type Fields map[string]interface{}

	var content = Fields{}
	json.Unmarshal(user.Source, &content)

	if key == "" {
		return content
	}

	return content[key]
}

func (user User) ContentMap(keys ...string) map[string]interface{} {
	type Fields map[string]interface{}

	var content = Fields{}
	json.Unmarshal(user.Source, &content)

	if len(keys) == 0 {
		return content
	}

	values := make(map[string]interface{})

	for _, key := range keys {
		values[key] = content[key]
	}

	return values
}

func (profile Profile) Policies() []string {
	type Policies struct {
		Policies []struct {
			RoleId string `json:"roleId"`
		} `json:"policies"`
	}

	var policies = Policies{}
	json.Unmarshal(profile.Source, &policies)

	roleIDs := []string{}

	for _, role := range policies.Policies {
		roleIDs = append(roleIDs, role.RoleId)
	}

	return roleIDs
}

func (role Role) Controllers() map[string]struct {
	Actions map[string]bool `json:"actions"`
} {
	type Controllers struct {
		Controllers map[string]struct {
			Actions map[string]bool `json:"actions"`
		} `json:"controllers"`
	}

	var controllers = Controllers{}
	json.Unmarshal(role.Source, &controllers)

	return controllers.Controllers
}
