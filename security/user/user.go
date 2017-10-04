package user

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security/profile"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityUser struct {
	Kuzzle *kuzzle.Kuzzle
}

type User struct {
	Id     string            `json:"_id"`
	Source json.RawMessage   `json:"_source"`
	Meta   *types.KuzzleMeta `json:"_meta"`
	SU     *SecurityUser
}

type UserSearchResult struct {
	Hits     []*User `json:"hits"`
	Total    int     `json:"total"`
	ScrollId string  `json:"scrollId"`
}

// SetContent updates the Source of the current user
// Updating an user will have no impact until the create or replace method is called.
func (u *User) SetContent(data *types.UserData) *User {
	u.Source, _ = json.Marshal(data)

	return u
}

// SetCredentials update the userData of the current user
// Updating user credentials will have no impact until the create method is called.
// The credentials to send depends entirely on the authentication plugin and strategy you want to create credentials for.
func (u User) SetCredentials(credentials *types.UserCredentials) User {
	userData := u.UserData()
	userData.Credentials = credentials

	return u.SetContent(userData)
}

// AddProfile adds a profile to the current user
// Updating an user will have no impact until the create or replace method is called.
func (u User) AddProfile(profile *profile.Profile) *User {
	userData := u.UserData()
	userData.ProfileIds = append(userData.ProfileIds, profile.Id)

	return u.SetContent(userData)
}

// GetProfiles returns the associated Profile instances from the Kuzzle API, using the profile identifiers attached to this user (see getProfileIds).
func (u User) GetProfiles(options types.QueryOptions) ([]*profile.Profile, error) {
	fetchedProfiles := []*profile.Profile{}

	if len(u.UserData().ProfileIds) == 0 {
		return fetchedProfiles, nil
	}

	for _, profileId := range u.UserData().ProfileIds {
		p, err := profile.SecurityProfile{Kuzzle: u.SU.Kuzzle}.Fetch(profileId, options)

		if err != nil {
			return []*profile.Profile{}, err
		}

		fetchedProfiles = append(fetchedProfiles, p)
	}

	return fetchedProfiles, nil
}

// GetProfileIds returns the list of profile identifiers associated to this user.
func (u User) GetProfileIds() []string {
	return u.UserData().ProfileIds
}

// SetProfiles updates the profiles of the current user
// Updating a user will have no impact until the create or replace method is called.
func (u User) SetProfiles(profiles []*profile.Profile) *User {
	profileIds := make([]string{}, profiles.len())

	userData := u.UserData()

	for _, p := range profiles {
		profileIds = append(profileIds, p.Id)
	}
	userData.ProfileIds = profileIds

	return u.SetContent(userData)
}

// Create the user in kuzzle. Credentials can be created during the process by using setCredentials beforehand.
func (u User) Create(options types.QueryOptions) (*User, error) {
	return u.SU.Create(u.Id, u.UserData(), options)
}

// SaveRestricted stores the current user as restricted into Kuzzle.
func (u User) SaveRestricted(options types.QueryOptions) (*User, error) {
	return u.SU.CreateRestrictedUser(u.Id, u.UserData(), options)
}

// Replace the user in kuzzle.
func (u User) Replace(options types.QueryOptions) (*User, error) {
	return u.SU.Replace(u.Id, u.UserData(), options)
}

// Update the user in kuzzle.
func (u User) Update(content *types.UserData, options types.QueryOptions) (User, error) {
	return u.SU.Update(u.Id, content, options)
}

// Delete the user in Kuzzle.
func (u User) Delete(options types.QueryOptions) (string, error) {
	return u.SU.Delete(u.Id, options)
}

// UserData returns the current user's data
func (u User) UserData() *types.UserData {
	userData := *types.UserData{}
	json.Unmarshal(u.Source, userData)

	rawContent := map[string]interface{}{}
	json.Unmarshal(u.Source, &rawContent)

	for key, value := range rawContent {
		if key != "profileIds" && key != "credentials" && value != nil {
			if userData.Content == nil {
				userData.Content = make(map[string]interface{})
			}
			userData.Content[key] = value
		}
	}

	return userData
}

// Content returns the current user's content
func (u User) Content(key string) interface{} {
	return u.UserData().Content[key]
}

// ContentMap returns the current user's content map
func (u User) ContentMap(keys ...string) map[string]interface{} {
	values := make(map[string]interface{})

	for _, key := range keys {
		values[key] = u.UserData().Content[key]
	}

	return values
}

// Fetch retrieves an User using its provided unique id.
func (su *SecurityUser) Fetch(id string, options types.QueryOptions) (*User, error) {
	if id == "" {
		return &User{}, errors.New("Security.User.Fetch: user id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUser",
		Id:         id,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return User{}, errors.New(res.Error.Message)
	}

	u := &User{SU: su}
	json.Unmarshal(res.Result, u)

	return u, nil
}

// Search executes a search on Users according to filters.
func (su SecurityUser) Search(filters interface{}, options types.QueryOptions) (*UserSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return &UserSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := &UserSearchResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

// Scroll executes a scroll search on Users.
func (su SecurityUser) Scroll(scrollId string, options types.QueryOptions) (*UserSearchResult, error) {
	if scrollId == "" {
		return &UserSearchResult{}, errors.New("Security.User.Scroll: scroll id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollUsers",
		ScrollId:   scrollId,
	}

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return &UserSearchResult{}, errors.New(res.Error.Message)
	}

	searchResult := &UserSearchResult{}
	json.Unmarshal(res.Result, searchResult)

	return searchResult, nil
}

// Create a new User in Kuzzle.
func (su *SecurityUser) Create(kuid string, content *types.UserData, options types.QueryOptions) (*User, error) {
	if kuid == "" {
		return &User{}, errors.New("Security.User.Create: user kuid required")
	}

	ch := make(chan *types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}
	type createBody struct {
		Content     *userData              `json:"content"`
		Credentials *types.UserCredentials `json:"credentials"`
	}

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createUser",
		Body:       createBody{Content: ud, Credentials: content.Credentials},
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return &User{}, errors.New(res.Error.Message)
	}

	user := &User{SU: su}
	json.Unmarshal(res.Result, user)

	return user, nil
}

// CreateRestrictedUser creates a new restricted User in Kuzzle.
func (su *SecurityUser) CreateRestrictedUser(kuid string, content *types.UserData, options types.QueryOptions) (*User, error) {
	if kuid == "" {
		return &User{}, errors.New("Security.User.CreateRestrictedUser: user kuid required")
	}

	ch := make(chan *types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	for key, value := range content.Content {
		ud[key] = value
	}
	type createBody struct {
		Content     *userData              `json:"content"`
		Credentials *types.UserCredentials `json:"credentials"`
	}

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "createRestrictedUser",
		Body:       createBody{Content: ud, Credentials: content.Credentials},
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return User{}, errors.New(res.Error.Message)
	}

	user := &User{SU: su}
	json.Unmarshal(res.Result, user)

	return user, nil
}

// Replace an User in Kuzzle.
func (su *SecurityUser) Replace(kuid string, content *types.UserData, options types.QueryOptions) (*User, error) {
	if kuid == "" {
		return &User{}, errors.New("Security.User.Replace: user kuid required")
	}

	ch := make(chan *types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "replaceUser",
		Body:       ud,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return &User{}, errors.New(res.Error.Message)
	}

	user := &User{SU: su}
	json.Unmarshal(res.Result, user)

	return user, nil
}

// Update an User in Kuzzle.
func (su *SecurityUser) Update(kuid string, content *types.UserData, options types.QueryOptions) (*User, error) {
	if kuid == "" {
		return &User{}, errors.New("Security.User.Update: user kuid required")
	}

	ch := make(chan *types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "updateUser",
		Body:       ud,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return &User{}, errors.New(res.Error.Message)
	}

	user := &User{SU: su}
	json.Unmarshal(res.Result, user)

	return user, nil
}

// Delete an User in Kuzzle.
// There is a small delay between user deletion and their deletion in our advanced search layer, usually a couple of seconds.
// This means that a user that has just been deleted will still be returned by this function.
func (su SecurityUser) Delete(kuid string, options types.QueryOptions) (string, error) {
	if kuid == "" {
		return "", errors.New("Security.User.Delete: user kuid required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteUser",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	shardResponse := types.ShardResponse{}
	json.Unmarshal(res.Result, &shardResponse)

	return shardResponse.Id, nil
}

// GetRights returns the rights of an User using its provided unique id.
func (su SecurityUser) GetRights(kuid string, options types.QueryOptions) ([]*types.UserRights, error) {
	if kuid == "" {
		return []*types.UserRights{}, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return []*types.UserRights{}, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []*types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}

// IsActionAllowed indicates whether an action is allowed, denied or conditional based on user rights provided as the first argument.
// An action is defined as a couple of action and controller (mandatory), plus an index and a collection(optional).
func (su SecurityUser) IsActionAllowed(rights []*types.UserRights, controller string, action string, index string, collection string) (string, error) {
	if rights == nil {
		return "", errors.New("Security.User.IsActionAllowed: Rights parameter is mandatory")
	}
	if controller == "" {
		return "", errors.New("Security.User.IsActionAllowed: Controller parameter is mandatory")
	}
	if action == "" {
		return "", errors.New("Security.User.IsActionAllowed: Action parameter is mandatory")
	}

	filteredUserRights := make([]*types.UserRights{}, rights.len())

	for _, ur := range rights {
		if (ur.Controller == controller || ur.Controller == "*") && (ur.Action == action || ur.Action == "*") && (ur.Index == index || ur.Index == "*") && (ur.Collection == collection || ur.Collection == "*") {
			filteredUserRights = append(filteredUserRights, ur)
		}
	}

	for _, ur := range filteredUserRights {
		if ur.Value == "allowed" {
			return "allowed", nil
		}
		if ur.Value == "conditional" {
			return "conditional", nil
		}
	}

	return "denied", nil
}
