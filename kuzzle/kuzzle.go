// Kuzzle Entry point and main struct for the entire SDK
package kuzzle

import (
	"sync"
	"time"

	"github.com/kuzzleio/sdk-go/auth"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/document"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/index"
	"github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/server"
	"github.com/kuzzleio/sdk-go/types"
)

const version = "1.0.0"

type Kuzzle struct {
	socket connection.Connection

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	defaultIndex   string
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory map[string]time.Time
	volatile       types.VolatileData

	MemoryStorage *ms.Ms
	Security      *security.Security
	Auth          *auth.Auth
	Server        *server.Server
	Document      *document.Document
	Index         *index.Index
	Collection    *collection.Collection
}

// NewKuzzle is the Kuzzle constructor
func NewKuzzle(c connection.Connection, options types.Options) (*Kuzzle, error) {
	if c == nil {
		return nil, types.NewError("Connection is nil")
	}

	if options == nil {
		options = types.NewOptions()
	}

	k := &Kuzzle{
		socket:       c,
		version:      version,
		defaultIndex: options.DefaultIndex(),
	}

	k.RequestHistory = k.socket.RequestHistory()
	k.MemoryStorage = &ms.Ms{k}
	k.Security = security.NewSecurity(k)
	k.Auth = auth.NewAuth(k)

	k.RequestHistory = k.socket.RequestHistory()

	k.defaultIndex = options.DefaultIndex()
	k.Server = server.NewServer(k)
	k.Collection = collection.NewCollection(k)
	k.Document = document.NewDocument(k)
	k.Index = index.NewIndex(k)

	var err error

	if options.Connect() == types.Auto {
		err = k.Connect()
	}

	return k, err
}

// Connect connects to a Kuzzle instance using the provided host and port.
func (k *Kuzzle) Connect() error {
	wasConnected, err := k.socket.Connect()
	if err == nil {
		if k.lastUrl != k.socket.Host() {
			k.wasConnected = false
			k.lastUrl = k.socket.Host()
		}

		if wasConnected {
			if k.jwt != "" {
				go func() {
					res, err := k.Auth.CheckToken(k.jwt)

					if err != nil {
						k.jwt = ""
						k.socket.EmitEvent(event.TokenExpired, nil)
						return
					}

					if !res.Valid {
						k.jwt = ""
						k.socket.EmitEvent(event.TokenExpired, nil)
					}
				}()
			}
		}
		return nil
	}
	return types.NewError(err.Error())
}

// Jwt get internal jwtToken used to request kuzzle.
func (k *Kuzzle) Jwt() string {
	return k.jwt
}

func (k *Kuzzle) SetJwt(token string) {
	k.jwt = token

	if token != "" {
		k.socket.EmitEvent(event.LoginAttempt, &types.LoginAttempt{Success: true})
	}
}

// UnsetJwt unset the authentication token and cancel all subscriptions
func (k *Kuzzle) UnsetJwt() {
	k.jwt = ""

	rooms := k.socket.Rooms()
	if rooms != nil {
		k.socket.Rooms().Range(func(key, value interface{}) bool {
			value.(*sync.Map).Range(func(key, value interface{}) bool {
				room := value.(types.IRoom)
				room.Subscribe(room.RealtimeChannel())
				return true
			})

			return true
		})
	}
}

func (k *Kuzzle) RegisterRoom(room types.IRoom) {
	k.socket.RegisterRoom(room)
}

func (k *Kuzzle) UnregisterRoom(roomId string) {
	k.socket.UnregisterRoom(roomId)
}

func (k *Kuzzle) State() int {
	return k.socket.State()
}

func (k *Kuzzle) AutoQueue() bool {
	return k.socket.AutoQueue()
}

func (k *Kuzzle) AutoReconnect() bool {
	return k.socket.AutoReconnect()
}

func (k *Kuzzle) AutoResubscribe() bool {
	return k.socket.AutoResubscribe()
}

func (k *Kuzzle) AutoReplay() bool {
	return k.socket.AutoReplay()
}

func (k *Kuzzle) Host() string {
	return k.socket.Host()
}

func (k *Kuzzle) OfflineQueue() []*types.QueryObject {
	return k.socket.OfflineQueue()
}

func (k *Kuzzle) OfflineQueueLoader() connection.OfflineQueueLoader {
	return k.socket.OfflineQueueLoader()
}

func (k *Kuzzle) Port() int {
	return k.socket.Port()
}

func (k *Kuzzle) QueueFilter() connection.QueueFilter {
	return k.socket.QueueFilter()
}

func (k *Kuzzle) QueueMaxSize() int {
	return k.socket.QueueMaxSize()
}

func (k *Kuzzle) QueueTTL() time.Duration {
	return k.socket.QueueTTL()
}

func (k *Kuzzle) ReplayInterval() time.Duration {
	return k.socket.ReplayInterval()
}

func (k *Kuzzle) ReconnectionDelay() time.Duration {
	return k.socket.ReconnectionDelay()
}

func (k *Kuzzle) SslConnection() bool {
	return k.socket.SslConnection()
}

func (k *Kuzzle) SetAutoQueue(v bool) {
	k.socket.SetAutoQueue(v)
}

func (k *Kuzzle) SetAutoReplay(v bool) {
	k.socket.SetAutoReplay(v)
}

func (k *Kuzzle) SetOfflineQueueLoader(v connection.OfflineQueueLoader) {
	k.socket.SetOfflineQueueLoader(v)
}

func (k *Kuzzle) SetQueueFilter(v connection.QueueFilter) {
	k.socket.SetQueueFilter(v)
}

func (k *Kuzzle) SetQueueMaxSize(v int) {
	k.socket.SetQueueMaxSize(v)
}

func (k *Kuzzle) SetQueueTTL(v time.Duration) {
	k.socket.SetQueueTTL(v)
}

func (k *Kuzzle) SetReplayInterval(v time.Duration) {
	k.socket.SetReplayInterval(v)
}

func (k *Kuzzle) DefaultIndex() string {
	return k.defaultIndex
}

// SetDefaultIndex set the default data index. Has the same effect than the defaultIndex constructor option.
func (k *Kuzzle) SetDefaultIndex(index string) error {
	if index == "" {
		return types.NewError("Kuzzle.SetDefaultIndex: index required", 400)
	}

	k.defaultIndex = index
	return nil
}

func (k *Kuzzle) Volatile() types.VolatileData {
	return k.volatile
}

func (k *Kuzzle) SetVolatile(v types.VolatileData) {
	k.volatile = v
}

func (k *Kuzzle) EmitEvent(e int, arg interface{}) {
	k.socket.EmitEvent(e, arg)
}
