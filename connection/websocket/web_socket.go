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

package websocket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
)

const (
	MAX_EMIT_TIMEOUT  = 10
	MAX_CONNECT_RETRY = 10
)

type subscription struct {
	roomID              string
	channel             string
	filters             json.RawMessage
	notificationChannel chan<- types.KuzzleNotification
	onReconnectChannel  chan<- interface{}
	subscribeToSelf     bool
}

type webSocket struct {
	ws      *websocket.Conn
	mu      *sync.Mutex
	queuing bool
	state   int

	listenChan         chan []byte
	channelsResult     sync.Map
	subscriptions      sync.Map
	lastUrl            string
	wasConnected       bool
	eventListeners     map[int]map[chan<- interface{}]bool
	eventListenersOnce map[int]map[chan<- interface{}]bool

	retrying              bool
	nbRetried             int
	stopRetryingToConnect bool
	requestHistory        map[string]time.Time

	autoQueue          bool
	autoReconnect      bool
	autoReplay         bool
	autoResubscribe    bool
	host               string
	offlineQueue       []*types.QueryObject
	offlineQueueLoader connection.OfflineQueueLoader
	port               int
	queueFilter        connection.QueueFilter
	queueMaxSize       int
	queueTTL           time.Duration
	reconnectionDelay  time.Duration
	replayInterval     time.Duration
	ssl                bool
}

var defaultQueueFilter connection.QueueFilter

// NewWebSocket instanciates a new webSocket connection object
func NewWebSocket(host string, options types.Options) connection.Connection {
	defaultQueueFilter = func([]byte) bool {
		return true
	}

	var opts types.Options

	if options == nil {
		opts = types.NewOptions()
	} else {
		opts = options
	}

	ws := &webSocket{
		mu:                    &sync.Mutex{},
		queueTTL:              opts.QueueTTL(),
		offlineQueue:          []*types.QueryObject{},
		queueMaxSize:          opts.QueueMaxSize(),
		channelsResult:        sync.Map{},
		subscriptions:         sync.Map{},
		eventListeners:        make(map[int]map[chan<- interface{}]bool),
		eventListenersOnce:    make(map[int]map[chan<- interface{}]bool),
		requestHistory:        make(map[string]time.Time),
		autoQueue:             opts.AutoQueue(),
		autoReconnect:         opts.AutoReconnect(),
		autoReplay:            opts.AutoReplay(),
		autoResubscribe:       opts.AutoResubscribe(),
		reconnectionDelay:     opts.ReconnectionDelay(),
		replayInterval:        opts.ReplayInterval(),
		state:                 state.Ready,
		retrying:              false,
		nbRetried:             0,
		stopRetryingToConnect: false,
		queueFilter:           defaultQueueFilter,
		port:                  opts.Port(),
		ssl:                   opts.SslConnection(),
	}
	ws.host = host

	if opts.OfflineMode() == types.Auto {
		ws.autoReconnect = true
		ws.autoQueue = true
		ws.autoReplay = true
		ws.autoResubscribe = true
	}
	ws.state = state.Offline

	return ws
}

//Connect connects to a kuzzle instance
func (ws *webSocket) Connect() (bool, error) {
	if ws.state != state.Offline {
		return false, nil
	}

	ws.state = state.Connecting

	if ws.autoQueue {
		ws.queuing = true
	}

	addr := fmt.Sprintf("%s:%d", ws.host, ws.port)

	if ws.lastUrl != addr {
		ws.wasConnected = false
		ws.lastUrl = addr
	}

	var scheme string

	if ws.ssl {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	u := url.URL{Scheme: scheme, Host: addr}
	socket, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		ws.state = state.Offline
		ws.EmitEvent(event.NetworkError, err)

		if ws.autoReconnect && !ws.retrying && !ws.stopRetryingToConnect && ws.nbRetried < MAX_CONNECT_RETRY {
			ws.retrying = true
			time.Sleep(ws.reconnectionDelay)
			ws.retrying = false
			ws.nbRetried++
			ws.Connect()
		} else {
			ws.EmitEvent(event.Disconnected, nil)
		}

		return false, err
	}

	ws.ws = socket
	ws.state = state.Connected
	ws.stopRetryingToConnect = false
	ws.queuing = false

	if ws.wasConnected {
		ws.EmitEvent(event.Reconnected, nil)
	} else {
		ws.wasConnected = true
		ws.EmitEvent(event.Connected, nil)
	}

	ws.listenChan = make(chan []byte)

	go ws.listen()

	go func() {
		for {
			_, message, err := ws.ws.ReadMessage()
			ws.listenChan <- message
			// TODO: either send a Disconnected event on a proper disconnection,
			// or a NetworkError one if the socket has been unexpectedly closed
			if err != nil {
				close(ws.listenChan)
				ws.ws.Close()
				ws.state = state.Offline
				if ws.autoQueue {
					ws.queuing = true
				}
				ws.EmitEvent(event.Disconnected, nil)
				return
			}
		}
	}()

	ws.PlayQueue()

	return ws.wasConnected, err
}

func (ws *webSocket) Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error {
	queuable := options == nil || options.Queuable()
	queuable = queuable && ws.queueFilter(query)

	if ws.queuing && queuable {
		ws.cleanQueue()
		qo := &types.QueryObject{
			Timestamp: time.Now(),
			ResChan:   responseChannel,
			Query:     query,
			RequestId: requestId,
			Options:   options,
		}
		ws.offlineQueue = append(ws.offlineQueue, qo)
		ws.EmitEvent(event.OfflineQueuePush, qo)
		return nil
	}

	if ws.state == state.Connected {
		return ws.emitRequest(&types.QueryObject{
			Query:     query,
			ResChan:   responseChannel,
			RequestId: requestId,
		})
	}

	ws.discardRequest(responseChannel, query)
	return nil
}

func (ws *webSocket) discardRequest(responseChannel chan<- *types.KuzzleResponse, query []byte) {
	if responseChannel != nil {
		responseChannel <- &types.KuzzleResponse{Status: 400, Error: types.NewError("Unable to execute request: not connected to a Kuzzle server.\nDiscarded request: "+string(query), 400)}
	}
}

// Clean up the queue, ensuring the queryTTL and queryMaxSize properties are respected
func (ws *webSocket) cleanQueue() {
	now := time.Now()
	now = now.Add(-ws.queueTTL * time.Millisecond)

	// Clean queue of timed out query
	if ws.queueTTL > 0 {
		var query *types.QueryObject
		for _, query = range ws.offlineQueue {
			if query.Timestamp.Before(now) {
				ws.offlineQueue = ws.offlineQueue[1:]
			} else {
				break
			}
		}
	}

	if ws.queueMaxSize > 0 && len(ws.offlineQueue) > ws.queueMaxSize {
		for len(ws.offlineQueue) > ws.queueMaxSize {
			eventListener := ws.eventListeners[event.OfflineQueuePop]
			for c := range eventListener {
				c <- ws.offlineQueue[0]
			}

			eventListener = ws.eventListenersOnce[event.OfflineQueuePop]
			for c := range eventListener {
				c <- ws.offlineQueue[0]
				delete(ws.eventListenersOnce[event.OfflineQueuePop], c)
			}

			ws.offlineQueue = ws.offlineQueue[1:]
		}
	}
}

func (ws *webSocket) RegisterSub(channel, roomID string, filters json.RawMessage, subscribeToSelf bool, notifChan chan<- types.KuzzleNotification, onReconnectChannel chan<- interface{}) {
	subs, found := ws.subscriptions.Load(channel)

	if !found {
		subs = map[string]subscription{}
	}

	subs.(map[string]subscription)[roomID] = subscription{
		channel:             channel,
		roomID:              roomID,
		notificationChannel: notifChan,
		onReconnectChannel:  onReconnectChannel,
		filters:             filters,
		subscribeToSelf:     subscribeToSelf,
	}

	ws.subscriptions.Store(channel, subs)
}

func (ws *webSocket) UnregisterSub(roomID string) {
	ws.subscriptions.Range(func(k, v interface{}) bool {
		for k, sub := range v.(map[string]subscription) {
			if sub.roomID == roomID {
				close(sub.onReconnectChannel)
				close(sub.notificationChannel)
				delete(v.(map[string]subscription), k)
			}
		}
		return true
	})
}

func (ws *webSocket) CancelSubs() {
	ws.subscriptions.Range(func(roomId, s interface{}) bool {
		for _, sub := range s.(map[string]subscription) {
			if sub.notificationChannel != nil {
				close(sub.onReconnectChannel)
			}
			ws.subscriptions.Delete(roomId)
		}
		return true
	})
}

func (ws *webSocket) listen() {
	for {
		msg := <-ws.listenChan

		var message types.KuzzleResponse
		json.Unmarshal(msg, &message)

		if s, found := ws.subscriptions.Load(message.RoomId); found {
			var notification types.KuzzleNotification
			_, fromSelf := ws.requestHistory[message.RequestId]

			json.Unmarshal(msg, &notification)

			for _, sub := range s.(map[string]subscription) {
				if sub.notificationChannel != nil && (!fromSelf || sub.subscribeToSelf) {
					sub.notificationChannel <- notification
				}
			}

		} else if c, found := ws.channelsResult.Load(message.RequestId); found {
			if message.Error.Error() != "" && message.Error.Message == "Token expired" {
				ws.EmitEvent(event.TokenExpired, nil)
			}

			// If this is a response to a query we simply broadcast the response to the corresponding channel
			c.(chan<- *types.KuzzleResponse) <- &message
			close(c.(chan<- *types.KuzzleResponse))
			ws.channelsResult.Delete(message.RequestId)
		} else if c, found := ws.channelsResult.Load(message.RoomId); found {
			c.(chan<- *types.KuzzleResponse) <- &message
			close(c.(chan<- *types.KuzzleResponse))
			ws.channelsResult.Delete(message.RequestId)
		} else {
			ws.EmitEvent(event.Discarded, &message)
		}
	}
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (ws *webSocket) AddListener(event int, channel chan<- interface{}) {
	if ws.eventListeners[event] == nil {
		ws.eventListeners[event] = make(map[chan<- interface{}]bool)
	}
	ws.eventListeners[event][channel] = true
}

// Removes all listeners, either from all events and close channels
func (ws *webSocket) RemoveAllListeners(event int) {
	for k := range ws.eventListeners {
		if event == k || event == -1 {
			delete(ws.eventListeners, k)
		}
	}

	for k := range ws.eventListenersOnce {
		if event == k || event == -1 {
			delete(ws.eventListenersOnce, k)
		}
	}
}

// Removes a listener from an event.
func (ws *webSocket) RemoveListener(event int, c chan<- interface{}) {
	delete(ws.eventListeners[event], c)
	delete(ws.eventListenersOnce[event], c)
}

func (ws *webSocket) Once(event int, channel chan<- interface{}) {
	if ws.eventListenersOnce[event] == nil {
		ws.eventListenersOnce[event] = make(map[chan<- interface{}]bool)
	}
	ws.eventListenersOnce[event][channel] = true
}

func (ws *webSocket) ListenerCount(event int) int {
	return len(ws.eventListenersOnce[event]) + len(ws.eventListeners[event])
}

// Emit an event to all registered listeners
func (ws *webSocket) EmitEvent(event int, arg interface{}) {
	for c := range ws.eventListeners[event] {
		c <- arg
	}
	for c := range ws.eventListenersOnce[event] {
		c <- arg
		close(c)
		delete(ws.eventListenersOnce[event], c)
	}
}

func (ws *webSocket) StartQueuing() {
	if ws.state == state.Offline && !ws.autoQueue {
		ws.queuing = true
	}
}

func (ws *webSocket) StopQueuing() {
	if ws.state == state.Offline && !ws.autoQueue {
		ws.queuing = false
	}
}

func (ws *webSocket) ClearQueue() {
	ws.offlineQueue = nil
}

// PlayQueue replays the requests queued during offline mode. Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
func (ws *webSocket) PlayQueue() {
	if ws.state != state.Offline && !ws.autoReplay {
		ws.cleanQueue()
		ws.dequeue()
	}
}

func (ws *webSocket) mergeOfflineQueueWithLoader() error {
	type query struct {
		requestId  string `json:"requestId"`
		controller string `json:"controller"`
		action     string `json:"action"""`
	}

	additionalOfflineQueue := ws.offlineQueueLoader.Load()

	for _, additionalQuery := range additionalOfflineQueue {
		for _, offlineQuery := range ws.offlineQueue {
			q := query{}
			json.Unmarshal(additionalQuery.Query, &q)
			if q.requestId != "" || q.action != "" || q.controller != "" {
				offlineQ := query{}
				json.Unmarshal(offlineQuery.Query, &offlineQ)
				if q.requestId != offlineQ.requestId {
					ws.offlineQueue = append(ws.offlineQueue, additionalQuery)
				} else {
					additionalOfflineQueue = additionalOfflineQueue[:1]
				}
			} else {
				return types.NewError("Invalid offline queue request. One or more missing properties: requestId, action, controller.")
			}
		}
	}
	return nil
}

func (ws *webSocket) dequeue() error {
	if ws.offlineQueueLoader != nil {
		err := ws.mergeOfflineQueueWithLoader()
		if err != nil {
			return err
		}
	}
	// Example from sdk where we have a good use of _
	if len(ws.offlineQueue) > 0 {
		for _, query := range ws.offlineQueue {
			ws.emitRequest(query)
			ws.offlineQueue = ws.offlineQueue[:1]
			ws.EmitEvent(event.OfflineQueuePop, query)
			time.Sleep(ws.replayInterval * time.Millisecond)
			ws.offlineQueue = ws.offlineQueue[:1]
		}
	} else {
		ws.queuing = false
	}
	return nil
}

func (ws *webSocket) emitRequest(query *types.QueryObject) error {
	now := time.Now()
	now = now.Add(-MAX_EMIT_TIMEOUT * time.Second)

	ws.channelsResult.Store(query.RequestId, query.ResChan)

	ws.mu.Lock()
	defer ws.mu.Unlock()
	err := ws.ws.WriteMessage(websocket.TextMessage, query.Query)
	if err != nil {
		return err
	}

	// Track requests made to allow Room.subscribeToSelf to work
	ws.requestHistory[query.RequestId] = time.Now()
	for i, request := range ws.requestHistory {
		if request.Before(now) {
			delete(ws.requestHistory, i)
		}
	}

	return nil
}

func (ws *webSocket) Close() error {
	ws.stopRetryingToConnect = true
	ws.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	ws.state = state.Disconnected

	return ws.ws.Close()
}

func (ws *webSocket) isValidState() bool {
	switch ws.state {
	case state.Initializing, state.Ready, state.Disconnected, state.Error, state.Offline:
		return true
	}
	return false
}

func (ws *webSocket) State() int {
	return ws.state
}

func (ws *webSocket) RequestHistory() map[string]time.Time {
	return ws.requestHistory
}

func (ws *webSocket) AutoQueue() bool {
	return ws.autoQueue
}

func (ws *webSocket) AutoReconnect() bool {
	return ws.autoReconnect
}

func (ws *webSocket) AutoResubscribe() bool {
	return ws.autoResubscribe
}

func (ws *webSocket) AutoReplay() bool {
	return ws.autoReplay
}

func (ws *webSocket) Host() string {
	return ws.host
}

func (ws *webSocket) OfflineQueue() []*types.QueryObject {
	return ws.offlineQueue
}

func (ws *webSocket) OfflineQueueLoader() connection.OfflineQueueLoader {
	return ws.offlineQueueLoader
}

func (ws *webSocket) Port() int {
	return ws.port
}

func (ws *webSocket) QueueFilter() connection.QueueFilter {
	return ws.queueFilter
}

func (ws *webSocket) QueueMaxSize() int {
	return ws.queueMaxSize
}

func (ws *webSocket) QueueTTL() time.Duration {
	return ws.queueTTL
}

func (ws *webSocket) ReplayInterval() time.Duration {
	return ws.replayInterval
}

func (ws *webSocket) ReconnectionDelay() time.Duration {
	return ws.reconnectionDelay
}

func (ws *webSocket) SslConnection() bool {
	return ws.ssl
}

func (ws *webSocket) SetAutoQueue(v bool) {
	ws.autoQueue = v
}

func (ws *webSocket) SetAutoReplay(v bool) {
	ws.autoReplay = v
}

func (ws *webSocket) SetOfflineQueueLoader(v connection.OfflineQueueLoader) {
	ws.offlineQueueLoader = v
}

func (ws *webSocket) SetQueueFilter(v connection.QueueFilter) {
	if v == nil {
		ws.queueFilter = defaultQueueFilter
	} else {
		ws.queueFilter = v
	}
}

func (ws *webSocket) SetQueueMaxSize(v int) {
	ws.queueMaxSize = v
}

func (ws *webSocket) SetQueueTTL(v time.Duration) {
	ws.queueTTL = v
}

func (ws *webSocket) SetReplayInterval(v time.Duration) {
	ws.replayInterval = v
}
