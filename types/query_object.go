package types

import "time"

type QueryObject struct {
  Query     []byte
  Options   *Options
  ResChan   chan<- KuzzleResponse
  Timestamp time.Time
  RequestId string
}
