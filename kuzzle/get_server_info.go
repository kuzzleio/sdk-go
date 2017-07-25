package kuzzle

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

func (k Kuzzle) GetServerInfo(options *types.Options) (json.RawMessage, error) {
	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "server",
		Action:     "info",
	}

	go k.Query(query, options, result)

	res := <-result

	if res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	type serverInfo struct {
		ServerInfo json.RawMessage `json:"serverInfo"`
	}
	info := serverInfo{}
	json.Unmarshal(res.Result, &info)

	return info.ServerInfo, nil
}
