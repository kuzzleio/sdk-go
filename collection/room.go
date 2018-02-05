package collection

import (
	"container/list"
	"encoding/json"

	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

type Room struct {
	requestId       string
	roomId          string
	channel         string
	result          json.RawMessage
	scope           string
	state           string
	users           string
	SubscribeToSelf bool

	collection                  *Collection
	realtimeNotificationChannel chan<- *types.KuzzleNotification
	subscribeResponseChan       chan<- *types.SubscribeResponse

	pendingSubscriptions map[string]chan<- *types.KuzzleNotification
	subscribing          bool
	queue                *list.List
	id                   string
	Volatile             types.VolatileData
	filters              interface{}
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
		scope:                opts.Scope(),
		state:                opts.State(),
		users:                opts.Users(),
		id:                   u.String(),
		collection:           c,
		pendingSubscriptions: make(map[string]chan<- *types.KuzzleNotification),
		SubscribeToSelf:      opts.SubscribeToSelf(),
		Volatile:             opts.Volatile(),
		queue:                &list.List{},
		filters:              filters,
	}
	r.queue.Init()

	return r
}

// RealtimeChannel return the room's ReatimeNotificationChannel
func (room *Room) RealtimeChannel() chan<- *types.KuzzleNotification {
	return room.realtimeNotificationChannel
}

// isReady returns true if the room is ready
func (room *Room) isReady() bool {
	return room.collection.Kuzzle.State() == state.Connected && !room.subscribing
}

// RoomId returns the room's id
func (room *Room) RoomId() string {
	return room.roomId
}

// Filters returns the room's filters
func (room *Room) Filters() interface{} {
	return room.filters
}

// ResponseChannel returns the room's response channel
func (room *Room) ResponseChannel() chan<- *types.SubscribeResponse {
	return room.subscribeResponseChan
}

// Channel is the getter for the channel unique identifier value
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
