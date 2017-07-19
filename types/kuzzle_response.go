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
  }

  KuzzleResponse struct {
    RequestId string `json:"requestId"`
    Result    json.RawMessage `json:"result"`
    RoomId    string `json:"room"`
    Error     MessageError `json:"error"`
  }

  KuzzleSearchResult struct {
    Hits  []KuzzleResult `json:"hits"`
    Total int `json:"total"`
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
)
