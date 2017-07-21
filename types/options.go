package types

import (
  "time"
)

const (
  Auto   = iota
  Manual
)

type Options struct {
  Queuable          bool
  QueueTTL          time.Duration
  QueueMaxSize      int
  OfflineMode       int
  AutoQueue         bool
  AutoReconnect     bool
  AutoReplay        bool
  AutoResubscribe   bool
  ReconnectionDelay time.Duration
  ReplayInterval    time.Duration
  Connect           int
  Volatile          interface{}
  Refresh           string
  IfExist           string
  DefaultIndex      string
  From              int
  Size              int
  Scroll            string
  ScrollId          string
}

func DefaultOptions() *Options {
  return &Options{
    QueueTTL:          120000,
    QueueMaxSize:      500,
    OfflineMode:       Manual,
    AutoQueue:         false,
    AutoReconnect:     true,
    AutoReplay:        false,
    AutoResubscribe:   true,
    ReconnectionDelay: 1000,
    ReplayInterval:    10,
    Connect:           Auto,
    From:              0,
    Size:              10,
    Scroll:            "1m",
    ScrollId:          "",
  }
}

type QueryOptions struct {
  Queuable bool
}
