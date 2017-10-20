package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type User struct {
	Id     string
	Content map[string]interface{}
	ProfileIds []string
	Security *Security
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
	u.ProfileIds = append(u.ProfileIds, profile.Id)
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

// Replace the user in Kuzzle.
func (u *User) Replace(options types.QueryOptions) (*User, error) {
	body := make(map[string]interface{})
	for k, v := range u.Content {
		body[k] = v
	}
	if u.ProfileIds != nil {
		body["profileIds"] = u.ProfileIds
	}

	return u.persist("replaceUser", body, options)
}

// SaveRestricted stores the current user as restricted into Kuzzle.
func (u *User) SaveRestricted(options types.QueryOptions) (*User, error) {
	if u.Id == "" {
		return &User{}, errors.New("Security.User.SaveRestricted: user kuid required")
	}

	type body struct {
		content map[string]interface{} `json:"content"`
		credentials types.UserCredentials `json:"credentials,omitempty"`
	}

	return u.persist("createRestrictedUser", body{content: u.Content, credentials: u.credentials}, options)
}

func (u *User) SetCredentials(credentials types.UserCredentials) *User {
	u.credentials = credentials
	return u
}

// SetProfiles updates the profiles of the current user
// Updating a user will have no impact until the create or replace method is called.
func (u *User) SetProfiles(profiles []*Profile) *User {
	u.ProfileIds = make([]string, 0, len(profiles))

	for _, p := range profiles {
		u.ProfileIds = append(u.ProfileIds, p.Id)
	}

	return u
}

// Update the user in kuzzle.
func (u *User) Update(content *types.UserData, options types.QueryOptions) (*User, error) {
	body := make(map[string]interface{})
	for k, v := range u.Content {
		body[k] = v
	}
	for k, v := range content.Content {
		body[k] = v
	}

	profileIdsMap := make(map[string]bool)
	if u.ProfileIds != nil {
		for _, profileId := range u.ProfileIds {
			profileIdsMap[profileId] = true
		}
	}
	if content.ProfileIds != nil {
		for _, profileId := range content.ProfileIds {
			profileIdsMap[profileId] = true
		}
	}
	if len(profileIdsMap) > 0 {
		profileIds := []string{}
		for profileId := range profileIdsMap {
			profileIds = append(profileIds, profileId)
		}
		body["profileIds"] = profileIds
	}

	return u.persist("updateUser", body, options)
}

func (u *User) persist(action string, body interface{}, options types.QueryOptions) (*User, error) {
	if options == nil {
		options = types.NewQueryOptions()
	}

	if action != "createUser" && u.Id == "" {
		return nil, errors.New("User." + action + ": id is required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action: action,
		Body: body,
		Id: u.Id,
	}
	go u.Security.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	jsonUser := &jsonUser{}
	json.Unmarshal(res.Result, jsonUser)

	updatedUser := u.Security.jsonUserToUser(jsonUser)

	u.Id = updatedUser.Id
	u.Content = updatedUser.Content
	u.ProfileIds = updatedUser.ProfileIds

	return u, nil
}
