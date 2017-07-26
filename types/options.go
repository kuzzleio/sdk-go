package types

import (
	"time"
)

const (
	Auto = iota
	Manual
)

type Options interface {
	GetQueueTTL() time.Duration
	SetQueueTTL(time.Duration) *options
	GetQueueMaxSize() int
	SetQueueMaxSize(int) *options
	GetOfflineMode() int
	SetOfflineMode(int) *options
	GetAutoQueue() bool
	SetAutoQueue(bool) *options
	GetAutoReconnect() bool
	SetAutoReconnect(bool) *options
	GetAutoReplay() bool
	SetAutoReplay(bool) *options
	GetAutoResubscribe() bool
	SetAutoResubscribe(bool) *options
	GetReconnectionDelay() time.Duration
	SetReconnectionDelay(time.Duration) *options
	GetReplayInterval() time.Duration
	SetReplayInterval(time.Duration) *options
	GetConnect() int
	SetConnect(int) *options
	GetRefresh() string
	SetRefresh(string) *options
	GetDefaultIndex() string
	SetDefaultIndex(string) *options
	GetHeaders() HeadersData
	SetHeaders(HeadersData) *options
}

type options struct {
	queueTTL          time.Duration
	queueMaxSize      int
	offlineMode       int
	autoQueue         bool
	autoReconnect     bool
	autoReplay        bool
	autoResubscribe   bool
	reconnectionDelay time.Duration
	replayInterval    time.Duration
	connect           int
	refresh           string
	defaultIndex      string
	headers           HeadersData
}

func (o options) GetQueueTTL() time.Duration {
	return o.queueTTL
}

func (o *options) SetQueueTTL(queueTTL time.Duration) *options {
	o.queueTTL = queueTTL
	return o
}

func (o options) GetQueueMaxSize() int {
	return o.queueMaxSize
}

func (o *options) SetQueueMaxSize(queueMaxSize int) *options {
	o.queueMaxSize = queueMaxSize
	return o
}

func (o options) GetOfflineMode() int {
	return o.offlineMode
}

func (o *options) SetOfflineMode(offlineMode int) *options {
	o.offlineMode = offlineMode
	return o
}

func (o options) GetAutoQueue() bool {
	return o.autoQueue
}

func (o *options) SetAutoQueue(autoQueue bool) *options {
	o.autoQueue = autoQueue
	return o
}

func (o options) GetAutoReconnect() bool {
	return o.autoReconnect
}

func (o *options) SetAutoReconnect(autoReconnect bool) *options {
	o.autoReconnect = autoReconnect
	return o
}

func (o options) GetAutoReplay() bool {
	return o.autoReplay
}

func (o *options) SetAutoReplay(autoReplay bool) *options {
	o.autoReplay = autoReplay
	return o
}

func (o options) GetAutoResubscribe() bool {
	return o.autoResubscribe
}

func (o *options) SetAutoResubscribe(autoResubscribe bool) *options {
	o.autoResubscribe = autoResubscribe
	return o
}

func (o options) GetReconnectionDelay() time.Duration {
	return o.reconnectionDelay
}

func (o *options) SetReconnectionDelay(reconnectionDelay time.Duration) *options {
	o.reconnectionDelay = reconnectionDelay
	return o
}

func (o options) GetReplayInterval() time.Duration {
	return o.replayInterval
}

func (o *options) SetReplayInterval(replayInterval time.Duration) *options {
	o.replayInterval = replayInterval
	return o
}

func (o options) GetConnect() int {
	return o.connect
}

func (o *options) SetConnect(connect int) *options {
	o.connect = connect
	return o
}

func (o options) GetRefresh() string {
	return o.refresh
}

func (o *options) SetRefresh(refresh string) *options {
	o.refresh = refresh
	return o
}

func (o options) GetDefaultIndex() string {
	return o.defaultIndex
}

func (o *options) SetDefaultIndex(defaultIndex string) *options {
	o.defaultIndex = defaultIndex
	return o
}

func (o options) GetHeaders() HeadersData {
	return o.headers
}

func (o *options) SetHeaders(headers HeadersData) *options {
	o.headers = headers
	return o
}

func NewOptions() *options {
	return &options{
		queueTTL:          120000,
		queueMaxSize:      500,
		offlineMode:       Manual,
		autoQueue:         false,
		autoReconnect:     true,
		autoReplay:        false,
		autoResubscribe:   true,
		reconnectionDelay: 1000,
		replayInterval:    10,
		connect:           Auto,
		headers:           make(HeadersData),
	}
}

type HeadersData map[string]interface{}
