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
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

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
func goToCDocument(col *C.collection, gDoc *collection.Document, dest *C.document) *C.document {
	var result *C.document
	if dest == nil {
		result = (*C.document)(C.calloc(1, C.sizeof_document))
	} else {
		result = dest
	}

	result.id = C.CString(gDoc.Id)
	result.index = C.CString(gDoc.Index)
	result.result = C.CString(gDoc.Result)
	result.collection = C.CString(gDoc.Collection)
	result.meta = goToCMeta(gDoc.Meta)
	result.shards = goToCShards(gDoc.Shards)
	result._collection = col

	if string(gDoc.Content) != "" {
		buffer := C.CString(string(gDoc.Content))
		result.content = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	} else {
		result.content = C.json_object_new_object()
	}

	result.version = C.int(gDoc.Version)
	result.created = C.bool(gDoc.Created)

	return result
}

// Allocates memory
func goToCNotificationContent(gNotifContent *types.NotificationResult) *C.notification_content {
	result := (*C.notification_content)(C.calloc(1, C.sizeof_notification_content))
	result.id = C.CString(gNotifContent.Id)
	result.meta = goToCMeta(gNotifContent.Meta)
	result.count = C.int(gNotifContent.Count)

	r, _ := json.Marshal(gNotifContent.Content)
	buffer := C.CString(string(r))
	result.content = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return result
}

// Allocates memory
func goToCNotificationResult(gNotif *types.KuzzleNotification) *C.notification_result {
	result := (*C.notification_result)(C.calloc(1, C.sizeof_notification_result))

	if gNotif.Error != nil {
		Set_notification_result_error(result, gNotif.Error)
		return result
	}

	result.request_id = C.CString(gNotif.RequestId)
	result.result = goToCNotificationContent(gNotif.Result)

	r, _ := json.Marshal(gNotif.Volatile)
	buffer := C.CString(string(r))
	result.volatiles = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

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
	result.result = C.json_tokener_parse(bufResult)
	C.free(unsafe.Pointer(bufResult))

	r, _ := json.Marshal(gRes.Volatile)
	bufVolatile := C.CString(string(r))
	result.volatiles = C.json_tokener_parse(bufVolatile)
	C.free(unsafe.Pointer(bufVolatile))

	result.index = C.CString(gRes.Index)
	result.collection = C.CString(gRes.Collection)
	result.controller = C.CString(gRes.Controller)
	result.action = C.CString(gRes.Action)
	result.room_id = C.CString(gRes.RoomId)
	result.channel = C.CString(gRes.Channel)
	result.status = C.int(gRes.Status)

	if gRes.Error != nil {
		// The error might be a partial error
		Set_kuzzle_response_error(result, gRes.Error)
	}

	return result
}

// Allocates memory
func goToCDocumentResult(col *C.collection, goRes *collection.Document, err error) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(col, goRes, nil)

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
		collections := (*[1<<30 - 1]*C.char)(unsafe.Pointer(crestriction.collections))[:len(restriction.Collections)]

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

func goToCStringArrayResult(goRes []string, err error) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	if goRes != nil {
		result.result = (**C.char)(C.calloc(C.size_t(len(goRes)), C.sizeof_char_ptr))
		result.result_length = C.size_t(len(goRes))

		cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

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
		restrictions := (*[1<<30 - 1]C.policy_restriction)(unsafe.Pointer(cpolicy.restricted_to))[:len(policy.RestrictedTo)]

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
		policies := (*[1<<30 - 1]C.policy)(unsafe.Pointer(cprofile.policies))[:len(profile.Policies)]
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
		profiles := (*[1<<30 - 1]C.profile)(unsafe.Pointer(result.result.hits))[:len(res.Hits)]

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
		cArray := (*[1<<30 - 1]C.role)(unsafe.Pointer(result.result.hits))[:len(res.Hits):len(res.Hits)]

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
func goToCSearchResult(col *C.collection, goRes *collection.SearchResult, err error) *C.search_result {
	result := (*C.search_result)(C.calloc(1, C.sizeof_search_result))

	if err != nil {
		Set_search_result_error(result, err)
		return result
	}

	result.result = (*C.document_search)(C.calloc(1, C.sizeof_document_search))
	result.result.hits_length = C.size_t(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	if goRes.ScrollId != "" {
		result.result.scroll_id = C.CString(goRes.ScrollId)
	}

	if len(goRes.Hits) > 0 {
		result.result.hits = (*C.document)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_document))
		cArray := (*[1<<30 - 1]C.document)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, doc := range goRes.Hits {
			goToCDocument(col, doc, &cArray[i])
		}
	}

	return result
}

// Allocates memory
func goToCMapping(c *C.collection, goMapping *collection.Mapping) *C.mapping {
	result := (*C.mapping)(C.calloc(1, C.sizeof_mapping))

	result.collection = c
	r, _ := json.Marshal(goMapping.Mapping)
	buffer := C.CString(string(r))
	result.mapping = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return result
}

// Allocates memory
func goToCMappingResult(c *C.collection, goRes *collection.Mapping, err error) *C.mapping_result {
	result := (*C.mapping_result)(C.calloc(1, C.sizeof_mapping_result))

	if err != nil {
		Set_mapping_result_error(result, err)
		return result
	}

	result.result = goToCMapping(c, goRes)

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
		j, _ := json.Marshal(role.Controllers)
		buffer := C.CString(string(j))
		crole.controllers = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	}

	return crole
}

// Allocates memory
func goToCSpecification(goSpec *types.Specification) *C.specification {
	result := (*C.specification)(C.calloc(1, C.sizeof_specification))

	result.strict = C.bool(goSpec.Strict)

	f, _ := json.Marshal(goSpec.Fields)
	v, _ := json.Marshal(goSpec.Validators)
	bufferFields := C.CString(string(f))
	bufferValidators := C.CString(string(v))

	result.fields = C.json_tokener_parse(bufferFields)
	result.validators = C.json_tokener_parse(bufferValidators)

	C.free(unsafe.Pointer(bufferFields))
	C.free(unsafe.Pointer(bufferValidators))

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
		cArray := (*[1<<30 - 1]C.specification_entry)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, spec := range goRes.Hits {
			goToCSpecificationEntry(&spec.Source, &cArray[i])
		}
	}

	return result
}

func goToCJson(data interface{}) (*C.json_object, error) {
	r, err := json.Marshal(data)
	if err != nil {
		return nil, types.NewError(err.Error(), 400)
	}

	buffer := C.CString(string(r))
	defer C.free(unsafe.Pointer(buffer))

	tok := C.json_tokener_new()
	j := C.json_tokener_parse_ex(tok, buffer, C.int(C.strlen(buffer)))
	jerr := C.json_tokener_get_error(tok)
	if jerr != C.json_tokener_success {
		return nil, types.NewError(C.GoString(C.json_tokener_error_desc(jerr)), 400)
	}

	return j, nil
}

func goToCJsonResult(goRes interface{}, err error) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))

	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	result.result, err = goToCJson(goRes)
	if err != nil {
		Set_json_result_error(result, err)
		return result
	}

	return result
}

func goToCJsonArrayResult(goRes []interface{}, err error) *C.json_array_result {
	result := (*C.json_array_result)(C.calloc(1, C.sizeof_json_array_result))

	if err != nil {
		Set_json_array_result_error(result, err)
		return result
	}

	result.result_length = C.size_t(len(goRes))
	if goRes != nil {
		result.result = (**C.json_object)(C.calloc(result.result_length, C.sizeof_json_object_ptr))
		cArray := (*[1<<30 - 1]*C.json_object)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

		for i, res := range goRes {
			cArray[i], err = goToCJson(res)
			if err != nil {
				Set_json_array_result_error(result, err)
				return result
			}
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
		jsonO, err := goToCJson(data.Content)
		if err != nil {
			return nil, err
		}
		cdata.content = jsonO
	}

	if data.ProfileIds != nil {
		cdata.profile_ids_length = C.size_t(len(data.ProfileIds))
		cdata.profile_ids = (**C.char)(C.calloc(C.size_t(len(data.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cdata.profile_ids))[:len(data.ProfileIds):len(data.ProfileIds)]

		for i, profileId := range data.ProfileIds {
			carray[i] = C.CString(profileId)
		}
	}

	return cdata, nil
}

func goToCUser(k *C.kuzzle, user *security.User, dest *C.user) (*C.user, error) {
	if user == nil {
		return nil, nil
	}

	var cuser *C.user
	if dest == nil {
		cuser = (*C.user)(C.calloc(1, C.sizeof_user))
	} else {
		cuser = dest
	}

	cuser.id = C.CString(user.Id)
	cuser.kuzzle = k

	if user.Content != nil {
		jsonO, err := goToCJson(user.Content)
		if err != nil {
			return nil, err
		}
		cuser.content = jsonO
	}

	if user.ProfileIds != nil {
		cuser.profile_ids_length = C.size_t(len(user.ProfileIds))
		cuser.profile_ids = (**C.char)(C.calloc(C.size_t(len(user.ProfileIds)), C.sizeof_char_ptr))
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cuser.profile_ids))[:len(user.ProfileIds):len(user.ProfileIds)]

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

	result.user = cuser

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
		carray := (*[1<<30 - 1]C.profile)(unsafe.Pointer(result.profiles))[:len(profiles):len(profiles)]

		for i, profile := range profiles {
			goToCProfile(k, profile, &carray[i])
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
		result.user_rights = (*C.user_right)(C.calloc(C.size_t(len(rights)), C.sizeof_user_right))
		carray := (*[1<<30 - 1]C.user_right)(unsafe.Pointer(result.user_rights))[:len(rights):len(rights)]

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
		result.result.hits = (*C.user)(C.calloc(C.size_t(len(res.Hits)), C.sizeof_user))
		users := (*[1<<30 - 1]C.user)(unsafe.Pointer(result.result.hits))[:len(res.Hits)]

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
		carray := (*[1<<30 - 1]C.collection_entry)(unsafe.Pointer(result.result))[:len(collections):len(collections)]

		for i, collection := range collections {
			fillCollectionList(collection, &carray[i])
		}
	}

	return result
}

// Allocates memory
func fillStatistics(res *types.Statistics, statistics *C.statistics) {
	ongoing, _ := json.Marshal(res.OngoingRequests)
	completedRequests, _ := json.Marshal(res.CompletedRequests)
	connections, _ := json.Marshal(res.Connections)
	failedRequests, _ := json.Marshal(res.FailedRequests)

	cOnGoing := C.CString(string(ongoing))
	cCompleteRequest := C.CString(string(completedRequests))
	cConnections := C.CString(string(connections))
	cFailedRequests := C.CString(string(failedRequests))

	statistics.ongoing_requests = C.json_tokener_parse(cOnGoing)
	statistics.completed_requests = C.json_tokener_parse(cCompleteRequest)
	statistics.connections = C.json_tokener_parse(cConnections)
	statistics.failed_requests = C.json_tokener_parse(cFailedRequests)
	statistics.timestamp = C.ulonglong(res.Timestamp)

	C.free(unsafe.Pointer(cOnGoing))
	C.free(unsafe.Pointer(cCompleteRequest))
	C.free(unsafe.Pointer(cConnections))
	C.free(unsafe.Pointer(cFailedRequests))
}

// Allocates memory
func goToCVoidResult(err error) *C.void_result {
	if err == nil {
		return nil
	}

	result := (*C.void_result)(C.calloc(1, C.sizeof_void_result))
	Set_void_result_error(result, err)

	return result
}
