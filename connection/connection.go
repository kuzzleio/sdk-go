package connection

import "github.com/kuzzleio/sdk-go/types"

type Connection interface {
	AddListener(event int, channel chan<- interface{})
	Connect() (bool, error)
	Send([]byte, *types.Options, chan<- types.KuzzleResponse, string) error
	Close() error
	GetOfflineQueue() *[]types.QueryObject
	GetState() *int
	EmitEvent(int, interface{})
}
