package core

import (
	"github.com/kuzzleio/sdk-go/wrappers"
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/kuzzleio/sdk-go/event"
	"sync"
	"github.com/kuzzleio/sdk-go/state"
)

type IKuzzle interface {
	Query(types.KuzzleRequest, chan<- types.KuzzleResponse, chan<- types.KuzzleNotification)
}

type Kuzzle struct {
	Host 													string
	socket 												*wrappers.WebSocket

	channelsResult							  map [string] chan<- types.KuzzleResponse
	subscriptions                 map[string] chan<- types.KuzzleNotification
	wasConnected                  bool
	lastUrl                       string
	message                       chan []byte
	eventListeners								map [int] chan <- interface{}
	mu														*sync.Mutex
	state 												int
	jwtToken											string
}

// Kuzzle constructor
func NewKuzzle(host string) *Kuzzle {
	return &Kuzzle{
		Host: host,
		channelsResult: make(map[string] chan<- types.KuzzleResponse),
		subscriptions: make(map[string] chan<- types.KuzzleNotification),
		eventListeners: make(map[int] chan<- interface{}),
		mu: &sync.Mutex{},
		socket: wrappers.NewWebSocket(),
		state: state.Initializing,
	}
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (k *Kuzzle) AddListener(event int, channel chan <- interface{}) {
	k.eventListeners[event] = channel
}

// Emit an event to all registered listeners
func (k *Kuzzle) emitEvent(event int, arg interface{}) {
	if k.eventListeners[event] != nil {
		k.eventListeners[event] <- arg
	}
}

// Connects to a Kuzzle instance using the provided host and port.
func (k *Kuzzle) Connect() error {
	var err error

	if !k.isValidState() {
		return nil
	}

	k.message, err = k.socket.Connect(k.Host)
	if err == nil {
		if k.lastUrl != k.Host {
			k.wasConnected = false
			k.lastUrl = k.Host
		}

		if k.wasConnected {
			k.emitEvent(event.Reconnected, nil)
			k.state = state.Connected
			if k.jwtToken != "" {
				// todo avoid import cycle
				//go func() {
				//	res, err := kuzzle.CheckToken(k, k.jwtToken)
        //
				//	if err != nil {
				//		k.jwtToken = ""
				//		k.emitEvent(event.JwtTokenExpired, nil)
				//		k.Reconnect()
				//		return
				//	}
        //
				//	if !res.Valid {
				//		k.jwtToken = ""
				//		k.emitEvent(event.JwtTokenExpired, nil)
				//	}
				//	k.Reconnect()
				//}()
			}
		} else {
			k.emitEvent(event.Connected, nil)
			k.state = state.Connected
			//todo renew subscription
		}

		go k.listen()
		return nil
	}
	k.state = state.Error
	k.emitEvent(event.Error, err)

	return err
}

func (k Kuzzle) Reconnect() {
	// todo auto resubscribe

	//todo auto replay

	k.emitEvent(event.Reconnected, nil)
}

func (k *Kuzzle) listen() {
	for {
		var message types.KuzzleResponse
		var room types.Room

		msg := <-k.message

		json.Unmarshal(msg, &message)
		m := message

		if len(k.channelsResult) > 0 {

			json.Unmarshal(m.Result, &room)

			if room.Channel != "" {
				// If this is a response to a subscribe we save the channel to a map
			  k.subscriptions[room.Channel] = k.subscriptions[m.RoomId]
			} else if m.RoomId != "" && k.subscriptions[m.RoomId] != nil {
				// If this is a notification from a subscribe then we unmarshal it as a KuzzleSubscription object
				var notification types.KuzzleNotification

				json.Unmarshal(m.Result, &notification.Result)
				k.subscriptions[m.RoomId] <- notification
			}

			if k.channelsResult[m.RequestId] != nil {
				// If this is a response to a query we simply broadcast the response to the corresponding channel
				k.channelsResult[m.RequestId] <- message
				close(k.channelsResult[m.RequestId])
				delete(k.channelsResult, m.RequestId)
			}
		}
	}
}

// Instantiates a new Collection object.
func (k *Kuzzle) Collection(collection, index string) *Collection {
	return NewCollection(k, collection, index)
}

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query types.KuzzleRequest, res chan<- types.KuzzleResponse, subscription chan<- types.KuzzleNotification) {
	requestId := uuid.NewV4().String()

	query.RequestId = requestId
	k.mu.Lock()
	k.channelsResult[requestId] = res
	k.mu.Unlock()
	if subscription != nil {
		k.mu.Lock()
		k.subscriptions[requestId] = subscription
		k.mu.Unlock()
	}

	type body struct {}
	if query.Body == nil {
		query.Body = &body{}
	}

	json, err := json.Marshal(query)
	if err != nil {
		res <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
		return
	}

	err = k.socket.Send(json)
	if err != nil {
		res <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
		return
	}
}

func (k *Kuzzle) Disconnect() error {
	err := k.socket.Close()

	if err != nil {
		return err
	}
	k.wasConnected = false

	return nil
}

func (k Kuzzle) isValidState() bool {
	switch k.state {
	case state.Initializing, state.Ready, state.Disconnected, state.Error, state.Offline:
		return true
	}
	return false
}
