package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_query
func kuzzle_wrapper_query(k *C.kuzzle, request *C.kuzzle_request, options *C.query_options) *C.kuzzle_response {
	opts := SetQueryOptions(options)

	req := types.KuzzleRequest{
		RequestId:  C.GoString(request.request_id),
		Controller: C.GoString(request.controller),
		Action:     C.GoString(request.action),
		Index:      C.GoString(request.index),
		Collection: C.GoString(request.collection),
		Id:         C.GoString(request.id),
		From:       int(request.from),
		Size:       int(request.size),
		Scroll:     C.GoString(request.scroll),
		ScrollId:   C.GoString(request.scroll_id),
		Strategy:   C.GoString(request.strategy),
		ExpiresIn:  int(request.expires_in),
		Scope:      C.GoString(request.scope),
		State:      C.GoString(request.state),
		Users:      C.GoString(request.user),
		Start:      int(request.start),
		Stop:       int(request.stop),
		End:        int(request.end),
		Bit:        int(request.bit),
		Member:     C.GoString(request.member),
		Member1:    C.GoString(request.member1),
		Member2:    C.GoString(request.member2),
		Lon:        float64(request.lon),
		Lat:        float64(request.lat),
		Distance:   float64(request.distance),
		Unit:       C.GoString(request.unit),
		Cursor:     int(request.cursor),
		Offset:     int(request.offset),
		Field:      C.GoString(request.field),
		Subcommand: C.GoString(request.subcommand),
		Pattern:    C.GoString(request.pattern),
		Idx:        int(request.idx),
		Min:        C.GoString(request.min),
		Max:        C.GoString(request.max),
		Limit:      C.GoString(request.limit),
		Count:      int(request.count),
		Match:      C.GoString(request.match),
	}

	if request.body != nil {
		req.Body = JsonCConvert(request.body)
	}

	if request.volatiles != nil {
		req.Volatile = JsonCConvert(request.volatiles).(map[string]interface{})
	}

	start := int(request.start)
	req.Start = start
	req.Members = cToGoStrings(request.members, request.members_length)
	req.Keys = cToGoStrings(request.keys, request.keys_length)
	req.Fields = cToGoStrings(request.fields, request.fields_length)

	resC := make(chan *types.KuzzleResponse)
	(*kuzzle.Kuzzle)(k.instance).Query(&req, opts, resC)

	res := <-resC

	return goToCKuzzleResponse(res)
}
