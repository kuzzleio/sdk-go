package security

import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

type jsonProfile struct {
	Id string `json:"_id"`
	Source types.Policies `json:"_source"`
}

type jsonProfileSearchResult struct {
	Total int `json:"total"`
	Hits []*jsonProfile `json:"hits"`
	ScrollId string     `json:"scrollId"`
}

func (s *Security) jsonProfileToProfile(j *jsonProfile) (*Profile) {
	p := &Profile{
		Id: j.Id,
		Policies: j.Source.Policies,
	}
	p.Security = s

	return p
}

func ProfileToJson(p *Profile) ([]byte, error) {
	j := &jsonProfile{
		Id: p.Id,
		Source: types.Policies{
			Policies: p.Policies,
		},
	}

	return json.Marshal(j)
}
