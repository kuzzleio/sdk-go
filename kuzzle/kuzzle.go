// Kuzzle Entry point and main struct for the entire SDK
package kuzzle

import (
	"errors"
	"time"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/event"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/ms"
	"github.com/kuzzleio/sdk-go/security"
)

const version = "1.0.0"

type Kuzzle struct {
	Host   string
	socket connection.Connection
	State  int

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	defaultIndex   string
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory map[string]time.Time

	Ms             *ms.Ms
	Security       *security.Security
}

// New is the Kuzzle constructor
func NewKuzzle(c connection.Connection, options types.Options) (*Kuzzle, error) {
	if c == nil {
		return nil, errors.New("Connection is nil")
	}

	if options == nil {
		options = types.NewOptions()
	}

	k := &Kuzzle{
		socket:  c,
		headers: options.GetHeaders(),
		version: version,
	}

	k.Ms = &ms.Ms{k}
	k.Security = &security.Security{k}

	k.RequestHistory = k.socket.GetRequestHistory()

	headers := options.GetHeaders()
	if headers != nil {
		k.headers = headers
	}

	k.State = k.socket.GetState()

	k.defaultIndex = options.GetDefaultIndex()

	var err error

	if options.GetConnect() == types.Auto {
		err = k.Connect()
	}

	return k, err
}

// Connect connects to a Kuzzle instance using the provided host and port.
func (k Kuzzle) Connect() error {
	wasConnected, err := k.socket.Connect()
	if err == nil {
		if k.lastUrl != k.Host {
			k.wasConnected = false
			k.lastUrl = k.Host
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

	return err
}

func (k Kuzzle) GetOfflineQueue() *[]*types.QueryObject {
	return k.socket.GetOfflineQueue()
}

// GetJwt get internal jwtToken used to request kuzzle.
func (k Kuzzle) GetJwt() string {
	return k.jwt
}

func (k Kuzzle) RegisterRoom(roomId, id string, room types.IRoom) {
	k.socket.RegisterRoom(roomId, id, room)
}

func (k Kuzzle) UnregisterRoom(roomId string) {
	k.socket.UnregisterRoom(roomId)
}
