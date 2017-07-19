package types

import (
  "encoding/json"
)

type MessageError struct {
  Message string `json:"message"`
  Status  int `json:"status"`
}

type KuzzleMeta struct {
  Author    string `json:"author"`
  CreatedAt int `json:"createdAt"`
  UpdatedAt int `json:"updatedAt"`
  Updater   string `json:"updater"`
  Active    bool `json:"active"`
  DeletedAt int `json:"deletedAt"`
}

type KuzzleNotification struct {
  RequestId string `json:"requestId"`
  Result    KuzzleResult `json:"result"`
  RoomId    string `json:"room"`
  Error     MessageError `json:"error"`
}

type KuzzleResult struct {
  Id     string `json:"_id"`
  Meta   KuzzleMeta `json:"_meta"`
  Source json.RawMessage `json:"_source"`
  Version int `json:"_version"`
}

type KuzzleResponse struct {
  RequestId string `json:"requestId"`
  Result    json.RawMessage `json:"result"`
  RoomId    string `json:"room"`
  Error     MessageError `json:"error"`
}

type KuzzleSearchResult struct {
  Hits     []KuzzleResult `json:"hits"`
  Total    int `json:"total"`
  ScrollId string `json:"_scroll_id"`
}

type AckResponse struct {
  Acknowledged       bool
  ShardsAcknowledged bool
}

type Document KuzzleResult
