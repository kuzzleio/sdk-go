package collection

import (
	"container/list"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
	"time"
)

type Room struct {
	RequestId       string          `json:"RequestId"`
	RoomId          string          `json:"roomId"`
	Channel         string          `json:"channel"`
	result          json.RawMessage `json:"result"`
	scope           string          `json:"scope"`
	state           string          `json:"state"`
	user            string          `json:"user"`
	subscribeToSelf bool            `json:"subscribeToSelf"`

	collection                  Collection                      `json:"-"`
	lastRenewal                 time.Duration                   `json:"-"`
	renewalDelay                time.Duration                   `json:"-"`
	RealtimeNotificationChannel chan<- types.KuzzleNotification `json:"-"`
	subscribeResponseChan       chan<- types.SubscribeResponse

	pendingSubscriptions map[string]chan<- types.KuzzleNotification `json:"-"`
	subscribing          bool                                       `json:"-"`
	queue                list.List                                  `json:"-"`
	id                   string                                     `json:"-"`
	Volatile             types.VolatileData                         `json:"-"`
	filters              interface{}                                `json:"-"`
}

/**
 * This object is the result of a subscription request, allowing to manipulate the subscription itself.
 * In Kuzzle, you don't exactly subscribe to a room or a topic but, instead, you subscribe to documents.
 * What it means is that, to subscribe, you provide to Kuzzle a set of matching filters.
 * Once you have subscribed, if a pub/sub message is published matching your filters, or if a matching stored
 * document change (because it is created, updated or deleted), then you'll receive a notification about it.
 */
func NewRoom(c Collection, opts types.RoomOptions) *Room {
	if opts == nil {
		opts = types.NewRoomOptions()
	}
	r := &Room{
		renewalDelay:         500 * time.Millisecond,
		scope:                opts.GetScope(),
		state:                opts.GetState(),
		user:                 opts.GetUser(),
		id:                   uuid.NewV4().String(),
		collection:           c,
		pendingSubscriptions: make(map[string]chan<- types.KuzzleNotification),
		subscribeToSelf:      opts.GetSubscribeToSelf(),
		Volatile:             opts.GetVolatile(),
	}
	r.queue.Init()

	return r
}

func (room Room) GetRealtimeChannel() chan<- types.KuzzleNotification {
	return room.RealtimeNotificationChannel
}

func (room Room) isReady() bool {
	return *room.collection.Kuzzle.State == state.Connected && !room.subscribing
}

func (room Room) GetRoomId() string {
	return room.RoomId
}

func (room Room) GetFilters() interface{} {
	return room.filters
}

func (room Room) GetResponseChannel() chan<- types.SubscribeResponse {
	return room.subscribeResponseChan
}
