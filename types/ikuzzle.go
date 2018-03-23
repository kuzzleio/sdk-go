package types

type IKuzzle interface {
	Query(query *KuzzleRequest, options QueryOptions, responseChannel chan<- *KuzzleResponse)
	EmitEvent(int, interface{})
	SetJwt(string)
	RegisterSub(string, string, chan<- KuzzleNotification)
	UnregisterSub(string)
	AddListener(event int, notifChan chan<- KuzzleNotification)
	AutoResubscribe() bool
}
