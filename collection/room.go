package collection

import (
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

type roomState int

const (
	inactive roomState = iota
	subscribing
	active
)

type Room struct {
	// object configuration
	filters         interface{}
	channel         string
	roomId          string
	requestId       string
	scope           string
	state           string
	subscribeToSelf bool
	users           string
	Volatile        types.VolatileData
	err             error
	autoResubscribe bool

	// internal properties
	collection    *Collection
	id            string
	internalState roomState
	isListening   bool
	subscribing   bool

	// channels
	realtimeNotificationChannel chan<- *types.KuzzleNotification
	subscribeResponseChan       chan *types.SubscribeResponse
	onReconnect                 chan interface{}
	onDisconnect                chan interface{}
	onTokenExpired              chan interface{}
}

// NewRoom instanciates a new Room; this type is the result of a subscription request,
// allowing to manipulate the subscription itself.
// In Kuzzle, you don't exactly subscribe to a room or a topic but, instead, you subscribe to documents.
// What it means is that, to subscribe, you provide to Kuzzle a set of matching filters.
// Once you have subscribed, if a pub/sub message is published matching your filters, or if a matching stored
// document change (because it is created, updated or deleted), then you'll receive a notification about it.
func NewRoom(c *Collection, filters interface{}, opts types.RoomOptions) *Room {
	if opts == nil {
		opts = types.NewRoomOptions()
	}

	u, _ := uuid.NewV4()
	r := &Room{
		scope:           opts.Scope(),
		state:           opts.State(),
		users:           opts.Users(),
		filters:         filters,
		id:              u.String(),
		collection:      c,
		subscribeToSelf: opts.SubscribeToSelf(),
		Volatile:        opts.Volatile(),
		internalState:   inactive,
		isListening:     false,
		autoResubscribe: opts.AutoResubscribe(),
		onReconnect:     make(chan interface{}),
		onDisconnect:    make(chan interface{}),
		onTokenExpired:  make(chan interface{}),
	}
	return r
}

// AddListener Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (room *Room) AddListener(event int, channel chan<- interface{}) {
	room.collection.Kuzzle.AddListener(event, channel)
}

// On is an alias to the AddListener function
func (room *Room) On(event int, channel chan<- interface{}) {
	room.collection.Kuzzle.AddListener(event, channel)
}

// Remove all listener by event type or all listener if event == -1
func (room *Room) RemoveAllListeners(event int) {
	room.collection.Kuzzle.RemoveAllListeners(event)
}

// RemoveListener removes a listener
func (room *Room) RemoveListener(event int, channel chan<- interface{}) {
	room.collection.Kuzzle.RemoveListener(event, channel)
}

func (room *Room) Once(event int, channel chan<- interface{}) {
	room.collection.Kuzzle.Once(event, channel)
}

func (room *Room) ListenerCount(event int) int {
	return room.collection.Kuzzle.ListenerCount(event)
}

// RealtimeChannel returns the channel handling the notifications received by this room
func (room *Room) RealtimeChannel() chan<- *types.KuzzleNotification {
	return room.realtimeNotificationChannel
}

// RoomId returns the kuzzle room unique identifier
func (room *Room) RoomId() string {
	return room.roomId
}

// Id returns the internal object identifier (internal use only)
func (room *Room) Id() string {
	return room.id
}

// Filters returns the room's filters
func (room *Room) Filters() interface{} {
	return room.filters
}

// ResponseChannel returns the channel handling the subscription request result
func (room *Room) ResponseChannel() chan<- *types.SubscribeResponse {
	return room.subscribeResponseChan
}

// Channel is the getter for the channel's unique identifier
func (room *Room) Channel() string {
	return room.channel
}

func (room *Room) Scope() string {
	return room.scope
}

func (room *Room) State() string {
	return room.state
}

func (room *Room) Users() string {
	return room.users
}

//OnDone Calls the provided callback when the subscription finishes.
func (room *Room) OnDone(c chan *types.SubscribeResponse) *Room {
	if room.err != nil {
		c <- &types.SubscribeResponse{Error: room.err}
	} else if room.internalState == active {
		c <- nil
	} else {
		room.subscribeResponseChan = c
	}

	return room
}

func (room *Room) SubscribeToSelf() bool {
	return room.subscribeToSelf
}
