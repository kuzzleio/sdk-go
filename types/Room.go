package types

import "encoding/json"

type Room struct {
  RoomId string `json:"room"`
  Channel string `json:"channel"`
  Result json.RawMessage `json:"result"`
}
