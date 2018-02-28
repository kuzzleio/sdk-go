package index

import "github.com/kuzzleio/sdk-go/types"

func (Index) SetAutoRefresh(index string, autoRefresh bool) error {
	if index == "" {
		return false, types.NewError("Index.SetAutoRefresh: index required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "index",
		Action:     "setAutoRefresh",
		Index:      index,
		Body: struct {
			AutoRefresh bool `json:"autoRefresh"`
		}{autoRefresh},
	}

	go k.Query(query, nil, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	return nil
}
