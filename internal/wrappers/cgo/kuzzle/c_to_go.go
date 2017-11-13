package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzlesdk.h"
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

func cToGoControllers(c *C.controllers) (*types.Controllers, error) {
	if c == nil {
		return nil, nil
	}

	if JsonCType(c) != C.json_type_object {
		return nil, types.NewError("Invalid controller structure", 400)
	}
	j := JsonCConvert(c)
	controllers := &types.Controllers{}

	for controller, val := range j.(map[string]interface{}) {
		rawController, ok := val.(map[string]interface{})
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		rawActions, ok := rawController["Actions"]
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		actionsMap, ok := rawActions.(map[string]interface{})
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		controllers.Controllers[controller] = &types.Controller{}
		for action, value := range actionsMap {
			boolValue, ok := value.(bool)
			if !ok {
				return nil, types.NewError("Invalid controllers structure", 400)
			}

			controllers.Controllers[controller].Actions[action] = boolValue
		}
	}

	return controllers, nil
}

func cToGoRole(crole *C.role) (*security.Role, error) {
	id := C.GoString(crole.id)
	var controllers *types.Controllers

	if crole.controllers != nil {
		c, err := cToGoControllers(crole.controllers)

		if err != nil {
			return nil, err
		}
		controllers = c
	}

	role := (*kuzzle.Kuzzle)(crole.kuzzle.instance).Security.NewRole(id, controllers)

	return role, nil
}

func cToGoSearchFilters(searchFilters *C.search_filters) *types.SearchFilters {
	return &types.SearchFilters{
		Query:        JsonCConvert(searchFilters.query),
		Sort:         JsonCConvert(searchFilters.sort).([]interface{}),
		Aggregations: JsonCConvert(searchFilters.aggregations),
		SearchAfter:  JsonCConvert(searchFilters.search_after).([]interface{}),
	}
}

// convert a C char** to a go array of string
func cToGoStrings(arr **C.char, len C.uint) []string {
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

// Helper to convert a C document** to a go array of document pointers
func cToGoDocuments(col *C.collection, docs **C.document, length C.uint) []*collection.Document {
	if length == 0 {
		return nil
	}
	tmpslice := (*[1 << 30]*C.document)(unsafe.Pointer(docs))[:length:length]
	godocuments := make([]*collection.Document, length)
	for i, doc := range tmpslice {
		godocuments[i] = cToGoDocument(col, doc)
	}
	return godocuments
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
	return collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
}

func cToGoMapping(cMapping *C.mapping) *collection.Mapping {
	mapping := collection.NewMapping(cToGoCollection(cMapping.collection))
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cMapping.mapping))), &mapping.Mapping)

	return mapping
}

func cToGoSpecification(cSpec *C.specification) *types.Specification {
	spec := types.Specification{}
	spec.Strict = bool(cSpec.strict)
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cSpec.fields))), &spec.Fields)
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cSpec.validators))), &spec.Validators)

	return &spec
}

func cToGoDocument(c *C.collection, cDoc *C.document) *collection.Document {
	gDoc := cToGoCollection(c).Document()
	gDoc.Id = C.GoString(cDoc.id)
	gDoc.Index = C.GoString(cDoc.index)
	gDoc.Meta = cToGoKuzzleMeta(cDoc.meta)
	gDoc.Shards = cToGoShards(cDoc.shards)
	gDoc.Content = []byte(C.GoString(C.json_object_to_json_string(cDoc.content)))
	gDoc.Version = int(cDoc.version)
	gDoc.Result = C.GoString(cDoc.result)
	gDoc.Collection = C.GoString(cDoc.collection)
	gDoc.Created = bool(cDoc.created)

	return gDoc
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
		Id:       C.GoString(p.id),
		Security: (*kuzzle.Kuzzle)(p.kuzzle.instance).Security,
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

func cToGoUserData(cdata *C.user_data) (*types.UserData, error) {
	if cdata == nil {
		return nil, nil
	}

	data := &types.UserData{}

	if cdata.content != nil {
		err := json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cdata.content))), data.Content)
		if err != nil {
			return nil, types.NewError(err.Error(), 400)
		}
	}

	if cdata.profile_ids_length > 0 {
		size := int(cdata.profile_ids_length)
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cdata.profile_ids))[:size:size]
		for i := 0; i < size; i++ {
			data.ProfileIds = append(data.ProfileIds, C.GoString(carray[i]))
		}
	}

	return data, nil
}

func cToGoUser(u *C.user) *security.User {
	if u == nil {
		return nil
	}

	user := (*kuzzle.Kuzzle)(u.kuzzle.instance).Security.NewUser(C.GoString(u.id), nil)

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
