package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query types.KuzzleRequest, options types.QueryOptions, responseChannel chan<- types.KuzzleResponse) {
	requestId := uuid.NewV4().String()

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

	volatile := options.GetVolatile()
	if options.GetVolatile() != nil {
		query.Volatile = volatile
		query.Volatile["sdkVersion"] = version
	} else {
		query.Volatile = types.VolatileData{"sdkVersion": version}
	}

	jsonRequest, _ := json.Marshal(query)
	out := map[string]interface{}{}
	json.Unmarshal(jsonRequest, &out)
	k.addHeaders(&out, query)

	refresh := options.GetRefresh()
	if refresh != "" {
		out["refresh"] = refresh
	}

	out["from"] = options.GetFrom()
	out["size"] = options.GetSize()

	scroll := options.GetScroll()
	if scroll != "" {
		out["scroll"] = scroll
	}

	scrollId := options.GetScrollId()
	if scrollId != "" {
		out["scrollId"] = scrollId
	}

	retryOnConflict := options.GetRetryOnConflict()
	if retryOnConflict > 0 {
		out["retryOnConflict"] = retryOnConflict
	}

	finalRequest, err := json.Marshal(out)

	if err != nil {
		if responseChannel != nil {
			responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
		}
		return
	}

	err = k.socket.Send(finalRequest, options, responseChannel, requestId)
	if err != nil {
		if responseChannel != nil {
			responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
		}
		return
	}
}

// Helper function copying headers to the query data
func (k Kuzzle) addHeaders(request *map[string]interface{}, query types.KuzzleRequest) {
	if k.jwt != "" && !(query.Controller == "auth" && query.Action == "checkToken") {
		(*request)["jwt"] = k.jwt
	}

	for k, v := range k.headers {
		if (*request)[k] == nil {
			(*request)[k] = v
		}
	}
}
