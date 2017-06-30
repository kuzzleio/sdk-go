package event

// Consts for event
const (
	Connected = iota
	Discarded = iota
	Disconnected = iota
	LoginAttempt = iota
	NetworkError = iota
	OfflineQueuePop = iota
	OfflieQueuePush = iota
	QueryError = iota
	Reconnected = iota
	JwtTokenExpired = iota
	Error = iota
)