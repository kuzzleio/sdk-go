package websocket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
)

const (
	MAX_EMIT_TIMEOUT = 10
)

type webSocket struct {
	ws      *websocket.Conn
	mu      *sync.Mutex
	queuing bool
	state   int

	listenChan     chan []byte
	channelsResult sync.Map
	subscriptions  *types.RoomList
	lastUrl        string
	wasConnected   bool
	eventListeners map[int]map[chan<- interface{}]bool

	retrying              bool
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
		subscriptions:         &types.RoomList{},
		eventListeners:        make(map[int]map[chan<- interface{}]bool),
		requestHistory:        make(map[string]time.Time),
		autoQueue:             opts.AutoQueue(),
		autoReconnect:         opts.AutoReconnect(),
		autoReplay:            opts.AutoReplay(),
		autoResubscribe:       opts.AutoResubscribe(),
		reconnectionDelay:     opts.ReconnectionDelay(),
		replayInterval:        opts.ReplayInterval(),
		state:                 state.Ready,
		retrying:              false,
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
	if !ws.isValidState() {
		return false, nil
	}

	addr := fmt.Sprintf("%s:%d", ws.host, ws.port)

	var scheme string

	if ws.ssl {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	u := url.URL{Scheme: scheme, Host: addr}
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

	ws.RenewSubscriptions()

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

func (ws *webSocket) Send(query []byte, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse, requestId string) error {
	if ws.state == state.Connected || (options != nil && !options.Queuable()) {
		ws.emitRequest(&types.QueryObject{
			Query:     query,
			ResChan:   responseChannel,
			RequestId: requestId,
		})
	} else if ws.queuing || (options != nil && options.Queuable()) || ws.state == state.Initializing || ws.state == state.Connecting {
		ws.cleanQueue()

		if ws.queueFilter(query) {
			qo := &types.QueryObject{
				Timestamp: time.Now(),
				ResChan:   responseChannel,
				Query:     query,
				RequestId: requestId,
				Options:   options,
			}
			ws.offlineQueue = append(ws.offlineQueue, qo)
			ws.EmitEvent(event.OfflineQueuePush, qo)
		}
	} else {
		ws.discardRequest(responseChannel, query)
	}
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
			ws.offlineQueue = ws.offlineQueue[1:]
		}
	}
}

func (ws *webSocket) RegisterRoom(roomId, id string, room types.IRoom) {
	_, ok := ws.subscriptions.Load(roomId)
	if !ok {
		s := &sync.Map{}
		s.Store(id, room)
		ws.subscriptions.Store(roomId, s)
	}
}

func (ws *webSocket) UnregisterRoom(roomId string) {
	_, ok := ws.subscriptions.Load(roomId)
	if ok {
		ws.subscriptions.Delete(roomId)
	}
}

func (ws *webSocket) listen() {
	for {
		var message types.KuzzleResponse
		var r collection.Room

		msg := <-ws.listenChan

		json.Unmarshal(msg, &message)
		m := message

		json.Unmarshal(m.Result, &r)

		s, ok := ws.subscriptions.Load(m.RoomId)
		if m.RoomId != "" && ok {
			var notification types.KuzzleNotification

			json.Unmarshal(msg, &notification)

			s.(*sync.Map).Range(func(key, value interface{}) bool {
				channel := value.(types.IRoom).RealtimeChannel()
				if channel != nil {
					value.(types.IRoom).RealtimeChannel() <- &notification
				}
				return true
			})
		}

		c, ok := ws.channelsResult.Load(m.RequestId)
		if ok {
			if message.Error != nil && message.Error.Message == "Token expired" {
				ws.EmitEvent(event.JwtExpired, nil)
			}

			// If this is a response to a query we simply broadcast the response to the corresponding channel
			c.(chan<- *types.KuzzleResponse) <- &message
			close(c.(chan<- *types.KuzzleResponse))
			ws.channelsResult.Delete(m.RequestId)
		}
	}
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (ws *webSocket) AddListener(event int, channel chan<- interface{}) {
	ws.eventListeners[event] = make(map[chan<- interface{}]bool)
	ws.eventListeners[event][channel] = true
}

// Removes all listeners, either from all events and close channels
func (ws *webSocket) RemoveAllListeners(event int) {
	for k := range ws.eventListeners {
		if event == k || event == -1 {
			for c := range ws.eventListeners[k] {
				close(c)
			}
			delete(ws.eventListeners, k)
		}
	}
}

// Removes a listener from an event.
func (ws *webSocket) RemoveListener(event int, c chan<- interface{}) {
	delete(ws.eventListeners[event], c)
}

// Emit an event to all registered listeners
func (ws *webSocket) EmitEvent(event int, arg interface{}) {
	for c := range ws.eventListeners[event] {
		c <- arg
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

// ReplayQueue replays the requests queued during offline mode. Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
func (ws *webSocket) ReplayQueue() {
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

func (ws *webSocket) RenewSubscriptions() {
	ws.subscriptions.Range(func(key, value interface{}) bool {
		value.(*sync.Map).Range(func(key, value interface{}) bool {
			value.(types.IRoom).Renew(value.(types.IRoom).Filters(), value.(types.IRoom).RealtimeChannel(), value.(types.IRoom).ResponseChannel())
			return true
		})

		return true
	})
}

func (ws *webSocket) Rooms() *types.RoomList {
	return ws.subscriptions
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
