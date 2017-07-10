package event

// Consts for event
const (
	Connected = iota
	Discarded
	Disconnected
	LoginAttempt
	NetworkError
	OfflineQueuePop
	OfflieQueuePush
	QueryError
	Reconnected
	JwtTokenExpired
	Error
)