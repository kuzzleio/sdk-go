// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"net/http"
	"time"
)

const (
	Auto = iota
	Manual
)

// Options provides a public access to options private struct
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
	Port() int
	SetPort(int) *options
	SslConnection() bool
	SetSslConnection(bool) *options
	Headers() *http.Header
	SetHeaders(*http.Header) *options
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
	port              int
	sslConnection     bool
	headers           *http.Header
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

func (o *options) Port() int {
	return o.port
}

func (o *options) SetPort(v int) *options {
	o.port = v
	return o
}

func (o *options) SslConnection() bool {
	return o.sslConnection
}

func (o *options) SetSslConnection(v bool) *options {
	o.sslConnection = v
	return o
}

func (o *options) Headers() *http.Header {
	return o.headers
}

func (o *options) SetHeaders(h *http.Header) *options {
	o.headers = h
	return o
}

// NewOptions instanciates new Options with default values
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
		port:              7512,
		sslConnection:     false,
		headers:           nil,
	}
}
