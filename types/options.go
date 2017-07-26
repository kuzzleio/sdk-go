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
	SetQueueTTL(time.Duration)
	GetQueueMaxSize() int
	SetQueueMaxSize(int)
	GetOfflineMode() int
	SetOfflineMode(int)
	GetAutoQueue() bool
	SetAutoQueue(bool)
	GetAutoReconnect() bool
	SetAutoReconnect(bool)
	GetAutoReplay() bool
	SetAutoReplay(bool)
	GetAutoResubscribe() bool
	SetAutoResubscribe(bool)
	GetReconnectionDelay() time.Duration
	SetReconnectionDelay(time.Duration)
	GetReplayInterval() time.Duration
	SetReplayInterval(time.Duration)
	GetConnect() int
	SetConnect(int)
	GetRefresh() string
	SetRefresh(string)
	GetDefaultIndex() string
	SetDefaultIndex(string)
	GetHeaders() HeadersData
	SetHeaders(HeadersData)
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

func (o *options) SetQueueTTL(queueTTL time.Duration) {
	o.queueTTL = queueTTL
}

func (o options) GetQueueMaxSize() int {
	return o.queueMaxSize
}

func (o *options) SetQueueMaxSize(queueMaxSize int) {
	o.queueMaxSize = queueMaxSize
}

func (o options) GetOfflineMode() int {
	return o.offlineMode
}

func (o *options) SetOfflineMode(offlineMode int) {
	o.offlineMode = offlineMode
}

func (o options) GetAutoQueue() bool {
	return o.autoQueue
}

func (o *options) SetAutoQueue(autoQueue bool) {
	o.autoQueue = autoQueue
}

func (o options) GetAutoReconnect() bool {
	return o.autoReconnect
}

func (o *options) SetAutoReconnect(autoReconnect bool) {
	o.autoReconnect = autoReconnect
}

func (o options) GetAutoReplay() bool {
	return o.autoReplay
}

func (o *options) SetAutoReplay(autoReplay bool) {
	o.autoReplay = autoReplay
}

func (o options) GetAutoResubscribe() bool {
	return o.autoResubscribe
}

func (o *options) SetAutoResubscribe(autoResubscribe bool) {
	o.autoResubscribe = autoResubscribe
}

func (o options) GetReconnectionDelay() time.Duration {
	return o.reconnectionDelay
}

func (o *options) SetReconnectionDelay(reconnectionDelay time.Duration) {
	o.reconnectionDelay = reconnectionDelay
}

func (o options) GetReplayInterval() time.Duration {
	return o.replayInterval
}

func (o *options) SetReplayInterval(replayInterval time.Duration) {
	o.replayInterval = replayInterval
}

func (o options) GetConnect() int {
	return o.connect
}

func (o *options) SetConnect(connect int) {
	o.connect = connect
}

func (o options) GetRefresh() string {
	return o.refresh
}

func (o *options) SetRefresh(refresh string) {
	o.refresh = refresh
}

func (o options) GetDefaultIndex() string {
	return o.defaultIndex
}

func (o *options) SetDefaultIndex(defaultIndex string) {
	o.defaultIndex = defaultIndex
}

func (o options) GetHeaders() HeadersData {
	return o.headers
}

func (o *options) SetHeaders(headers HeadersData) {
	o.headers = headers
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
