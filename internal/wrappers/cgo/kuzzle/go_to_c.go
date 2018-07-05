// Copyright 2015-2017 Kuzzle
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
	#include <string.h>
	#include <stdlib.h>
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

func duplicateCollection(ptr *C.collection) *C.collection {
	dest := (*C.collection)(C.calloc(1, C.sizeof_collection))

	dest.kuzzle = ptr.kuzzle

	return dest
}

// Allocates memory
func goToCMeta(gMeta *types.Meta) *C.meta {
	if gMeta == nil {
		return nil
	}

	result := (*C.meta)(C.calloc(1, C.sizeof_meta))
	result.author = C.CString(gMeta.Author)
	result.created_at = C.ulonglong(gMeta.CreatedAt)
	result.updated_at = C.ulonglong(gMeta.UpdatedAt)
	result.deleted_at = C.ulonglong(gMeta.DeletedAt)
	result.updater = C.CString(gMeta.Updater)
	result.active = C.bool(gMeta.Active)

	return result
}

// Allocates memory
func goToCShards(gShards *types.Shards) *C.shards {
	if gShards == nil {
		return nil
	}

	result := (*C.shards)(C.calloc(1, C.sizeof_shards))
	result.failed = C.int(gShards.Failed)
	result.successful = C.int(gShards.Successful)
	result.total = C.int(gShards.Total)

	return result
}

// Allocates memory
func goToCNotificationContent(gNotifContent *types.NotificationResult) *C.notification_content {
	result := (*C.notification_content)(C.calloc(1, C.sizeof_notification_content))
	result.id = C.CString(gNotifContent.Id)
	result.meta = goToCMeta(gNotifContent.Meta)
	result.count = C.int(gNotifContent.Count)
	marshalled, _ := json.Marshal(gNotifContent)
	result.content = C.CString(string(marshalled))

	return result
}

// Allocates memory
func goToCNotificationResult(gNotif *types.KuzzleNotification) *C.notification_result {
	result := (*C.notification_result)(C.calloc(1, C.sizeof_notification_result))

	if gNotif.Error.Error() != "" {
		Set_notification_result_error(result, gNotif.Error)
		return result
	}

	result.volatiles = C.CString(string(gNotif.Volatile))
	result.request_id = C.CString(gNotif.RequestId)
	result.result = goToCNotificationContent(gNotif.Result)
	result.index = C.CString(gNotif.Index)
	result.collection = C.CString(gNotif.Collection)
	result.controller = C.CString(gNotif.Controller)
	result.action = C.CString(gNotif.Action)
	result.protocol = C.CString(gNotif.Protocol)
	result.scope = C.CString(gNotif.Scope)
	result.state = C.CString(gNotif.State)
	result.user = C.CString(gNotif.User)
	result.n_type = C.CString(gNotif.Type)
	result.room_id = C.CString(gNotif.RoomId)
	result.timestamp = C.ulonglong(gNotif.Timestamp)
	result.status = C.int(gNotif.Status)

	return result
}

func goToCKuzzleResponse(gRes *types.KuzzleResponse) *C.kuzzle_response {
	result := (*C.kuzzle_response)(C.calloc(1, C.sizeof_kuzzle_response))

	result.request_id = C.CString(gRes.RequestId)

	bufResult := C.CString(string(gRes.Result))
	result.result = bufResult
	C.free(unsafe.Pointer(bufResult))

	result.volatiles = C.CString(string(gRes.Volatile))
	result.index = C.CString(gRes.Index)
	result.collection = C.CString(gRes.Collection)
	result.controller = C.CString(gRes.Controller)
	result.action = C.CString(gRes.Action)
	result.room_id = C.CString(gRes.RoomId)
	result.channel = C.CString(gRes.Channel)
	result.status = C.int(gRes.Status)

	if gRes.Error.Error() != "" {
		// The error might be a partial error
		Set_kuzzle_response_error(result, gRes.Error)
	}

	return result
}

func goToCPolicyRestriction(restriction *types.PolicyRestriction, dest *C.policy_restriction) *C.policy_restriction {
	var crestriction *C.policy_restriction
	if dest == nil {
		crestriction = (*C.policy_restriction)(C.calloc(1, C.sizeof_policy_restriction))
	} else {
		crestriction = dest
	}
	crestriction.index = C.CString(restriction.Index)
	crestriction.collections_length = C.size_t(len(restriction.Collections))

	if restriction.Collections != nil {
		crestriction.collections = (**C.char)(C.calloc(C.size_t(len(restriction.Collections)), C.sizeof_char_ptr))
		collections := (*[1<<28 - 1]*C.char)(unsafe.Pointer(crestriction.collections))[:len(restriction.Collections)]

		for i, col := range restriction.Collections {
			collections[i] = C.CString(col)
		}
	}

	return crestriction
}

// Allocates memory
func goToCStringResult(goRes *string, err error) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))

	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	if goRes != nil {
		result.result = C.CString(*goRes)
	}

	return result
}

func goToCSubscribeResult(goRes *types.SubscribeResult, err error) *C.subscribe_result {
	result := (*C.subscribe_result)(C.calloc(1, C.sizeof_subscribe_result))

	if err != nil {
		Set_subscribe_error(result, err)
		return result
	}

	result.room = C.CString(goRes.Room)
	result.channel = C.CString(goRes.Channel)

	return result
}

func goToCStringArrayResult(goRes []string, err error) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	if goRes != nil {
		result.result = (**C.char)(C.calloc(C.size_t(len(goRes)), C.sizeof_char_ptr))
		result.result_length = C.size_t(len(goRes))

		cArray := (*[1<<28 - 1]*C.char)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

		for i, substring := range goRes {
			cArray[i] = C.CString(substring)
		}
	}

	return result
}

// Allocates memory
func goToCIntResult(goRes int, err error) *C.int_result {
	result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

	if err != nil {
		Set_int_result_error(result, err)
		return result
	}

	result.result = C.longlong(goRes)

	return result
}

//Allocates memory
func goToCIntArrayResult(goRes []int, err error) *C.int_array_result {
	result := (*C.int_array_result)(C.calloc(1, C.sizeof_int_array_result))

	if err != nil {
		Set_int_array_result_error(result, err)
		return result
	}

	if goRes != nil {
		result.result = (*C.longlong)(C.calloc(C.size_t(len(goRes)), C.sizeof_longlong))
		result.result_length = C.size_t(len(goRes))

		cArray := (*[1<<20 - 1]C.longlong)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

		for i, num := range goRes {
			cArray[i] = C.longlong(num)
		}
	}

	return result
}

// Allocates memory
func goToCDoubleResult(goRes float64, err error) *C.double_result {
	result := (*C.double_result)(C.calloc(1, C.sizeof_double_result))

	if err != nil {
		Set_double_result_error(result, err)
		return result
	}

	result.result = C.double(goRes)

	return result
}

func goToCPolicy(policy *types.Policy, dest *C.policy) *C.policy {
	var cpolicy *C.policy
	if dest == nil {
		cpolicy = (*C.policy)(C.calloc(1, C.sizeof_policy))
	} else {
		cpolicy = dest
	}

	cpolicy.role_id = C.CString(policy.RoleId)
	cpolicy.restricted_to_length = C.size_t(len(policy.RestrictedTo))

	if policy.RestrictedTo != nil {
		cpolicy.restricted_to = (*C.policy_restriction)(C.calloc(C.size_t(len(policy.RestrictedTo)), C.sizeof_policy_restriction))
		restrictions := (*[1<<27 - 1]C.policy_restriction)(unsafe.Pointer(cpolicy.restricted_to))[:len(policy.RestrictedTo)]

		for i, restriction := range policy.RestrictedTo {
			goToCPolicyRestriction(restriction, &restrictions[i])
		}
	}

	return cpolicy
}

func goToCProfile(k *C.kuzzle, profile *security.Profile, dest *C.profile) *C.profile {
	var cprofile *C.profile
	if dest == nil {
		cprofile = (*C.profile)(C.calloc(1, C.sizeof_profile))
	} else {
		cprofile = dest
	}

	cprofile.id = C.CString(profile.Id)
	cprofile.policies_length = C.size_t(len(profile.Policies))
	cprofile.kuzzle = k

	if profile.Policies != nil {
		cprofile.policies = (*C.policy)(C.calloc(C.size_t(len(profile.Policies)), C.sizeof_policy))
		policies := (*[1<<27 - 1]C.policy)(unsafe.Pointer(cprofile.policies))[:len(profile.Policies)]
		for i, policy := range profile.Policies {
			goToCPolicy(policy, &policies[i])
		}
	}

	return cprofile
}

func goToCProfileSearchResult(k *C.kuzzle, res *security.ProfileSearchResult, err error) *C.search_profiles_result {
	result := (*C.search_profiles_result)(C.calloc(1, C.sizeof_search_profiles_result))

	if err != nil {
		Set_search_profiles_result_error(result, err)
		return result
	}

	result.result = (*C.profile_search)(C.calloc(1, C.sizeof_profile_search))
	result.result.hits_length = C.size_t(len(res.Hits))
	result.result.total = C.uint(res.Total)
	if res.ScrollId != "" {
		result.result.scroll_id = C.CString(res.ScrollId)
	}

	if len(res.Hits) > 0 {
		result.result.hits = (*C.profile)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_profile))
		profiles := (*[1<<27 - 1]C.profile)(unsafe.Pointer(result.result.hits))[:len(res.Hits)]

		for i, profile := range res.Hits {
			goToCProfile(k, profile, &profiles[i])
		}
	}

	return result
}

func goToCRoleSearchResult(k *C.kuzzle, res *security.RoleSearchResult, err error) *C.search_roles_result {
	result := (*C.search_roles_result)(C.calloc(1, C.sizeof_search_roles_result))

	if err != nil {
		Set_search_roles_result_error(result, err)
		return result
	}

	result.result = (*C.role_search)(C.calloc(1, C.sizeof_role_search))
	result.result.hits_length = C.size_t(len(res.Hits))
	result.result.total = C.uint(res.Total)

	if len(res.Hits) > 0 {
		result.result.hits = (*C.role)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_role))
		cArray := (*[1<<27 - 1]C.role)(unsafe.Pointer(result.result.hits))[:len(res.Hits):len(res.Hits)]

		for i, role := range res.Hits {
			goToCRole(k, role, &cArray[i])
		}
	}

	return result
}

// Allocates memory
func goToCBoolResult(goRes bool, err error) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(goRes)

	return result
}

// Allocates memory
func goToCSearchResult(goRes *types.SearchResult, err error) *C.search_result {
	result := (*C.search_result)(C.calloc(1, C.sizeof_search_result))

	if err != nil {
		Set_search_result_error(result, err)
		return result
	}

	result.collection = C.CString(string(goRes.Collection))
	result.documents = C.CString(string(goRes.Documents))

	result.fetched = C.uint(goRes.Fetched)
	result.total = C.uint(goRes.Total)

	if goRes.Filters != nil {
		result.filters = C.CString(string(goRes.Filters))
	}

	result.options = (*C.query_options)(C.calloc(1, C.sizeof_query_options))
	if goRes.Options != nil {
		result.options.from = C.long(goRes.Options.From())
		result.options.size = C.long(goRes.Options.Size())

		if goRes.Options.ScrollId() != "" {
			result.options.scroll_id = C.CString(goRes.Options.ScrollId())
		}
	}

	if len(goRes.Aggregations) > 0 {
		result.aggregations = C.CString(string(goRes.Aggregations))
	}

	return result
}

func goToCRole(k *C.kuzzle, role *security.Role, dest *C.role) *C.role {
	var crole *C.role
	if dest == nil {
		crole = (*C.role)(C.calloc(1, C.sizeof_role))
	} else {
		crole = dest
	}

	crole.id = C.CString(role.Id)
	crole.kuzzle = k

	if role.Controllers != nil {
		ctrls, _ := json.Marshal(role.Controllers)
		crole.controllers = C.CString(string(ctrls))
	}

	return crole
}

// Allocates memory
func goToCSpecification(goSpec *types.Specification) *C.specification {
	result := (*C.specification)(C.calloc(1, C.sizeof_specification))

	result.strict = C.bool(goSpec.Strict)
	flds, _ := json.Marshal(goSpec.Fields)
	result.fields = C.CString(string(flds))
	result.validators = C.CString(string(goSpec.Validators))
	return result
}

// Allocates memory
func goToCSpecificationEntry(goEntry *types.SpecificationEntry, dest *C.specification_entry) *C.specification_entry {
	var result *C.specification_entry
	if dest == nil {
		result = (*C.specification_entry)(C.calloc(1, C.sizeof_specification_entry))
	} else {
		result = dest
	}

	result.index = C.CString(goEntry.Index)
	result.collection = C.CString(goEntry.Collection)
	result.validation = goToCSpecification(goEntry.Validation)

	return result
}

// Allocates memory
func goToCSpecificationResult(goRes *types.Specification, err error) *C.specification_result {
	result := (*C.specification_result)(C.calloc(1, C.sizeof_specification_result))

	if err != nil {
		Set_specification_result_err(result, err)
		return result
	}

	result.result = goToCSpecification(goRes)

	return result
}

// Allocates memory
func goToCSpecificationSearchResult(goRes *types.SpecificationSearchResult, err error) *C.specification_search_result {
	result := (*C.specification_search_result)(C.calloc(1, C.sizeof_specification_search_result))

	if err != nil {
		Set_specification_search_result_error(result, err)
		return result
	}

	result.result = (*C.specification_search)(C.calloc(1, C.sizeof_specification_search))
	result.result.hits_length = C.size_t(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)

	if goRes.ScrollId != "" {
		result.result.scroll_id = C.CString(goRes.ScrollId)
	}

	if len(goRes.Hits) > 0 {
		result.result.hits = (*C.specification_entry)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_specification_entry))
		cArray := (*[1<<27 - 1]C.specification_entry)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, spec := range goRes.Hits {
			goToCSpecificationEntry(&spec.Source, &cArray[i])
		}
	}

	return result
}

func goToCProfileResult(k *C.kuzzle, res *security.Profile, err error) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, res, nil)
	return result
}

func goToCUserData(data *types.UserData) (*C.user_data, error) {
	if data == nil {
		return nil, nil
	}

	cdata := (*C.user_data)(C.calloc(1, C.sizeof_user_data))

	if data.Content != nil {
		cnt, err := json.Marshal(data.Content)
		if err != nil {
			return nil, err
		}
		jsonO := C.CString(string(cnt))
		cdata.content = jsonO
	}

	if data.ProfileIds != nil {
		cdata.profile_ids_length = C.size_t(len(data.ProfileIds))
		cdata.profile_ids = (**C.char)(C.calloc(C.size_t(len(data.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<27 - 1]*C.char)(unsafe.Pointer(cdata.profile_ids))[:len(data.ProfileIds):len(data.ProfileIds)]

		for i, profileId := range data.ProfileIds {
			carray[i] = C.CString(profileId)
		}
	}

	return cdata, nil
}

func goToCUser(k *C.kuzzle, user *security.User, dest *C.kuzzle_user) (*C.kuzzle_user, error) {
	if user == nil {
		return nil, nil
	}

	var cuser *C.kuzzle_user
	if dest == nil {
		cuser = (*C.kuzzle_user)(C.calloc(1, C.sizeof_kuzzle_user))
	} else {
		cuser = dest
	}

	cuser.id = C.CString(user.Id)
	cuser.kuzzle = k

	if user.Content != nil {
		cnt, err := json.Marshal(user.Content)
		if err != nil {
			return nil, err
		}

		cuser.content = C.CString(string(cnt))
	}

	if user.ProfileIds != nil {
		cuser.profile_ids_length = C.size_t(len(user.ProfileIds))
		cuser.profile_ids = (**C.char)(C.calloc(C.size_t(len(user.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<28 - 1]*C.char)(unsafe.Pointer(cuser.profile_ids))[:len(user.ProfileIds):len(user.ProfileIds)]

		for i, profileId := range user.ProfileIds {
			carray[i] = C.CString(profileId)
		}
	}

	return cuser, nil
}

func goToCUserResult(k *C.kuzzle, user *security.User, err error) *C.user_result {
	result := (*C.user_result)(C.calloc(1, C.sizeof_user_result))
	if err != nil {
		Set_user_result_error(result, err)
		return result
	}

	cuser, err := goToCUser(k, user, nil)
	if err != nil {
		Set_user_result_error(result, err)
		return result
	}

	result.result = cuser

	return result
}

func goToCProfilesResult(k *C.kuzzle, profiles []*security.Profile, err error) *C.profiles_result {
	result := (*C.profiles_result)(C.calloc(1, C.sizeof_profiles_result))
	if err != nil {
		Set_profiles_result_error(result, err)
		return result
	}

	result.profiles_length = C.size_t(len(profiles))

	if profiles != nil {
		result.profiles = (*C.profile)(C.calloc(C.size_t(len(profiles)), C.sizeof_profile))
		carray := (*[1<<27 - 1]C.profile)(unsafe.Pointer(result.profiles))[:len(profiles):len(profiles)]

		for i, profile := range profiles {
			goToCProfile(k, profile, &carray[i])
		}
	}

	return result
}

func goToCRolesResult(k *C.kuzzle, roles []*security.Role, err error) *C.roles_result {
	result := (*C.roles_result)(C.calloc(1, C.sizeof_roles_result))
	if err != nil {
		Set_roles_result_error(result, err)
		return result
	}

	result.roles_length = C.size_t(len(roles))

	if roles != nil {
		result.roles = (*C.role)(C.calloc(C.size_t(len(roles)), C.sizeof_role))
		carray := (*[1<<27 - 1]C.role)(unsafe.Pointer(result.roles))[:len(roles):len(roles)]

		for i, role := range roles {
			goToCRole(k, role, &carray[i])
		}
	}

	return result
}

func goToCUserRight(right *types.UserRights, dest *C.user_right) *C.user_right {
	if right == nil {
		return nil
	}

	var cright *C.user_right
	if dest == nil {
		cright = (*C.user_right)(C.calloc(1, C.sizeof_user_right))
	} else {
		cright = dest
	}

	cright.controller = C.CString(right.Controller)
	cright.action = C.CString(right.Action)
	cright.index = C.CString(right.Index)
	cright.collection = C.CString(right.Collection)
	cright.value = C.CString(right.Value)

	return cright
}

func goToCUserRightsResult(rights []*types.UserRights, err error) *C.user_rights_result {
	result := (*C.user_rights_result)(C.calloc(1, C.sizeof_user_rights_result))
	if err != nil {
		Set_user_rights_error(result, err)
		return result
	}

	result.user_rights_length = C.size_t(len(rights))
	if rights != nil {
		result.result = (*C.user_right)(C.calloc(C.size_t(len(rights)), C.sizeof_user_right))
		carray := (*[1<<26 - 1]C.user_right)(unsafe.Pointer(result.result))[:len(rights):len(rights)]

		for i, right := range rights {
			goToCUserRight(right, &carray[i])
		}
	}

	return result
}

func goToCUserSearchResult(k *C.kuzzle, res *security.UserSearchResult, err error) *C.search_users_result {
	result := (*C.search_users_result)(C.calloc(1, C.sizeof_search_users_result))

	if err != nil {
		Set_search_users_result_error(result, err)
		return result
	}

	result.result = (*C.user_search)(C.calloc(1, C.sizeof_user_search))
	result.result.hits_length = C.size_t(len(res.Hits))
	result.result.total = C.uint(res.Total)
	if res.ScrollId != "" {
		result.result.scroll_id = C.CString(res.ScrollId)
	}

	if len(res.Hits) > 0 {
		result.result.hits = (*C.kuzzle_user)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_kuzzle_user))
		users := (*[1<<26 - 1]C.kuzzle_user)(unsafe.Pointer(result.result.hits))[:len(res.Hits)]

		for i, user := range res.Hits {
			goToCUser(k, user, &users[i])
		}
	}

	return result
}

func fillCollectionList(collection *types.CollectionsList, entry *C.collection_entry) {
	if collection == nil {
		return
	}

	entry.persisted = collection.Type == "persisted"
	entry.name = C.CString(collection.Name)
}

func goToCCollectionListResult(collections []*types.CollectionsList, err error) *C.collection_entry_result {
	result := (*C.collection_entry_result)(C.calloc(1, C.sizeof_collection_entry_result))
	if err != nil {
		Set_collection_entry_error(result, err)
		return result
	}

	if collections != nil {
		result.result = (*C.collection_entry)(C.calloc(C.size_t(len(collections)), C.sizeof_collection_entry))
		result.result_length = C.size_t(len(collections))
		carray := (*[1<<27 - 1]C.collection_entry)(unsafe.Pointer(result.result))[:len(collections):len(collections)]

		for i, collection := range collections {
			fillCollectionList(collection, &carray[i])
		}
	}

	return result
}

// Allocates memory
func fillStatistics(src *types.Statistics, dest *C.statistics) {
	dest.ongoing_requests = C.CString(string(src.OngoingRequests))
	dest.completed_requests = C.CString(string(src.CompletedRequests))
	dest.connections = C.CString(string(src.Connections))
	dest.failed_requests = C.CString(string(src.FailedRequests))
	dest.timestamp = C.ulonglong(src.Timestamp)
}

// Allocates memory
func goToCErrorResult(err error) *C.error_result {
	if err == nil {
		return nil
	}

	result := (*C.error_result)(C.calloc(1, C.sizeof_error_result))
	Set_error_result_error(result, err)

	return result
}

// Allocates memory
func goToCDateResult(goRes int, err error) *C.date_result {
	result := (*C.date_result)(C.calloc(1, C.sizeof_date_result))

	if err != nil {
		Set_date_result_error(result, err)
		return result
	}

	result.result = C.longlong(goRes)

	return result
}
