package types

import (
	"time"
)

const (
	Auto = iota
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
	Volatile          VolatileData
	Refresh           string
	IfExist           string
	DefaultIndex      string
	From              int
	Size              int
	Scroll            string
	ScrollId          string
	Headers           map[string]interface{}
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
		Headers:           make(map[string]interface{}),
	}
}

type QueryOptions struct {
	Queuable bool
}

type RoomOptions struct {
	Scope           string
	State           string
	User            string
	SubscribeToSelf bool
}

func DefaultRoomOptions() *RoomOptions {
	return &RoomOptions{
		Scope:           SCOPE_ALL,
		State:           STATE_DONE,
		User:            USER_NONE,
		SubscribeToSelf: true,
	}
}
