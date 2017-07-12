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
  // k.State = &k.socket.State
  if options.Connect == types.Auto {
    err = k.Connect()
  }

  return k, err
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func AddListener(k Kuzzle, event int, channel chan<- interface{}) {
  // k.socket.AddListener(event, channel)
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

// Instantiates a new collection object.
func (k *Kuzzle) Collection(collection, index string) *Collection {
  return NewCollection(k, collection, index)
}

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query types.KuzzleRequest, options *types.Options, responseChannel chan<- types.KuzzleResponse) {
  requestId := uuid.NewV4().String()

  query.RequestId = requestId

  type body struct{}
  if query.Body == nil {
    query.Body = &body{}
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

// Disconnect from Kuzzle and invalidate this instance.
// Does not fire a disconnected event.
func (k *Kuzzle) Disconnect() error {
  err := k.socket.Close()

  if err != nil {
    return err
  }
  k.wasConnected = false

  return nil
}

func buildQuery(collection, index, controller, action string, body interface{}) types.KuzzleRequest {
  return types.KuzzleRequest{
    Controller: controller,
    Action:     action,
    Index:      index,
    Collection: collection,
    Body:       body,
  }
}