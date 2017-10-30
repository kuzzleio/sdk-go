package security

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

type User struct {
	Id          string
	Content     map[string]interface{}
	ProfileIds  []string
	Security    *Security
	credentials types.UserCredentials
}

type UserSearchResult struct {
	Hits     []*User
	Total    int
	ScrollId string
}

// AddProfile adds a profile to the current user
// Updating an user will have no impact until the create or replace method is called.
func (u *User) AddProfile(profile *Profile) *User {
	u.addProfileIds(profile.Id)
	return u
}

// Create the user in Kuzzle.
// Credentials can be created during the process by using setCredentials beforehand.
func (u *User) Create(options types.QueryOptions) (*User, error) {
	return u.persist("createUser", types.UserData{Content: u.Content, Credentials: u.credentials, ProfileIds: u.ProfileIds}, options)
}

// Delete the user from Kuzzle.
func (u *User) Delete(options types.QueryOptions) (string, error) {
	return u.Security.rawDelete("deleteUser", u.Id, options)
}

// GetCredentials returns user credentials for the given strategy.
func (u *User) GetCredentials(strategy string, options types.QueryOptions) (json.RawMessage, error) {
	return u.Security.GetCredentials(strategy, u.Id, options)
}

// GetProfiles returns the associated Profile instances from the Kuzzle API, using the profile identifiers attached to this user (see getProfileIds).
func (u *User) GetProfiles(options types.QueryOptions) ([]*Profile, error) {
	fetchedProfiles := []*Profile{}

	if len(u.ProfileIds) == 0 {
		return fetchedProfiles, nil
	}

	for _, profileId := range u.ProfileIds {
		p, err := u.Security.FetchProfile(profileId, options)

		if err != nil {
			return nil, err
		}

		fetchedProfiles = append(fetchedProfiles, p)
	}

	return fetchedProfiles, nil
}

// GetRights returns user permissions the user is granted, per controller action
func (u *User) GetRights(options types.QueryOptions) ([]*types.UserRights, error) {
	return u.Security.GetUserRights(u.Id, options)
}

// Replace the user in Kuzzle.
func (u *User) Replace(options types.QueryOptions) (*User, error) {
	return u.persist("replaceUser", u.getFlatBody(), options)
}

// SaveRestricted stores the current user as restricted into Kuzzle.
func (u *User) SaveRestricted(options types.QueryOptions) (*User, error) {
	if u.Id == "" {
		return nil, types.NewError("User.SaveRestricted: id is required", 400)
	}

	type body struct {
		Content     map[string]interface{} `json:"content"`
		Credentials types.UserCredentials  `json:"credentials,omitempty"`
	}

	return u.persist("createRestrictedUser", body{Content: u.Content, Credentials: u.credentials}, options)
}

// SetCredentials adds credentials to user before creating it
// Updating user credentials will have no impact until the create method is called.
// The credentials to send depends entirely on the authentication plugin and strategy you want to create credentials for.
func (u *User) SetCredentials(strategy string, credentials interface{}) *User {
	if u.credentials == nil {
		u.credentials = types.UserCredentials{}
	}
	u.credentials[strategy] = credentials
	return u
}

// SetProfiles updates the profiles of the current user
// Updating a user will have no impact until the create or replace method is called.
func (u *User) SetProfiles(profiles []*Profile) *User {
	u.ProfileIds = make([]string, 0, len(profiles))

	profileIds := []string{}
	for _, p := range profiles {
		profileIds = append(profileIds, p.Id)
	}
	u.addProfileIds(profileIds...)

	return u
}

// Update the user in kuzzle.
func (u *User) Update(content *types.UserData, options types.QueryOptions) (*User, error) {
	if content != nil {
		for k, v := range content.Content {
			if u.Content == nil {
				u.Content = make(map[string]interface{})
			}
			u.Content[k] = v
		}

		u.addProfileIds(content.ProfileIds...)
	}

	return u.persist("updateUser", u.getFlatBody(), options)
}

func (u *User) addProfileIds(profileIds ...string) {
	if u.ProfileIds == nil {
		u.ProfileIds = []string{}
	}

	profileIdsMap := make(map[string]bool)
	for _, profileId := range u.ProfileIds {
		profileIdsMap[profileId] = true
	}

	for _, profileId := range profileIds {
		_, alreadyIn := profileIdsMap[profileId]
		if !alreadyIn {
			u.ProfileIds = append(u.ProfileIds, profileId)
		}
	}
}

func (u *User) getFlatBody() map[string]interface{} {
	body := make(map[string]interface{})
	for k, v := range u.Content {
		body[k] = v
	}
	if u.ProfileIds != nil {
		body["profileIds"] = u.ProfileIds
	}

	return body
}

func (u *User) persist(action string, body interface{}, options types.QueryOptions) (*User, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	if action != "createUser" && u.Id == "" {
		return nil, types.NewError("User."+action+": id is required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Body:       body,
		Id:         u.Id,
	}
	go u.Security.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res.Result, jsonUser)

	updatedUser := u.Security.jsonUserToUser(jsonUser)

	u.Id = updatedUser.Id
	u.Content = updatedUser.Content
	u.ProfileIds = updatedUser.ProfileIds

	return u, nil
}
