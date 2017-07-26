package kuzzle

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/satori/go.uuid"
)

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k Kuzzle) Query(query types.KuzzleRequest, options *types.Options, responseChannel chan<- types.KuzzleResponse) {
	requestId := uuid.NewV4().String()

	query.RequestId = requestId

	type body struct{}

	if query.Body == nil {
		query.Body = make(map[string]interface{})
	}

	if k.volatile != nil {
		query.Volatile = k.volatile
		query.Volatile["sdkVersion"] = version
	} else {
		query.Volatile = types.VolatileData{"sdkVersion": version}
	}

	jsonRequest, _ := json.Marshal(query)
	out := map[string]interface{}{}
	json.Unmarshal(jsonRequest, &out)
	k.addHeaders(&out, query)

	if options == nil {
		options = types.DefaultOptions()
	}

	if options.Refresh != "" {
		out["refresh"] = options.Refresh
	}

	out["from"] = options.From
	out["size"] = options.Size

	if options.Scroll != "" {
		out["scroll"] = options.Scroll
	}

	if options.ScrollId != "" {
		out["scrollId"] = options.ScrollId
	}

	finalRequest, err := json.Marshal(out)

	if err != nil {
		responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
		return
	}

	err = k.socket.Send(finalRequest, options, responseChannel, requestId)
	if err != nil {
		responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
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
