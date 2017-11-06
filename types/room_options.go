package types

type roomOptions struct {
	scope           string
	state           string
	users           string
	subscribeToSelf bool
	volatile        VolatileData
}

type RoomOptions interface {
	GetScope() string
	SetScope(string) *roomOptions
	GetState() string
	SetState(string) *roomOptions
	GetUsers() string
	SetUsers(string) *roomOptions
	GetSubscribeToSelf() bool
	SetSubscribeToSelf(bool) *roomOptions
	GetVolatile() VolatileData
	SetVolatile(VolatileData) *roomOptions
}

func (ro roomOptions) GetScope() string {
	return ro.scope
}

func (ro *roomOptions) SetScope(scope string) *roomOptions {
	ro.scope = scope
	return ro
}

func (ro roomOptions) GetState() string {
	return ro.state
}

func (ro *roomOptions) SetState(state string) *roomOptions {
	ro.state = state
	return ro
}

func (ro roomOptions) GetUsers() string {
	return ro.users
}

func (ro *roomOptions) SetUsers(users string) *roomOptions {
	ro.users = users
	return ro
}

func (ro roomOptions) GetSubscribeToSelf() bool {
	return ro.subscribeToSelf
}

func (ro *roomOptions) SetSubscribeToSelf(subscribeToSelf bool) *roomOptions {
	ro.subscribeToSelf = subscribeToSelf
	return ro
}

func (ro roomOptions) GetVolatile() VolatileData {
	return ro.volatile
}

func (ro *roomOptions) SetVolatile(volatile VolatileData) *roomOptions {
	ro.volatile = volatile
	return ro
}

func NewRoomOptions() RoomOptions {
	return &roomOptions{
		scope:           SCOPE_ALL,
		state:           STATE_DONE,
		users:           USERS_NONE,
		subscribeToSelf: true,
	}
}
