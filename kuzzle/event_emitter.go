package kuzzle

// KuzzleEventEmitter Is an interface used by Kuzzle and Room instances to emit and listen to events (see Event Handling section).
type KuzzleEventEmitter interface {
	AddListener(event int, channel chan<- interface{})
	On(event int, channel chan<- interface{})
	RemoveAllListeners(event int)
	RemoveListener(event int, channel chan<- interface{})
	Once(event int, channel chan<- interface{})
	ListenerCount(event int) int
}
