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
  }
}

type RoomOptions struct {
  Scope           string
  State           string
  User            string
  SubscribeToSelf bool
}

func DefaultRoomOptions() *RoomOptions {
  return &RoomOptions{
    Scope: SCOPE_ALL,
    State: STATE_DONE,
    User: USER_NONE,
    SubscribeToSelf: true,
  }
}
