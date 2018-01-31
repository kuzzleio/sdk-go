package kuzzle

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/state"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

// Query this is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query *types.KuzzleRequest, options types.QueryOptions, responseChannel chan<- *types.KuzzleResponse) {
	if k.State() == state.Disconnected || k.State() == state.Offline || k.State() == state.Ready {
		responseChannel <- &types.KuzzleResponse{Error: types.NewError("This Kuzzle object has been invalidated. Did you try to access it after a disconnect call?", 400)}
		return
	}

	u, _ := uuid.NewV4()
	requestId := u.String()

	if query.RequestId == "" {
		query.RequestId = requestId
	}

	type body struct{}

	if query.Body == nil {
		query.Body = make(map[string]interface{})
	}

	if options == nil {
		options = types.NewQueryOptions()
	}

	if options.Volatile() != nil {
		query.Volatile = options.Volatile()
		query.Volatile["sdkVersion"] = version
	} else {
		query.Volatile = types.VolatileData{"sdkVersion": version}
	}

	jsonRequest, _ := json.Marshal(query)
	out := map[string]interface{}{}
	json.Unmarshal(jsonRequest, &out)

	refresh := options.Refresh()
	if refresh != "" {
		out["refresh"] = refresh
	}

	out["from"] = options.From()
	out["size"] = options.Size()

	scroll := options.Scroll()
	if scroll != "" {
		out["scroll"] = scroll
	}

	scrollId := options.ScrollId()
	if scrollId != "" {
		out["scrollId"] = scrollId
	}

	retryOnConflict := options.RetryOnConflict()
	if retryOnConflict > 0 {
		out["retryOnConflict"] = retryOnConflict
	}

	finalRequest, err := json.Marshal(out)

	if err != nil {
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Error: types.NewError(err.Error())}
		}
		return
	}

	err = k.socket.Send(finalRequest, options, responseChannel, requestId)
	if err != nil {
		if responseChannel != nil {
			responseChannel <- &types.KuzzleResponse{Error: types.NewError(err.Error())}
		}
		return
	}
}
