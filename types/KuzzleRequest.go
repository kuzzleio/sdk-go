package types

type KuzzleRequest struct {
  RequestId string `json:"requestId"`
  Controller string `json:"controller"`
  Action string `json:"action"`
  Index string `json:"index"`
  Collection string `json:"collection"`
  Body interface{} `json:"body"`
}
