package types

import (
  "encoding/json"
)

type (
  MessageError struct {
    Message string `json:"message"`
    Status  int `json:"status"`
  }

  KuzzleMeta struct {
    Author    string `json:"author"`
    CreatedAt int `json:"createdAt"`
    UpdatedAt int `json:"updatedAt"`
    Updater   string `json:"updater"`
    Active    bool `json:"active"`
    DeletedAt int `json:"deletedAt"`
  }

  KuzzleNotification struct {
    RequestId string `json:"requestId"`
    Result    KuzzleResult `json:"result"`
    RoomId    string `json:"room"`
    Error     MessageError `json:"error"`
  }

  KuzzleResult struct {
  Id     string `json:"_id"`
  Meta   KuzzleMeta `json:"_meta"`
  Source json.RawMessage `json:"_source"`
  Version int `json:"_version"`
}

  KuzzleResponse struct {
    RequestId string `json:"requestId"`
    Result    json.RawMessage `json:"result"`
    RoomId    string `json:"room"`
    Error     MessageError `json:"error"`
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

  KuzzleSpecifications map[string]map[string]struct {
    Strict bool                   `json:"strict"`
    Fields KuzzleValidationFields `json:"fields"`
  }

  KuzzleSpecificationsResult struct {
    Validation KuzzleValidation `json:"validation"`
    Index      string           `json:"index"`
    Collection string           `json:"collection"`
  }

  ValidResponse struct {
    Valid bool `json:"valid"`
  }

  RealtimeResponse struct {
    Published bool `json:"published"`
  }

  AckResponse struct {
    Acknowledged       bool
    ShardsAcknowledged bool
  }

  Document KuzzleResult

  Statistics struct {
    CompletedRequests map[string]int `json:"completedRequests"`
    Connections       map[string]int `json:"connections"`
    FailedRequests    map[string]int `json:"failedRequests"`
    OngoingRequests   map[string]int `json:"ongoingRequests"`
    Timestamp         int `json:"timestamp"`
  }

  UserInterface interface {
    ProfileIds()
    Content(key string)
    ContentMap(keys ...string)
  }

  User struct {
    Id     string          `json:"_id"`
    Source json.RawMessage `json:"_source"`
    Meta   KuzzleMeta      `json:"_meta"`
  }
)

func (user User) ProfileIDs() ([]string) {
  type profileIds struct {
    ProfileIds []string `json:"profileIds"`
  }

  var content = profileIds{}
  json.Unmarshal(user.Source, &content)

  return content.ProfileIds
}

func (user User) Content(key string) (interface{}) {
  type Fields map[string]interface{}

  var content = Fields{}
  json.Unmarshal(user.Source, &content)

  if key == "" {
    return content
  }

  return content[key]
}

func (user User) ContentMap(keys ...string) (map[string]interface{}) {
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
