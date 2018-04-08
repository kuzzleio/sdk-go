package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

func cToGoSearchFilters(searchFilters *C.search_filters) *types.SearchFilters {
	if searchFilters == nil {
		return nil
	}
	return &types.SearchFilters{
		Query:        json.RawMessage(C.GoString(searchFilters.query)),
		Sort:         json.RawMessage(C.GoString(searchFilters.sort)),
		Aggregations: json.RawMessage(C.GoString(searchFilters.aggregations)),
		SearchAfter:  json.RawMessage(C.GoString(searchFilters.search_after)),
	}
}

// convert a C char** to a go array of string
func cToGoStrings(arr **C.char, len C.size_t) []string {
	if len == 0 {
		return nil
	}

	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(arr))[:len:len]
	goStrings := make([]string, len)
	for i, s := range tmpslice {
		goStrings[i] = C.GoString(s)
	}

	return goStrings
}

func cToGoShards(cShards *C.shards) *types.Shards {
	return &types.Shards{
		Total:      int(cShards.total),
		Successful: int(cShards.successful),
		Failed:     int(cShards.failed),
	}
}

func cToGoKuzzleMeta(cMeta *C.meta) *types.Meta {
	return &types.Meta{
		Author:    C.GoString(cMeta.author),
		CreatedAt: int(cMeta.created_at),
		UpdatedAt: int(cMeta.updated_at),
		Updater:   C.GoString(cMeta.updater),
		Active:    bool(cMeta.active),
		DeletedAt: int(cMeta.deleted_at),
	}
}

func cToGoCollection(c *C.collection) *collection.Collection {
	return collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance))
}

func cToGoPolicyRestriction(r *C.policy_restriction) *types.PolicyRestriction {
	restriction := &types.PolicyRestriction{
		Index: C.GoString(r.index),
	}

	if r.collections == nil {
		return restriction
	}

	slice := (*[1<<30 - 1]*C.char)(unsafe.Pointer(r.collections))[:r.collections_length:r.collections_length]
	for _, col := range slice {
		restriction.Collections = append(restriction.Collections, C.GoString(col))
	}

	return restriction
}

func cToGoPolicy(p *C.policy) *types.Policy {
	policy := &types.Policy{
		RoleId: C.GoString(p.role_id),
	}

	if p.restricted_to == nil {
		return policy
	}

	slice := (*[1<<30 - 1]*C.policy_restriction)(unsafe.Pointer(p.restricted_to))[:p.restricted_to_length]
	for _, crestriction := range slice {
		policy.RestrictedTo = append(policy.RestrictedTo, cToGoPolicyRestriction(crestriction))
	}

	return policy
}

func cToGoProfile(p *C.profile) *security.Profile {
	profile := &security.Profile{
		Id: C.GoString(p.id),
	}

	if p.policies == nil {
		return profile
	}

	slice := (*[1<<30 - 1]*C.policy)(unsafe.Pointer(p.policies))[:p.policies_length]
	for _, cpolicy := range slice {
		profile.Policies = append(profile.Policies, cToGoPolicy(cpolicy))
	}

	return profile
}

func cToGoUser(u *C.user) *security.User {
	if u == nil {
		return nil
	}

	user := security.NewUser(C.GoString(u.id), nil)

	return user
}

func cToGoUserRigh(r *C.user_right) *types.UserRights {
	right := &types.UserRights{
		Controller: C.GoString(r.controller),
		Action:     C.GoString(r.action),
		Index:      C.GoString(r.index),
		Value:      C.GoString(r.value),
	}

	return right
}

func cToGoSearchResult(s *C.search_result) *types.SearchResult {
	opts := types.NewQueryOptions()

	opts.SetSize(int(s.options.size))
	opts.SetFrom(int(s.options.from))
	opts.SetScrollId(C.GoString(s.options.scroll_id))

	var collections json.RawMessage
	c, _ := json.Marshal(s.collection)
	collections = c

	var documents json.RawMessage
	d, _ := json.Marshal(s.documents)
	documents = d

	return &types.SearchResult{
		Collection: collections,
		Documents:  documents,
		Total:      int(s.total),
		Fetched:    int(s.fetched),
		Options:    opts,
		Filters:    cToGoSearchFilters(s.filters),
	}
}

func cToGoKuzzleNotificationChannel(c *C.kuzzle_notification_listener) chan<- types.KuzzleNotification {
	return make(chan<- types.KuzzleNotification)
}
