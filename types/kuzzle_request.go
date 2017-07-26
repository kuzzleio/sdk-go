package types

type KuzzleRequest struct {
	RequestId  string       `json:"requestId,omitempty"`
	Controller string       `json:"controller,omitempty"`
	Action     string       `json:"action,omitempty"`
	Index      string       `json:"index,omitempty"`
	Collection string       `json:"collection,omitempty"`
	Body       interface{}  `json:"body"`
	Id         string       `json:"_id,omitempty"`
	From       int          `json:"from"`
	Size       int          `json:"size"`
	Scroll     string       `json:"scroll,omitempty"`
	ScrollId   string       `json:"scrollId,omitempty"`
	Strategy   string       `json:"strategy,omitempty"`
	ExpiresIn  int          `json:"expiresIn"`
	Volatile   VolatileData `json:"volatile"`
}

type VolatileData map[string]interface{}

type UserCredentials map[string]struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	ProfileIds []string `json:"profileIds"`
	Content map[string]interface{} `json:"content"`
	Credentials UserCredentials `json:"credentials"`
}

type Policy struct {
	RoleId string `json:"roleId"`
}

type Policies struct {
	Policies []Policy `json:"policies"`
}
