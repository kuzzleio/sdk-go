package event

// Consts for event
const (
	Connected = iota
	Discarded
	Disconnected
	LoginAttempt
	NetworkError
	OfflineQueuePop
	OfflineQueuePush
	QueryError
	Reconnected
	TokenExpired
	Error
	Done
)
