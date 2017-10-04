package types

import (
	"time"
)

type QueryObject struct {
	Query     []byte
	Options   QueryOptions
	ResChan   chan<- *KuzzleResponse
	NotifChan chan<- *KuzzleNotification
	Timestamp time.Time
	RequestId string
}
