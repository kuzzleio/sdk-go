package main

/*
  #cgo CFLAGS: -I../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include <errno.h>
  #include "kuzzlesdk.h"
  #include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_set_jwt
func kuzzle_wrapper_set_jwt(k *C.kuzzle, token *C.char) {
	(*kuzzle.Kuzzle)(k.instance).SetJwt(C.GoString(token))
}

//export kuzzle_wrapper_unset_jwt
func kuzzle_wrapper_unset_jwt(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).UnsetJwt()
}

// Allocates memory
//export kuzzle_wrapper_get_jwt
func kuzzle_wrapper_get_jwt(k *C.kuzzle) *C.char {
	return C.CString((*kuzzle.Kuzzle)(k.instance).GetJwt())
}

//export kuzzle_wrapper_login
func kuzzle_wrapper_login(k *C.kuzzle, strategy *C.char, credentials *C.json_object, expires_in *C.int) *C.string_result {
	var expire int
	if expires_in != nil {
		expire = int(*expires_in)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), &expire)

	return goToCStringResult(&res, err)
}

//export kuzzle_wrapper_logout
func kuzzle_wrapper_logout(k *C.kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Logout()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_wrapper_check_token
func kuzzle_wrapper_check_token(k *C.kuzzle, token *C.char) *C.token_validity {
	result := (*C.token_validity)(C.calloc(1, C.sizeof_token_validity))

	res, err := (*kuzzle.Kuzzle)(k.instance).CheckToken(C.GoString(token))
	if err != nil {
		Set_token_validity_error(result, err)
		return result
	}

	result.valid = C.bool(res.Valid)
	result.state = C.CString(res.State)
	result.expires_at = C.ulonglong(res.ExpiresAt)

	return result
}

//export kuzzle_wrapper_create_my_credentials
func kuzzle_wrapper_create_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).CreateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_delete_my_credentials
func kuzzle_wrapper_delete_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).DeleteMyCredentials(
		C.GoString(strategy),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_get_my_credentials
func kuzzle_wrapper_get_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyCredentials(
		C.GoString(strategy),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_update_my_credentials
func kuzzle_wrapper_update_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_validate_my_credentials
func kuzzle_wrapper_validate_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ValidateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_get_my_rights
func kuzzle_wrapper_get_my_rights(k *C.kuzzle, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyRights(SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_update_self
func kuzzle_wrapper_update_self(k *C.kuzzle, data *C.user_data, options *C.query_options) *C.json_result {
	userData, err := cToGoUserData(data)
	if err != nil {
		return goToCJsonResult(nil, err)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateSelf(
		userData,
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_wrapper_who_am_i
func kuzzle_wrapper_who_am_i(k *C.kuzzle) *C.user_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).WhoAmI()

	return goToCUserResult(k, res, err)
}
