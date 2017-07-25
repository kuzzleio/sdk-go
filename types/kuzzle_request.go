package types

type KuzzleRequest struct {
  RequestId string `json:"requestId"`
  Controller string `json:"controller"`
  Action string `json:"action"`
  Index string `json:"index"`
  Collection string `json:"collection"`
  Body interface{} `json:"body"`
  Id string `json:"_id"`
  From int `json:"from"`
  Size int `json:"size"`
  Scroll string `json:"scroll"`
  ScrollId string `json:"scrollId"`
}

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
