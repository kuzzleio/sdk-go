package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// MGetProfiles deletes all roles matching with given ids
func (s *Security) MGetProfiles(ids []string, options types.QueryOptions) ([]*Profile, error) {
	if len(ids) == 0 {
		return nil, types.NewError("Security.MGetProfiles: ids array can't be nil", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "mGetProfiles",
		Body: struct {
			Ids []string `json:"ids"`
		}{ids},
	}
	go s.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var fetchedRaw jsonProfileSearchResult
	var fetchedProfiles []*Profile
	json.Unmarshal(res.Result, &fetchedRaw)

	for _, jsonProfileRaw := range fetchedRaw.Hits {
		fetchedProfiles = append(fetchedProfiles, jsonProfileRaw.jsonProfileToProfile())
	}

	return fetchedProfiles, nil

}
