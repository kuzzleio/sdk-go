package main

/*
	#cgo CFLAGS: -I../../headers
	#include <errno.h>
	#include <stdlib.h>
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"

import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
)

// --- profile

//export kuzzle_security_new_profile
func kuzzle_security_new_profile(k *C.kuzzle, id *C.char, policies *C.policy, policies_length C.size_t) *C.profile {
	cprofile := (*C.profile)(C.calloc(1, C.sizeof_profile))
	cprofile.id = id
	cprofile.policies = policies
	cprofile.policies_length = policies_length
	cprofile.kuzzle = k

	return cprofile
}

//export kuzzle_security_destroy_profile
func kuzzle_security_destroy_profile(p *C.profile) {
	if p == nil {
		return
	}

	if p.policies != nil {
		size := int(p.policies_length)
		policies := (*[1<<30 - 1]*C.policy)(unsafe.Pointer(p.policies))[:size:size]
		for _, policy := range policies {
			C.free(unsafe.Pointer(policy.role_id))
			size = int(policy.restricted_to_length)
			restrictions := (*[1<<30 - 1]*C.policy_restriction)(unsafe.Pointer(policy.restricted_to))[:size:size]
			for _, restriction := range restrictions {
				C.free(unsafe.Pointer(restriction.index))
				size = int(restriction.collections_length)
				collections := (*[1<<30 - 1]*C.char)(unsafe.Pointer(restriction.collections))[:size:size]
				for _, collection := range collections {
					C.free(unsafe.Pointer(collection))
				}
				C.free(unsafe.Pointer(restriction.collections))
			}
			C.free(unsafe.Pointer(policy.restricted_to))
		}
		C.free(unsafe.Pointer(p.policies))
	}

	C.free(unsafe.Pointer(p))
}

// --- user

//export kuzzle_security_new_user
func kuzzle_security_new_user(k *C.kuzzle, id *C.char, d *C.user_data) *C.user {
	cuser := (*C.user)(C.calloc(1, C.sizeof_user))

	cuser.id = id
	cuser.kuzzle = k

	if d != nil {
		cuser.content = d.content
		cuser.profile_ids = d.profile_ids
		cuser.profile_ids_length = d.profile_ids_length
	}

	return cuser
}

//export kuzzle_security_destroy_user
func kuzzle_security_destroy_user(u *C.user) {
	if u == nil {
		return
	}

	if u.id != nil {
		C.free(unsafe.Pointer(u.id))
	}
	if u.content != nil {
		C.json_object_put(u.content)
	}
	if u.profile_ids != nil {
		size := int(u.profile_ids_length)
		carray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(u.profile_ids))[:size:size]

		for i := 0; i < size; i++ {
			C.free(unsafe.Pointer(carray[i]))
		}
		C.free(unsafe.Pointer(u.profile_ids))
	}

	C.free(unsafe.Pointer(u))
}

// --- role

//export kuzzle_security_new_role
func kuzzle_security_new_role(k *C.kuzzle, id *C.char, c *C.controllers) *C.role {
	crole := (*C.role)(C.calloc(1, C.sizeof_role))
	crole.id = id
	crole.controllers = c
	crole.kuzzle = k

	_, err := cToGoRole(crole)
	if err != nil {
		C.set_errno(C.ENOKEY)
		return nil
	}

	return crole
}

//export kuzzle_security_destroy_role
func kuzzle_security_destroy_role(r *C.role) {
	if r == nil {
		return
	}

	C.json_object_put(r.controllers)
	C.free(unsafe.Pointer(r))
}

//export kuzzle_security_get_profile
func kuzzle_security_get_profile(k *C.kuzzle, id *C.char, o *C.query_options) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	options := SetQueryOptions(o)

	profile, err := (*kuzzle.Kuzzle)(k.instance).Security.GetProfile(C.GoString(id), options)
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, profile, nil)

	return result
}

//export kuzzle_security_get_profile_rights
func kuzzle_security_get_profile_rights(k *C.kuzzle, id *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetProfileRights(C.GoString(id), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_profile_mapping
func kuzzle_security_get_profile_mapping(k *C.kuzzle, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetProfileMapping(SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_role
func kuzzle_security_get_role(k *C.kuzzle, id *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := (*kuzzle.Kuzzle)(k.instance).Security.GetRole(C.GoString(id), options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role, nil)

	return result
}

//export kuzzle_security_get_role_mapping
func kuzzle_security_get_role_mapping(k *C.kuzzle, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetRoleMapping(SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_user
func kuzzle_security_get_user(k *C.kuzzle, id *C.char, o *C.query_options) *C.user_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetUser(C.GoString(id), SetQueryOptions(o))
	return goToCUserResult(k, res, err)
}

//export kuzzle_security_get_user_rights
func kuzzle_security_get_user_rights(k *C.kuzzle, id *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetUserRights(C.GoString(id), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_user_mapping
func kuzzle_security_get_user_mapping(k *C.kuzzle, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetUserMapping(SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_credential_fields
func kuzzle_security_get_credential_fields(k *C.kuzzle, strategy *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetCredentialFields(C.GoString(strategy), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_all_credential_fields
func kuzzle_security_get_all_credential_fields(k *C.kuzzle, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetAllCredentialFields(SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_credential
func kuzzle_security_get_credential(k *C.kuzzle, strategy, id *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetCredentials(C.GoString(strategy), C.GoString(id), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_get_credential_by_id
func kuzzle_security_get_credential_by_id(k *C.kuzzle, strategy, id *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.GetCredentialsByID(C.GoString(strategy), C.GoString(id), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_search_profiles
func kuzzle_security_search_profiles(k *C.kuzzle, body *C.char, o *C.query_options) *C.search_profiles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchProfiles(json.RawMessage(C.GoString(body)), options)

	return goToCProfileSearchResult(k, res, err)
}

//export kuzzle_security_search_roles
func kuzzle_security_search_roles(k *C.kuzzle, body *C.char, o *C.query_options) *C.search_roles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchRoles(json.RawMessage(C.GoString(body)), options)

	return goToCRoleSearchResult(k, res, err)
}

//export kuzzle_security_search_users
func kuzzle_security_search_users(k *C.kuzzle, body *C.char, o *C.query_options) *C.search_users_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchUsers(json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	return goToCUserSearchResult(k, res, err)
}

//export kuzzle_security_delete_role
func kuzzle_security_delete_role(k *C.kuzzle, id *C.char, o *C.query_options) *C.string_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.DeleteRole(C.GoString(id), options)
	return goToCStringResult(&res, err)
}

//export kuzzle_security_delete_profile
func kuzzle_security_delete_profile(k *C.kuzzle, id *C.char, o *C.query_options) *C.string_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.DeleteProfile(C.GoString(id), options)
	return goToCStringResult(&res, err)
}

//export kuzzle_security_delete_user
func kuzzle_security_delete_user(k *C.kuzzle, id *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.DeleteUser(C.GoString(id), SetQueryOptions(o))
	return goToCStringResult(&res, err)
}

//export kuzzle_security_delete_credentials
func kuzzle_security_delete_credentials(k *C.kuzzle, strategy, id *C.char, o *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).Security.DeleteCredentials(C.GoString(strategy), C.GoString(id), SetQueryOptions(o))
	return goToCVoidResult(err)
}

//export kuzzle_security_create_profile
func kuzzle_security_create_profile(k *C.kuzzle, id, body *C.char, o *C.query_options) *C.profile_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateProfile(C.GoString(id), json.RawMessage(C.GoString(body)), options)

	return goToCProfileResult(k, res, err)
}

//export kuzzle_security_create_role
func kuzzle_security_create_role(k *C.kuzzle, id, body *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateRole(C.GoString(id), json.RawMessage(C.GoString(body)), options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role, nil)

	return result
}

//export kuzzle_security_create_user
func kuzzle_security_create_user(k *C.kuzzle, body *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateUser(json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_create_credentials
func kuzzle_security_create_credentials(k *C.kuzzle, cstrategy, id, body *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateCredentials(C.GoString(cstrategy), C.GoString(id), json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_create_restricted_user
func kuzzle_security_create_restricted_user(k *C.kuzzle, body *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateRestrictedUser(json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_create_first_admin
func kuzzle_security_create_first_admin(k *C.kuzzle, body *C.char, o *C.query_options) *C.string_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateFirstAdmin(json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	var result string
	json.Unmarshal(res, result)
	return goToCStringResult(&result, err)
}

//export kuzzle_security_create_or_replace_profile
func kuzzle_security_create_or_replace_profile(k *C.kuzzle, id, body *C.char, o *C.query_options) *C.profile_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateOrReplaceProfile(C.GoString(id), json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	return goToCProfileResult(k, res, err)
}

//export kuzzle_security_create_or_replace_role
func kuzzle_security_create_or_replace_role(k *C.kuzzle, id, body *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	role, err := (*kuzzle.Kuzzle)(k.instance).Security.CreateOrReplaceRole(C.GoString(id), json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role, nil)

	return result
}

//export kuzzle_security_has_credentials
func kuzzle_security_has_credentials(k *C.kuzzle, strategy *C.char, id *C.char, o *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.HasCredentials(C.GoString(strategy), C.GoString(id), SetQueryOptions(o))
	return goToCBoolResult(res, err)
}

//export kuzzle_security_replace_user
func kuzzle_security_replace_user(k *C.kuzzle, id, content *C.char, o *C.query_options) *C.user_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.ReplaceUser(C.GoString(id), json.RawMessage(C.GoString(content)), SetQueryOptions(o))
	return goToCUserResult(k, res, err)
}

//export kuzzle_security_update_credentials
func kuzzle_security_update_credentials(k *C.kuzzle, strategy *C.char, id *C.char, body *C.char, o *C.query_options) *C.void_result {
	err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateCredentials(C.GoString(strategy), C.GoString(id), json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	return goToCVoidResult(err)
}

//export kuzzle_security_update_profile
func kuzzle_security_update_profile(k *C.kuzzle, id *C.char, body *C.char, o *C.query_options) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	options := SetQueryOptions(o)

	profile, err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateProfile(C.GoString(id), json.RawMessage(C.GoString(body)), options)
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, profile, nil)

	return result
}

//export kuzzle_security_update_profile_mapping
func kuzzle_security_update_profile_mapping(k *C.kuzzle, body *C.char, o *C.query_options) *C.void_result {
	options := SetQueryOptions(o)
	err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateProfileMapping(json.RawMessage(C.GoString(body)), options)
	return goToCVoidResult(err)
}

//export kuzzle_security_update_role
func kuzzle_security_update_role(k *C.kuzzle, id *C.char, body *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateRole(C.GoString(id), json.RawMessage(C.GoString(body)), options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role, nil)

	return result
}

//export kuzzle_security_update_role_mapping
func kuzzle_security_update_role_mapping(k *C.kuzzle, body *C.char, o *C.query_options) *C.void_result {
	options := SetQueryOptions(o)
	err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateRoleMapping(json.RawMessage(C.GoString(body)), options)
	return goToCVoidResult(err)
}

//export kuzzle_security_update_user
func kuzzle_security_update_user(k *C.kuzzle, id *C.char, body *C.char, o *C.query_options) *C.user_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateUser(C.GoString(id), json.RawMessage(C.GoString(body)), SetQueryOptions(o))
	return goToCUserResult(k, res, err)
}

//export kuzzle_security_update_user_mapping
func kuzzle_security_update_user_mapping(k *C.kuzzle, body *C.char, o *C.query_options) *C.void_result {
	options := SetQueryOptions(o)
	err := (*kuzzle.Kuzzle)(k.instance).Security.UpdateUserMapping(json.RawMessage(C.GoString(body)), options)
	return goToCVoidResult(err)
}

//export kuzzle_security_is_action_allowed
func kuzzle_security_is_action_allowed(crights **C.user_right, crlength C.uint, controller *C.char, action *C.char, index *C.char, collection *C.char) C.uint {
	rights := make([]*types.UserRights, int(crlength))

	carray := (*[1<<30 - 1]*C.user_right)(unsafe.Pointer(crights))[:int(crlength):int(crlength)]
	for i := 0; i < int(crlength); i++ {
		rights[i] = cToGoUserRigh(carray[i])
	}

	res := security.IsActionAllowed(rights, C.GoString(controller), C.GoString(action), C.GoString(index), C.GoString(collection))
	return C.uint(res)
}

//export kuzzle_security_mdelete_credentials
func kuzzle_security_mdelete_credentials(k *C.kuzzle, ids **C.char, idsSize C.size_t, o *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.MDeleteCredentials(cToGoStrings(ids, idsSize), SetQueryOptions(o))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_security_mdelete_roles
func kuzzle_security_mdelete_roles(k *C.kuzzle, ids **C.char, idsSize C.size_t, o *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.MDeleteRoles(cToGoStrings(ids, idsSize), SetQueryOptions(o))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_security_mdelete_users
func kuzzle_security_mdelete_users(k *C.kuzzle, ids **C.char, idsSize C.size_t, o *C.query_options) *C.string_array_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.MDeleteUsers(cToGoStrings(ids, idsSize), SetQueryOptions(o))
	return goToCStringArrayResult(res, err)
}

//export kuzzle_security_mget_profiles
func kuzzle_security_mget_profiles(k *C.kuzzle, ids **C.char, idsSize C.size_t, o *C.query_options) *C.profiles_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.MGetProfiles(cToGoStrings(ids, idsSize), SetQueryOptions(o))
	return goToCProfilesResult(k, res, err)
}

//export kuzzle_security_mget_roles
func kuzzle_security_mget_roles(k *C.kuzzle, ids **C.char, idsSize C.size_t, o *C.query_options) *C.roles_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.MGetRoles(cToGoStrings(ids, idsSize), SetQueryOptions(o))
	return goToCRolesResult(k, res, err)
}
