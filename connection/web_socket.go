package connection

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"net/url"
	"sync"
	"time"
)

const (
	MAX_EMIT_TIMEOUT = 10
	EVENT_TIMEOUT    = 20
)

type WebSocket struct {
	ws      *websocket.Conn
	mu      *sync.Mutex
	queuing bool
	state   int

	listenChan     chan []byte
	channelsResult map[string]chan<- types.KuzzleResponse
	subscriptions  map[string]chan<- types.KuzzleNotification
	lastUrl        string
	host           string
	wasConnected   bool
	eventListeners map[int]chan<- interface{}

	autoQueue             bool
	autoReconnect         bool
	autoReplay            bool
	autoResubscribe       bool
	queueTTL              time.Duration
	offlineQueue          []types.QueryObject
	offlineQueueLoader    OfflineQueueLoader
	queueFilter           QueueFilter
	queueMaxSize          int
	reconnectionDelay     time.Duration
	replayInterval        time.Duration
	retrying              bool
	stopRetryingToConnect bool
	RequestHistory        map[string]time.Time
}

type QueueFilter interface {
	Filter(interface{}) bool
}

type defaultQueueFilter struct{}

func (qf defaultQueueFilter) Filter(interface{}) bool {
	return true
}

type OfflineQueueLoader interface {
	load() []types.QueryObject
}

func (ws *WebSocket) SetQueueFilter(queueFilter QueueFilter) {
	ws.queueFilter = queueFilter
}

func NewWebSocket(host string, options types.Options) Connection {
	var opts types.Options

	if options == nil {
		opts = types.NewOptions()
	} else {
		opts = options
	}
	ws := &WebSocket{
		mu:                    &sync.Mutex{},
		queueTTL:              opts.GetQueueTTL(),
		offlineQueue:          make([]types.QueryObject, 0),
		queueMaxSize:          opts.GetQueueMaxSize(),
		channelsResult:        make(map[string]chan<- types.KuzzleResponse),
		subscriptions:         make(map[string]chan<- types.KuzzleNotification),
		eventListeners:        make(map[int]chan<- interface{}),
		RequestHistory:        make(map[string]time.Time),
		autoQueue:             opts.GetAutoQueue(),
		autoReconnect:         opts.GetAutoReconnect(),
		autoReplay:            opts.GetAutoReplay(),
		autoResubscribe:       opts.GetAutoResubscribe(),
		reconnectionDelay:     opts.GetReconnectionDelay(),
		replayInterval:        opts.GetReplayInterval(),
		state:                 state.Ready,
		retrying:              false,
		stopRetryingToConnect: false,
		queueFilter:           &defaultQueueFilter{},
	}
	ws.host = host

	if opts.GetOfflineMode() == types.Auto {
		ws.autoReconnect = true
		ws.autoQueue = true
		ws.autoReplay = true
		ws.autoResubscribe = true
	}
	return ws
}

func (ws *WebSocket) Connect() (bool, error) {
	if !ws.isValidState() {
		return false, nil
	}

	addr := flag.String("addr", ws.host, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr}
	resChan := make(chan []byte)

	socket, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		if ws.autoReconnect && !ws.retrying && !ws.stopRetryingToConnect {
			for err != nil {
				ws.retrying = true
				ws.EmitEvent(event.NetworkError, err)
				time.Sleep(ws.reconnectionDelay)
				ws.retrying = false
				socket, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
			}
		}
		ws.EmitEvent(event.NetworkError, err)
		return ws.wasConnected, err
	}

	if ws.lastUrl != ws.host {
		ws.wasConnected = false
		ws.lastUrl = ws.host
	}

	if ws.wasConnected {
		ws.EmitEvent(event.Reconnected, nil)
		ws.state = state.Connected
	} else {
		ws.EmitEvent(event.Connected, nil)
		ws.state = state.Connected
	}

	ws.ws = socket
	ws.stopRetryingToConnect = false

	if ws.autoReplay {
		ws.cleanQueue()
		ws.dequeue()
	}
	//todo renew subscription

	go func() {
		for {
			_, message, err := ws.ws.ReadMessage()

			if err != nil {
				close(resChan)
				ws.ws.Close()
				ws.state = state.Offline
				if ws.autoQueue {
					ws.queuing = true
				}
				ws.EmitEvent(event.Disconnected, nil)
				return
			}
			go func() {
				resChan <- message
			}()
		}
	}()

	ws.listenChan = resChan
	go ws.listen()
	return ws.wasConnected, err
}

func (ws *WebSocket) Send(query []byte, options types.QueryOptions, responseChannel chan<- types.KuzzleResponse, requestId string) error {
	if ws.state == state.Connected || (options != nil && !options.GetQueuable()) {
		ws.emitRequest(types.QueryObject{
			Query:     query,
			ResChan:   responseChannel,
			RequestId: requestId,
		})
	} else if ws.queuing || (options != nil && options.GetQueuable()) || ws.state == state.Initializing || ws.state == state.Connecting {
		ws.cleanQueue()

		if ws.queueFilter.Filter(query) {
			qo := types.QueryObject{
				Timestamp: time.Now(),
				ResChan:   responseChannel,
				Query:     query,
				RequestId: requestId,
				Options:   options,
			}
			ws.offlineQueue = append(ws.offlineQueue, qo)
			ws.EmitEvent(event.OfflieQueuePush, qo)
		}
	} else {
		ws.discardRequest(responseChannel, query)
	}
	return nil
}

func (ws *WebSocket) discardRequest(responseChannel chan<- types.KuzzleResponse, query []byte) {
	if responseChannel != nil {
		responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: "Unable to execute request: not connected to a Kuzzle server.\nDiscarded request: " + string(query), Status: 400}}
	}
}

// Clean up the queue, ensuring the queryTTL and queryMaxSize properties are respected
func (ws *WebSocket) cleanQueue() {
	now := time.Now()
	now = now.Add(-ws.queueTTL * time.Millisecond)

	// Clean queue of timed out query
	if ws.queueTTL > 0 {
		var query types.QueryObject
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
			if eventListener != nil {
				ws.eventListeners[event.OfflineQueuePop] <- ws.offlineQueue[0]
			}
			ws.offlineQueue = ws.offlineQueue[1:]
		}
	}
}

func (ws *WebSocket) listen() {
	for {
		var message types.KuzzleResponse
		var room types.Room

		msg := <-ws.listenChan

		json.Unmarshal(msg, &message)
		m := message

		if len(ws.channelsResult) > 0 {

			json.Unmarshal(m.Result, &room)

			if room.Channel != "" {
				// If this is a response to a subscribe we save the channel to a map
				ws.subscriptions[room.Channel] = ws.subscriptions[m.RoomId]
			} else if m.RoomId != "" && ws.subscriptions[m.RoomId] != nil {
				// If this is a notification from a subscribe then we unmarshal it as a KuzzleSubscription object
				var notification types.KuzzleNotification

				json.Unmarshal(m.Result, &notification.Result)
				ws.subscriptions[m.RoomId] <- notification
			}

			if ws.channelsResult[m.RequestId] != nil {
				if message.Error.Message == "Token expired" {
					ws.EmitEvent(event.JwtExpired, nil)
				}

				// If this is a response to a query we simply broadcast the response to the corresponding channel
				ws.channelsResult[m.RequestId] <- message
				close(ws.channelsResult[m.RequestId])
				delete(ws.channelsResult, m.RequestId)
			}
		}
	}
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (ws *WebSocket) AddListener(event int, channel chan<- interface{}) {
	ws.eventListeners[event] = channel
}

// Removes all listeners, either from all events and close channels
func (ws *WebSocket) RemoveAllListeners() {
	for k := range ws.eventListeners {
		if ws.eventListeners[k] != nil {
			close(ws.eventListeners[k])
		}
		delete(ws.eventListeners, k)
	}
}

// Removes a listener from an event.
func (ws *WebSocket) RemoveListener(event int) {
	delete(ws.eventListeners, event)
}

// Emit an event to all registered listeners
func (ws *WebSocket) EmitEvent(event int, arg interface{}) {
	if ws.eventListeners[event] != nil {
		ws.eventListeners[event] <- arg
	}
}

func (ws *WebSocket) StartQueuing() {
	if ws.state == state.Offline && !ws.autoQueue {
		ws.queuing = true
	}
}

func (ws *WebSocket) StopQueuing() {
	if ws.state == state.Offline && !ws.autoQueue {
		ws.queuing = false
	}
}

func (ws *WebSocket) FlushQueue() {
	ws.offlineQueue = ws.offlineQueue[:cap(ws.offlineQueue)]
}

func (ws *WebSocket) ReplayQueue() {
	if ws.state != state.Offline && !ws.autoReplay {
		ws.cleanQueue()
		ws.dequeue()
	}
}

func (ws *WebSocket) mergeOfflineQueueWithLoader() error {
	type query struct {
		requestId  string `json:"requestId"`
		controller string `json:"controller"`
		action     string `json:"action"""`
	}

	additionalOfflineQueue := ws.offlineQueueLoader.load()

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
				return errors.New("Invalid offline queue request. One or more missing properties: requestId, action, controller.")
			}
		}
	}
	return nil
}

func (ws *WebSocket) dequeue() error {
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

func (ws *WebSocket) emitRequest(query types.QueryObject) error {
	now := time.Now()
	now = now.Add(-MAX_EMIT_TIMEOUT * time.Second)

	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.channelsResult[query.RequestId] = query.ResChan
	// todo write room feature for subscribe
	//if subscription != nil {
	//	k.mu.Lock()
	//	k.subscriptions[requestId] = subscription
	//	k.mu.Unlock()
	//}

	err := ws.ws.WriteMessage(websocket.TextMessage, query.Query)
	if err != nil {
		return err
	}

	// Track requests made to allow Room.subscribeToSelf to work
	ws.RequestHistory[query.RequestId] = time.Now()
	for i, request := range ws.RequestHistory {
		if request.Before(now) {
			delete(ws.RequestHistory, i)
		}
	}

	return nil
}

func (ws *WebSocket) Close() error {
	ws.stopRetryingToConnect = true
	ws.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	ws.state = state.Disconnected

	return ws.ws.Close()
}

func (ws WebSocket) GetOfflineQueue() *[]types.QueryObject {
	return &ws.offlineQueue
}

func (ws *WebSocket) isValidState() bool {
	switch ws.state {
	case state.Initializing, state.Ready, state.Disconnected, state.Error, state.Offline:
		return true
	}
	return false
}

func (ws *WebSocket) GetState() *int {
	return &ws.state
}
