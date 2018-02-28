package server

import (
	"github.com/kuzzleio/sdk-go/types"
)

type Server struct {
	k types.IKuzzle
}

func NewServer(k types.IKuzzle) *Server {
	return &Server{k}
}
