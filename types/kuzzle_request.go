package types

type KuzzleRequest struct {
	RequestId  string        `json:"requestId,omitempty"`
	Controller string        `json:"controller,omitempty"`
	Action     string        `json:"action,omitempty"`
	Index      string        `json:"index,omitempty"`
	Collection string        `json:"collection,omitempty"`
	Body       interface{}   `json:"body"`
	Id         string        `json:"_id,omitempty"`
	From       int           `json:"from"`
	Size       int           `json:"size"`
	Scroll     string        `json:"scroll,omitempty"`
	ScrollId   string        `json:"scrollId,omitempty"`
	Strategy   string        `json:"strategy,omitempty"`
	ExpiresIn  int           `json:"expiresIn"`
	Volatile   VolatileData  `json:"volatile"`
	Scope      string        `json:"scope"`
	State      string        `json:"state"`
	User       string        `json:"user"`
	Start      int           `json:"start,omitempty"`
	Stop       int           `json:"stop,omitempty"`
	End        int           `json:"end,omitempty"`
	Bit        int           `json:"bit,omitempty"`
	Member     string        `json:"member,omitempty"`
	Member1    string        `json:"member1,omitempty"`
	Member2    string        `json:"member2,omitempty"`
	Members    []string      `json:"members,omitempty"`
	Lon        float64       `json:"lon,omitempty"`
	Lat        float64       `json:"lat,omitempty"`
	Distance   float64       `json:"distance,omitempty"`
	Unit       string        `json:"unit,omitempty"`
	Options    []interface{} `json:"options,omitempty"`
	Keys       []string      `json:"keys,omitempty"`
	Cursor     int           `json:"cursor,omitempty"`
	Offset     int           `json:"offset,omitempty"`
	Field      string        `json:"field,omitempty"`
	Fields     []string      `json:"fields,omitempty"`
	Subcommand string        `json:"subcommand,omitempty"`
	Pattern    string        `json:"pattern,omitempty"`
	Idx        int           `json:"idx, omitempty"`
	Min        string        `json:"min,omitempty"`
	Max        string        `json:"max,omitempty"`
	Limit      string        `json:"limit,omitempty"`
	Count      int           `json:"count,omitempty"`
	Match      string        `json:"match,omitempty"`
}

type SubscribeQuery struct {
	Scope string      `json:"scope"`
	State string      `json:"state"`
	User  string      `json:"user"`
	Body  interface{} `json:"body"`
}

type VolatileData map[string]interface{}

type UserCredentials map[string]interface{}

type UserData struct {
	ProfileIds  []string               `json:"profileIds"`
	Content     map[string]interface{} `json:"content"`
	Credentials UserCredentials        `json:"credentials"`
}

type PolicyRestriction struct {
	Index       string   `json:"index"`
	Collections []string `json:"collections,omitempty"`
}

type Policy struct {
	RoleId             string               `json:"roleId"`
	RestrictedTo       []*PolicyRestriction `json:"restrictedTo,omitempty"`
	AllowInternalIndex bool                 `json:"allowInternalIndex,omitempty"`
}

type Policies struct {
	Policies []*Policy `json:"policies"`
}

type GeoPoint struct {
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Name string  `json:"name"`
}

type MsHashField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type MSKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MSSortedSet struct {
	Score  float64 `json:"score"`
	Member string  `json:"member"`
}

type SearchFilters struct {
	Query        interface{}         `json:"query,omitempty"`
	Sort         []map[string]string `json:"sort,omitempty"`
	Aggregations interface{}         `json:"aggregations,omitempty"`
	SearchAfter  []interface{}       `json:"search_after,omitempty"`
}
