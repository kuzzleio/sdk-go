package security

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

type User struct {
	Id         string
	Content    map[string]interface{}
	ProfileIds []string
	Security   *Security
}

type UserSearchResult struct {
	Hits     []*User
	Total    int
	ScrollId string
}

// Create the user in Kuzzle.
// Credentials can be created during the process by using setCredentials beforehand.
func (u *User) Create(options types.QueryOptions) (*User, error) {
	return u.persist("createUser", types.UserData{Content: u.Content, ProfileIds: u.ProfileIds}, options)
}

func (u *User) CreateCredentials(strategy string, credentials interface{}, options types.QueryOptions) (json.RawMessage, error) {
	return u.queryCredentials("createCredentials", strategy, credentials, options)
}

// CreateWithCredentials creates the user in Kuzzle along with supplied credentials
func (u *User) CreateWithCredentials(credentials types.Credentials, options types.QueryOptions) (*User, error) {
	type body struct {
		Content     map[string]interface{} `json:"content"`
		Credentials types.Credentials      `json:"credentials,omitempty"`
		ProfileIds  []string               `json:"profileIds"`
	}
	return u.persist("createUser", body{
		Content:     u.Content,
		Credentials: credentials,
		ProfileIds:  u.ProfileIds,
	}, options)
}

// Delete the user from Kuzzle.
func (u *User) Delete(options types.QueryOptions) (string, error) {
	return u.Security.rawDelete("deleteUser", u.Id, options)
}

// DeleteCredentials removes the user credentials for the given strategy
func (u *User) DeleteCredentials(strategy string, options types.QueryOptions) (bool, error) {
	_, err := u.queryCredentials("deleteCredentials", strategy, nil, options)
	return err != nil, err
}

// GetCredentialsInfo returns user credentials for the given strategy.
func (u *User) GetCredentialsInfo(strategy string, options types.QueryOptions) (json.RawMessage, error) {
	return u.queryCredentials("getCredentials", strategy, nil, options)
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
	if u.Id == "" {
		return nil, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         u.Id,
	}
	go u.Security.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []*types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}

// HasCredentials checks if the user has some credentials set for the given strategy.
func (u *User) HasCredentials(strategy string, options types.QueryOptions) (bool, error) {
	_, err := u.queryCredentials("hasCredentials", strategy, nil, options)
	return err != nil, err
}

// Replace the user in Kuzzle.
func (u *User) Replace(options types.QueryOptions) (*User, error) {
	return u.persist("replaceUser", u.getFlatBody(), options)
}

// SaveRestricted stores the current user as restricted into Kuzzle.
func (u *User) SaveRestricted(credentials types.Credentials, options types.QueryOptions) (*User, error) {
	if u.Id == "" {
		return nil, types.NewError("User.SaveRestricted: id is required", 400)
	}

	type body struct {
		Content     map[string]interface{} `json:"content"`
		Credentials types.Credentials      `json:"credentials,omitempty"`
	}

	return u.persist("createRestrictedUser", body{Content: u.Content, Credentials: credentials}, options)
}

// UpdateCredentials replaces the user credentials for the given strategy
// The credentials to send depends entirely on the authentication plugin and strategy you want to update credentials for.
func (u *User) UpdateCredentials(strategy string, credentials interface{}, options types.QueryOptions) (json.RawMessage, error) {
	return u.queryCredentials("updateCredentials", strategy, credentials, options)
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

func (u *User) queryCredentials(action string, strategy string, body interface{}, options types.QueryOptions) (json.RawMessage, error) {
	if strategy == "" {
		return nil, types.NewError("Security."+action+": strategy is required", 400)
	}

	if u.Id == "" {
		return nil, types.NewError("Security."+action+": user id is required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     action,
		Strategy:   strategy,
		Id:         u.Id,
		Body:       body,
	}
	go u.Security.Kuzzle.Query(query, options, result)
	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
