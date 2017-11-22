// Kuzzle Entry point and main struct for the entire SDK
package kuzzle

import (
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"sync"
	"time"
)

const version = "1.0.0"

type IKuzzle interface {
	Query(*types.KuzzleRequest, chan<- *types.KuzzleResponse, types.QueryOptions)
}

type Kuzzle struct {
	host   string
	port   int
	socket connection.Connection

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	defaultIndex   string
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory map[string]time.Time

	MemoryStorage *ms.Ms
	Security      *security.Security
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
		socket:  c,
		version: version,
	}

	k.MemoryStorage = &ms.Ms{k}
	k.Security = &security.Security{k}

	k.RequestHistory = k.socket.RequestHistory()

	k.defaultIndex = options.DefaultIndex()

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
		if k.lastUrl != k.host {
			k.wasConnected = false
			k.lastUrl = k.host
		}

		if wasConnected {
			if k.jwt != "" {
				go func() {
					res, err := k.CheckToken(k.jwt)

					if err != nil {
						k.jwt = ""
						k.socket.EmitEvent(event.JwtExpired, nil)
						return
					}

					if !res.Valid {
						k.jwt = ""
						k.socket.EmitEvent(event.JwtExpired, nil)
					}
				}()
			}
		}
		return nil
	}
	return types.NewError(err.Error())
}

func (k *Kuzzle) OfflineQueue() []*types.QueryObject {
	return k.socket.OfflineQueue()
}

func (k *Kuzzle) SetOfflineQueue(v []*types.QueryObject) {
	k.socket.SetOfflineQueue(v)
}

// Jwt get internal jwtToken used to request kuzzle.
func (k *Kuzzle) Jwt() string {
	return k.jwt
}

func (k *Kuzzle) SetJwt(token string) {
	k.jwt = token

	if token != "" {
		k.socket.RenewSubscriptions()
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
				room.Renew(room.Filters(), room.RealtimeChannel(), room.ResponseChannel())

				return true
			})

			return true
		})
	}
}

func (k *Kuzzle) RegisterRoom(roomId, id string, room types.IRoom) {
	k.socket.RegisterRoom(roomId, id, room)
}

func (k *Kuzzle) UnregisterRoom(roomId string) {
	k.socket.UnregisterRoom(roomId)
}

func (k *Kuzzle) State() int {
	return k.socket.State()
}
