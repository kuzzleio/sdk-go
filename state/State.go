package state

const (
	Connecting = iota
	Disconnected = iota
	Connected = iota
	Initializing = iota
	Ready = iota
	Logged_out = iota
	Error = iota
	Offline = iota
)