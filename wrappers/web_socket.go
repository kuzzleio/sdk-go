package wrappers

import (
  "github.com/gorilla/websocket"
  "flag"
  "log"
  "net/url"
  "sync"
)

type WebSocket struct {
  ws  *websocket.Conn
  mu  *sync.Mutex
}

func NewWebSocket() *WebSocket {
  return &WebSocket{
    mu: &sync.Mutex{},
  }
}

func (ws *WebSocket) Connect(connectUrl string) (chan []byte, error) {
  addr := flag.String("addr", connectUrl, "http service address")
  u := url.URL{Scheme: "ws", Host: *addr}
  resChan := make(chan []byte)
  socket, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

  if err != nil {
    return nil, err
  }

  ws.ws = socket

  go func() {
    for {
      _, message, err := ws.ws.ReadMessage()

      if err != nil {
        log.Fatal("read", err)
        close(resChan)
        return
      }
      go func() {
        resChan <- message
      }()
    }
  }()

  return resChan, err
}

func (ws WebSocket) Send(query []byte) error {
  ws.mu.Lock()
  defer ws.mu.Unlock()
  return ws.ws.WriteMessage(websocket.TextMessage, query)
}

func (ws *WebSocket) Close() error {
  ws.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
  return ws.ws.Close()
}