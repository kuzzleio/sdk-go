package kuzzle

import (
	"errors"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/types"
	"sync"
	"time"
)

const version = "0.1"

type IKuzzle interface {
	Query(types.KuzzleRequest, chan<- types.KuzzleResponse, types.QueryOptions)
}

type Kuzzle struct {
	Host   string
	socket connection.Connection
	State  *int

	wasConnected   bool
	lastUrl        string
	message        chan []byte
	mu             *sync.Mutex
	defaultIndex   string
	jwt            string
	headers        map[string]interface{}
	version        string
	RequestHistory *map[string]time.Time
}

// Kuzzle constructor
func NewKuzzle(c connection.Connection, options types.Options) (*Kuzzle, error) {
	var err error

	if c == nil {
		return nil, errors.New("Connection is nil")
	}

	if options == nil {
		options = types.NewOptions()
	}

	k := &Kuzzle{
		mu:      &sync.Mutex{},
		socket:  c,
		headers: options.GetHeaders(),
		version: version,
	}
	k.RequestHistory = k.socket.GetRequestHistory()

	headers := options.GetHeaders()
	if headers != nil {
		k.headers = headers
	}

	k.State = k.socket.GetState()

	k.defaultIndex = options.GetDefaultIndex()

	if options.GetConnect() == types.Auto {
		err = k.Connect()
	}

	return k, err
}

// Connects to a Kuzzle instance using the provided host and port.
func (k *Kuzzle) Connect() error {
	wasConnected, err := k.socket.Connect()
	if err == nil {
		//if k.lastUrl != k.Host {
		//  k.wasConnected = false
		//  k.lastUrl = k.Host
		//}

		if wasConnected {
			if k.jwt != "" {
				// todo avoid import cycle (kuzzle)
				//go func() {
				//	res, err := kuzzle.CheckToken(k, k.jwt)
				//
				//	if err != nil {
				//		k.jwt = ""
				//		k.emitEvent(event.jwtExpired, nil)
				//		k.Reconnect()
				//		return
				//	}
				//
				//	if !res.Valid {
				//		k.jwt = ""
				//		k.emitEvent(event.jwtExpired, nil)
				//	}
				//	k.Reconnect()
				//}()
			}
		}
		return nil
	}

	return err
}

func (k Kuzzle) GetOfflineQueue() *[]types.QueryObject {
	return k.socket.GetOfflineQueue()
}

func (k Kuzzle) GetJwt() string {
	return k.jwt
}

func (k Kuzzle) RegisterRoom(roomId, id string, room types.IRoom) {
	k.socket.RegisterRoom(roomId, id, room)
}

func (k Kuzzle) UnregisterRoom(roomId string) {
	k.socket.UnregisterRoom(roomId)
}
func (k Kuzzle) GetHeaders() map[string]interface{} {
	return k.headers
}
