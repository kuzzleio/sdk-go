package connection

import "github.com/kuzzleio/sdk-go/types"

type Connection interface {
	Connect() (bool, error)
	Send([]byte, *types.Options, chan<- types.KuzzleResponse, string) error
	Close() error
	GetState() *int
}