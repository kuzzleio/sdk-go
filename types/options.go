package types

import (
	"time"
)

const (
	Auto = iota
	Manual
)

type Options interface {
	QueueTTL() time.Duration
	SetQueueTTL(time.Duration) *options
	QueueMaxSize() int
	SetQueueMaxSize(int) *options
	OfflineMode() int
	SetOfflineMode(int) *options
	AutoQueue() bool
	SetAutoQueue(bool) *options
	AutoReconnect() bool
	SetAutoReconnect(bool) *options
	AutoReplay() bool
	SetAutoReplay(bool) *options
	AutoResubscribe() bool
	SetAutoResubscribe(bool) *options
	ReconnectionDelay() time.Duration
	SetReconnectionDelay(time.Duration) *options
	ReplayInterval() time.Duration
	SetReplayInterval(time.Duration) *options
	Connect() int
	SetConnect(int) *options
	Refresh() string
	SetRefresh(string) *options
	DefaultIndex() string
	SetDefaultIndex(string) *options
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
}

func (o options) QueueTTL() time.Duration {
	return o.queueTTL
}

func (o *options) SetQueueTTL(queueTTL time.Duration) *options {
	o.queueTTL = queueTTL
	return o
}

func (o options) QueueMaxSize() int {
	return o.queueMaxSize
}

func (o *options) SetQueueMaxSize(queueMaxSize int) *options {
	o.queueMaxSize = queueMaxSize
	return o
}

func (o options) OfflineMode() int {
	return o.offlineMode
}

func (o *options) SetOfflineMode(offlineMode int) *options {
	o.offlineMode = offlineMode
	return o
}

func (o options) AutoQueue() bool {
	return o.autoQueue
}

func (o *options) SetAutoQueue(autoQueue bool) *options {
	o.autoQueue = autoQueue
	return o
}

func (o options) AutoReconnect() bool {
	return o.autoReconnect
}

func (o *options) SetAutoReconnect(autoReconnect bool) *options {
	o.autoReconnect = autoReconnect
	return o
}

func (o options) AutoReplay() bool {
	return o.autoReplay
}

func (o *options) SetAutoReplay(autoReplay bool) *options {
	o.autoReplay = autoReplay
	return o
}

func (o options) AutoResubscribe() bool {
	return o.autoResubscribe
}

func (o *options) SetAutoResubscribe(autoResubscribe bool) *options {
	o.autoResubscribe = autoResubscribe
	return o
}

func (o options) ReconnectionDelay() time.Duration {
	return o.reconnectionDelay
}

func (o *options) SetReconnectionDelay(reconnectionDelay time.Duration) *options {
	o.reconnectionDelay = reconnectionDelay
	return o
}

func (o options) ReplayInterval() time.Duration {
	return o.replayInterval
}

func (o *options) SetReplayInterval(replayInterval time.Duration) *options {
	o.replayInterval = replayInterval
	return o
}

func (o options) Connect() int {
	return o.connect
}

func (o *options) SetConnect(connect int) *options {
	o.connect = connect
	return o
}

func (o options) Refresh() string {
	return o.refresh
}

func (o *options) SetRefresh(refresh string) *options {
	o.refresh = refresh
	return o
}

func (o options) DefaultIndex() string {
	return o.defaultIndex
}

func (o *options) SetDefaultIndex(defaultIndex string) *options {
	o.defaultIndex = defaultIndex
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
	}
}
