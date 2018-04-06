package types

import (
	"encoding/json"
)

type IKuzzle interface {
	Query(query *KuzzleRequest, options QueryOptions, responseChannel chan<- *KuzzleResponse)
	EmitEvent(int, interface{})
	SetJwt(string)
	RegisterSub(string, string, json.RawMessage, chan<- KuzzleNotification, chan<- interface{})
	UnregisterSub(string)
	AddListener(event int, notifChan chan<- interface{})
	AutoResubscribe() bool
}
