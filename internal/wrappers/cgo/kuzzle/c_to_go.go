// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// convert a C char** to a go array of string
func cToGoStrings(arr **C.char, len C.size_t) []string {
	if len == 0 {
		return nil
	}

	tmpslice := (*[1 << 27]*C.char)(unsafe.Pointer(arr))[:len:len]
	goStrings := make([]string, 0, len)

	for _, s := range tmpslice {
		goStrings = append(goStrings, C.GoString(s))
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

	slice := (*[1<<28 - 1]*C.char)(unsafe.Pointer(r.collections))[:r.collections_length:r.collections_length]
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

	slice := (*[1<<28 - 1]*C.policy_restriction)(unsafe.Pointer(p.restricted_to))[:p.restricted_to_length]
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

	slice := (*[1<<28 - 1]*C.policy)(unsafe.Pointer(p.policies))[:p.policies_length]
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

	c, _ := json.Marshal(s.collection)

	var documents json.RawMessage
	d, _ := json.Marshal(s.documents)
	documents = d

	return &types.SearchResult{
		Collection: string(c),
		Documents:  documents,
		Total:      int(s.total),
		Fetched:    int(s.fetched),
		Options:    opts,
		Filters:    json.RawMessage(C.GoString(s.filters)),
	}
}

func cToGoKuzzleNotificationChannel(c *C.kuzzle_notification_listener) chan<- types.KuzzleNotification {
	return make(chan<- types.KuzzleNotification)
}
