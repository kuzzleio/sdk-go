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

