package kuzzle

import (
  "github.com/kuzzleio/sdk-go/connection"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/satori/go.uuid"
  "sync"
)

type IKuzzle interface {
  Query(types.KuzzleRequest, chan<- types.KuzzleResponse, *types.Options)
}

type Kuzzle struct {
  Host   string
  socket connection.Connection
  State  *int

  wasConnected bool
  lastUrl      string
  message      chan []byte
  mu           *sync.Mutex
  jwtToken     string
}

// Kuzzle constructor
func NewKuzzle(c connection.Connection, options *types.Options) (*Kuzzle, error) {
  var err error

  if options == nil {
    options = types.DefaultOptions()
  }

  k := &Kuzzle{
    mu:     &sync.Mutex{},
    socket: c,
  }

  k.State = k.socket.GetState()

  if options.Connect == types.Auto {
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
      if k.jwtToken != "" {
        // todo avoid import cycle (kuzzle)
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
    }
    return nil
  }

  return err
}

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k Kuzzle) Query(query types.KuzzleRequest, options *types.Options, responseChannel chan<- types.KuzzleResponse) {
  requestId := uuid.NewV4().String()

  query.RequestId = requestId

  type body struct{}

  if query.Body == nil {
    query.Body = make(map[string]interface{})
  }

  jsonRequest, err := json.Marshal(query)
  if err != nil {
    responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
    return
  }

  err = k.socket.Send(jsonRequest, options, responseChannel, requestId)
  if err != nil {
    responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
    return
  }
}

func (k Kuzzle) GetOfflineQueue() *[]types.QueryObject {
  return k.socket.GetOfflineQueue()
}
