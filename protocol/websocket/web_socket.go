// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/gorilla/websocket"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	MAX_EMIT_TIMEOUT = 10
)

type subscription struct {
	roomID              string
	channel             string
	filters             json.RawMessage
	notificationChannel chan<- types.NotificationResult
	onReconnectChannel  chan<- interface{}
	subscribeToSelf     bool
}

type WebSocket struct {
	ws      *websocket.Conn
	mu      *sync.Mutex
	queuing bool
	state   int

	listenChan         chan []byte
	channelsResult     sync.Map
	subscriptions      sync.Map
	lastUrl            string
	wasConnected       bool
	eventListeners     map[int]map[chan<- json.RawMessage]bool
	eventListenersOnce map[int]map[chan<- json.RawMessage]bool

	requestHistory map[string]time.Time

	host    string
	port    int
	ssl     bool
	headers *http.Header
}

// NewWebSocket instanciates a new webSocket connection object
func NewWebSocket(host string, options types.Options) *WebSocket {
	var opts types.Options

	if options == nil {
		opts = types.NewOptions()
	} else {
		opts = options
	}

	ws := &WebSocket{
		mu:                 &sync.Mutex{},
		channelsResult:     sync.Map{},
		subscriptions:      sync.Map{},
		eventListeners:     make(map[int]map[chan<- json.RawMessage]bool),
		eventListenersOnce: make(map[int]map[chan<- json.RawMessage]bool),
		requestHistory:     make(map[string]time.Time),
		state:              state.Ready,
		host:               host,
		port:               opts.Port(),
		ssl:                opts.SslConnection(),
		headers:            opts.Headers(),
	}
	ws.host = host

	ws.state = state.Offline

	return ws
}

//Connect connects to a kuzzle instance
func (ws *WebSocket) Connect() (bool, error) {
	if ws.state != state.Offline {
		return false, nil
	}

	ws.state = state.Connecting

	addr := fmt.Sprintf("%s:%d", ws.host, ws.port)

	ws.wasConnected = ws.lastUrl == addr
	ws.lastUrl = addr

	var scheme string

	if ws.ssl {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	u := url.URL{Scheme: scheme, Host: addr}

	var headers http.Header

	if ws.headers != nil {
		headers = *ws.headers
	}

	socket, _, err := websocket.DefaultDialer.Dial(u.String(), headers)

	if err != nil {
		ws.state = state.Offline
		ws.EmitEvent(event.NetworkError, err)

		return false, err
	}

	ws.ws = socket
	ws.state = state.Connected
	ws.queuing = false

	if ws.wasConnected {
		ws.EmitEvent(event.Reconnected, nil)
	} else {
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
				ws.EmitEvent(event.Disconnected, nil)
				return
			}
		}
	}()

	return ws.wasConnected, err
}

func (ws *WebSocket) Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error {
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

func (ws *WebSocket) discardRequest(responseChannel chan<- *types.KuzzleResponse, query []byte) {
	if responseChannel != nil {
		responseChannel <- &types.KuzzleResponse{Status: 400, Error: types.NewError("Unable to execute request: not connected to a Kuzzle server.\nDiscarded request: "+string(query), 400)}
	}
}

func (ws *WebSocket) RegisterSub(channel, roomID string, filters json.RawMessage, subscribeToSelf bool, notifChan chan<- types.NotificationResult, onReconnectChannel chan<- interface{}) {
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

func (ws *WebSocket) UnregisterSub(roomID string) {
	ws.subscriptions.Range(func(k, v interface{}) bool {
		for k, sub := range v.(map[string]subscription) {
			if sub.roomID == roomID {
				if sub.onReconnectChannel != nil {
					close(sub.onReconnectChannel)
				}
				if sub.notificationChannel != nil {
					close(sub.notificationChannel)
				}
				delete(v.(map[string]subscription), k)
			}
		}
		return true
	})
}

func (ws *WebSocket) CancelSubs() {
	ws.subscriptions.Range(func(roomId, s interface{}) bool {
		for _, sub := range s.(map[string]subscription) {
			if sub.onReconnectChannel != nil {
				close(sub.onReconnectChannel)
			}
			ws.subscriptions.Delete(roomId)
		}
		return true
	})
}

func (ws *WebSocket) listen() {
	for msg := range ws.listenChan {
		var message types.KuzzleResponse
		json.Unmarshal(msg, &message)

		if s, found := ws.subscriptions.Load(message.RoomId); found {
			var notification types.NotificationResult
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
func (ws *WebSocket) AddListener(event int, channel chan<- json.RawMessage) {
	if ws.eventListeners[event] == nil {
		ws.eventListeners[event] = make(map[chan<- json.RawMessage]bool)
	}
	ws.eventListeners[event][channel] = true
}

// Removes all listeners, either from all events and close channels
func (ws *WebSocket) RemoveAllListeners(event int) {
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
func (ws *WebSocket) RemoveListener(event int, c chan<- json.RawMessage) {
	delete(ws.eventListeners[event], c)
	delete(ws.eventListenersOnce[event], c)
}

func (ws *WebSocket) Once(event int, channel chan<- json.RawMessage) {
	if ws.eventListenersOnce[event] == nil {
		ws.eventListenersOnce[event] = make(map[chan<- json.RawMessage]bool)
	}
	ws.eventListenersOnce[event][channel] = true
}

func (ws *WebSocket) ListenerCount(event int) int {
	return len(ws.eventListenersOnce[event]) + len(ws.eventListeners[event])
}

// Emit an event to all registered listeners
func (ws *WebSocket) EmitEvent(event int, arg interface{}) {
	for c := range ws.eventListeners[event] {
		json, _ := json.Marshal(arg)
		c <- json
	}
	for c := range ws.eventListenersOnce[event] {
		json, _ := json.Marshal(arg)
		c <- json
		close(c)
		delete(ws.eventListenersOnce[event], c)
	}
}

func (ws *WebSocket) IsReady() bool {
	return ws != nil && ws.state == state.Connected
}

func (ws *WebSocket) emitRequest(query *types.QueryObject) error {
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

func (ws *WebSocket) Close() error {
	if ws.ws == nil {
		return nil
	}

	ws.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	ws.state = state.Disconnected

	return ws.ws.Close()
}

func (ws *WebSocket) isValidState() bool {
	switch ws.state {
	case state.Initializing, state.Ready, state.Disconnected, state.Error, state.Offline:
		return true
	}
	return false
}

func (ws *WebSocket) State() int {
	return ws.state
}

func (ws *WebSocket) RequestHistory() map[string]time.Time {
	return ws.requestHistory
}

func (ws *WebSocket) Host() string {
	return ws.host
}

func (ws *WebSocket) Port() int {
	return ws.port
}

func (ws *WebSocket) SslConnection() bool {
	return ws.ssl
}
