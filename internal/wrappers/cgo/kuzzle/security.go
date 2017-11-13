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
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

// --- profile

//export kuzzle_wrapper_security_new_profile
func kuzzle_wrapper_security_new_profile(k *C.kuzzle, id *C.char, policies *C.policy, policies_length C.int) *C.profile {
	cprofile := (*C.profile)(C.calloc(1, C.sizeof_profile))
	cprofile.id = id
	cprofile.policies = policies
	cprofile.policies_length = policies_length
	cprofile.kuzzle = k

	return cprofile
}

//export kuzzle_wrapper_security_destroy_profile
func kuzzle_wrapper_security_destroy_profile(p *C.profile) {
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

//export kuzzle_wrapper_security_fetch_profile
func kuzzle_wrapper_security_fetch_profile(k *C.kuzzle, id *C.char, o *C.query_options) *C.profile_result {
	result := (*C.profile_result)(C.calloc(1, C.sizeof_profile_result))
	options := SetQueryOptions(o)

	profile, err := (*kuzzle.Kuzzle)(k.instance).Security.FetchProfile(C.GoString(id), options)
	if err != nil {
		Set_profile_result_error(result, err)
		return result
	}

	result.profile = goToCProfile(k, profile, nil)

	return result
}

//export kuzzle_wrapper_security_scroll_profiles
func kuzzle_wrapper_security_scroll_profiles(k *C.kuzzle, s *C.char, o *C.query_options) *C.search_profiles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.ScrollProfiles(C.GoString(s), options)

	return goToCProfileSearchResult(k, res, err)
}

//export kuzzle_wrapper_security_search_profiles
func kuzzle_wrapper_security_search_profiles(k *C.kuzzle, f *C.search_filters, o *C.query_options) *C.search_profiles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchProfiles(cToGoSearchFilters(f), options)

	return goToCProfileSearchResult(k, res, err)
}

//export kuzzle_wrapper_security_profile_add_policy
func kuzzle_wrapper_security_profile_add_policy(p *C.profile, policy *C.policy) *C.profile {
	profile := cToGoProfile(p).AddPolicy(cToGoPolicy(policy))

	// @TODO: check if this method is useful and if so, the original pointer to p should be returned instead of a new allocated struct
	return goToCProfile(p.kuzzle, profile, nil)
}

//export kuzzle_wrapper_security_profile_delete
func kuzzle_wrapper_security_profile_delete(p *C.profile, o *C.query_options) *C.string_result {
	options := SetQueryOptions(o)
	profile := cToGoProfile(p)

	res, err := profile.Delete(options)

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_security_profile_save
func kuzzle_wrapper_security_profile_save(p *C.profile, o *C.query_options) *C.profile_result {
	options := SetQueryOptions(o)
	res, err := cToGoProfile(p).Save(options)

	return goToCProfileResult(p.kuzzle, res, err)
}

// --- role

//export kuzzle_wrapper_security_new_role
func kuzzle_wrapper_security_new_role(k *C.kuzzle, id *C.char, c *C.controllers) *C.role {
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

//export kuzzle_wrapper_security_destroy_role
func kuzzle_wrapper_security_destroy_role(r *C.role) {
	if r == nil {
		return
	}

	C.json_object_put(r.controllers)
	C.free(unsafe.Pointer(r))
}

//export kuzzle_wrapper_security_fetch_role
func kuzzle_wrapper_security_fetch_role(k *C.kuzzle, id *C.char, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := (*kuzzle.Kuzzle)(k.instance).Security.FetchRole(C.GoString(id), options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(k, role, nil)

	return result
}

//export kuzzle_wrapper_security_search_roles
func kuzzle_wrapper_security_search_roles(k *C.kuzzle, f *C.search_filters, o *C.query_options) *C.search_roles_result {
	options := SetQueryOptions(o)
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchRoles(cToGoSearchFilters(f), options)

	return goToCRoleSearchResult(k, res, err)
}

//export kuzzle_wrapper_security_role_delete
func kuzzle_wrapper_security_role_delete(r *C.role, o *C.query_options) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))
	opts := SetQueryOptions(o)

	role, err := cToGoRole(r)
	if err != nil {
		Set_string_result_error(result, err)
		return result
	}
	res, err := role.Delete(opts)

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_security_role_save
func kuzzle_wrapper_security_role_save(r *C.role, o *C.query_options) *C.role_result {
	result := (*C.role_result)(C.calloc(1, C.sizeof_role_result))
	options := SetQueryOptions(o)

	role, err := cToGoRole(r)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}
	res, err := role.Save(options)
	if err != nil {
		Set_role_result_error(result, err)
		return result
	}

	result.role = goToCRole(r.kuzzle, res, nil)

	return result
}

// --- user

//export kuzzle_wrapper_security_new_user
func kuzzle_wrapper_security_new_user(k *C.kuzzle, id *C.char, d *C.user_data) *C.user {
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

//export kuzzle_wrapper_security_destroy_user
func kuzzle_wrapper_security_destroy_user(u *C.user) {
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

//export kuzzle_wrapper_security_fetch_user
func kuzzle_wrapper_security_fetch_user(k *C.kuzzle, id *C.char, o *C.query_options) *C.user_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.FetchUser(C.GoString(id), SetQueryOptions(o))
	return goToCUserResult(k, res, err)
}

//export kuzzle_wrapper_security_scroll_users
func kuzzle_wrapper_security_scroll_users(k *C.kuzzle, s *C.char, o *C.query_options) *C.search_users_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.ScrollUsers(C.GoString(s), SetQueryOptions(o))
	return goToCUserSearchResult(k, res, err)
}

//export kuzzle_wrapper_security_search_users
func kuzzle_wrapper_security_search_users(k *C.kuzzle, f *C.search_filters, o *C.query_options) *C.search_users_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).Security.SearchUsers(cToGoSearchFilters(f), SetQueryOptions(o))
	return goToCUserSearchResult(k, res, err)
}

//export kuzzle_wrapper_security_user_create
func kuzzle_wrapper_security_user_create(u *C.user, o *C.query_options) *C.user_result {
	res, err := cToGoUser(u).Create(SetQueryOptions(o))
	return goToCUserResult(u.kuzzle, res, err)
}

//export kuzzle_wrapper_security_user_create_credentials
func kuzzle_wrapper_security_user_create_credentials(u *C.user, cstrategy *C.char, ccredentials *C.json_object, o *C.query_options) *C.json_result {
	strategy := C.GoString(cstrategy)
	credentials := JsonCConvert(ccredentials)
	res, err := cToGoUser(u).CreateCredentials(strategy, credentials, SetQueryOptions(o))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_security_user_create_with_credentials
func kuzzle_wrapper_security_user_create_with_credentials(u *C.user, ccredentials *C.json_object, o *C.query_options) *C.user_result {
	credentials, ok := JsonCConvert(ccredentials).(types.Credentials)
	if !ok {
		return goToCUserResult(u.kuzzle, nil, types.NewError("Invalid credentials given", 400))
	}

	res, err := cToGoUser(u).CreateWithCredentials(credentials, SetQueryOptions(o))

	return goToCUserResult(u.kuzzle, res, err)
}

//export kuzzle_wrapper_security_user_delete
func kuzzle_wrapper_security_user_delete(u *C.user, o *C.query_options) *C.string_result {
	res, err := cToGoUser(u).Delete(SetQueryOptions(o))
	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_security_user_delete_credentials
func kuzzle_wrapper_security_user_delete_credentials(u *C.user, strategy *C.char, o *C.query_options) *C.bool_result {
	res, err := cToGoUser(u).DeleteCredentials(C.GoString(strategy), SetQueryOptions(o))
	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_security_user_get_credentials_info
func kuzzle_wrapper_security_user_get_credentials_info(u *C.user, strategy *C.char, o *C.query_options) *C.json_result {
	res, err := cToGoUser(u).GetCredentialsInfo(C.GoString(strategy), SetQueryOptions(o))
	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_security_user_get_profiles
func kuzzle_wrapper_security_user_get_profiles(u *C.user, o *C.query_options) *C.profiles_result {
	res, err := cToGoUser(u).GetProfiles(SetQueryOptions(o))
	return goToCProfilesResult(u.kuzzle, res, err)
}

//export kuzzle_wrapper_security_user_get_rights
func kuzzle_wrapper_security_user_get_rights(u *C.user, o *C.query_options) *C.user_rights_result {
	res, err := cToGoUser(u).GetRights(SetQueryOptions(o))
	return goToCUserRightsResult(res, err)
}

//export kuzzle_wrapper_security_user_has_credentials
func kuzzle_wrapper_security_user_has_credentials(u *C.user, strategy *C.char, o *C.query_options) *C.bool_result {
	res, err := cToGoUser(u).HasCredentials(C.GoString(strategy), SetQueryOptions(o))
	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_security_user_replace
func kuzzle_wrapper_security_user_replace(u *C.user, o *C.query_options) *C.user_result {
	res, err := cToGoUser(u).Replace(SetQueryOptions(o))
	return goToCUserResult(u.kuzzle, res, err)
}

func kuzzle_wrapper_security_save_restricted(u *C.user, ccredentials *C.json_object, o *C.query_options) *C.user_result {
	credentials, ok := JsonCConvert(ccredentials).(types.Credentials)
	if !ok {
		return goToCUserResult(u.kuzzle, nil, types.NewError("Invalid credentials", 400))
	}

	res, err := cToGoUser(u).SaveRestricted(credentials, SetQueryOptions(o))
	return goToCUserResult(u.kuzzle, res, err)
}

//export kuzzle_wrapper_security_update_credentials
func kuzzle_wrapper_security_update_credentials(u *C.user, strategy *C.char, credentials *C.json_object, o *C.query_options) *C.json_result {
	res, err := cToGoUser(u).UpdateCredentials(C.GoString(strategy), JsonCConvert(credentials), SetQueryOptions(o))
	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_security_is_action_allowed
func kuzzle_wrapper_security_is_action_allowed(crights **C.user_right, crights_length C.uint, controller *C.char, action *C.char, index *C.char, collection *C.char) C.uint {
	rights := make([]*types.UserRights, int(crights_length))

	carray := (*[1<<30 - 1]*C.user_right)(unsafe.Pointer(crights))[:int(crights_length):int(crights_length)]
	for i := 0; i < int(crights_length); i++ {
		rights[i] = cToGoUserRigh(carray[i])
	}

	res := security.IsActionAllowed(rights, C.GoString(controller), C.GoString(action), C.GoString(index), C.GoString(collection))
	return C.uint(res)
}
