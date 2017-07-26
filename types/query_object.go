package types

import (
	"time"
)

type QueryObject struct {
	Query     []byte
	Options   QueryOptions
	ResChan   chan<- KuzzleResponse
	Timestamp time.Time
	RequestId string
}
