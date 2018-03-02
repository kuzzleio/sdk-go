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
	"unsafe"

	"github.com/kuzzleio/sdk-go/auth"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

// map which stores instances to keep references in case the gc passes
var authInstances map[interface{}]bool

//register new instance of server
func registerAuth(instance interface{}) {
	authInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterAuth
func unregisterAuth(a *C.auth) {
	delete(authInstances, (*auth.Auth)(a.instance))
}

// Allocates memory
//export kuzzle_new_auth
func kuzzle_new_auth(a *C.auth, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	auth := auth.NewAuth(kuz)

	if authInstances == nil {
		authInstances = make(map[interface{}]bool)
	}

	a.instance = unsafe.Pointer(auth)
	a.kuzzle = k

	registerAuth(a)
}

//export kuzzle_login
func kuzzle_login(k *C.kuzzle, strategy *C.char, credentials *C.json_object, expires_in *C.int) *C.string_result {
	var expire int
	if expires_in != nil {
		expire = int(*expires_in)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).Login(C.GoString(strategy), JsonCConvert(credentials).(map[string]interface{}), &expire)

	return goToCStringResult(&res, err)
}

//export kuzzle_logout
func kuzzle_logout(k *C.kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Logout()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_check_token
func kuzzle_check_token(a *C.auth, token *C.char) *C.token_validity {
	result := (*C.token_validity)(C.calloc(1, C.sizeof_token_validity))

	res, err := (*auth.Auth)(a.instance).CheckToken(C.GoString(token))
	if err != nil {
		Set_token_validity_error(result, err)
		return result
	}

	result.valid = C.bool(res.Valid)
	result.state = C.CString(res.State)
	result.expires_at = C.ulonglong(res.ExpiresAt)

	return result
}

//export kuzzle_create_my_credentials
func kuzzle_create_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).CreateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_delete_my_credentials
func kuzzle_delete_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).DeleteMyCredentials(
		C.GoString(strategy),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_get_my_credentials
func kuzzle_get_my_credentials(k *C.kuzzle, strategy *C.char, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyCredentials(
		C.GoString(strategy),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_update_my_credentials
func kuzzle_update_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_validate_my_credentials
func kuzzle_validate_my_credentials(k *C.kuzzle, strategy *C.char, credentials *C.json_object, options *C.query_options) *C.bool_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).ValidateMyCredentials(
		C.GoString(strategy),
		JsonCConvert(credentials).(map[string]interface{}),
		SetQueryOptions(options))

	return goToCBoolResult(res, err)
}

//export kuzzle_get_my_rights
func kuzzle_get_my_rights(k *C.kuzzle, options *C.query_options) *C.json_result {
	res, err := (*kuzzle.Kuzzle)(k.instance).GetMyRights(SetQueryOptions(options))

	return goToCJsonResult(res, err)
}

//export kuzzle_update_self
func kuzzle_update_self(k *C.kuzzle, data *C.user_data, options *C.query_options) *C.json_result {
	userData, err := cToGoUserData(data)
	if err != nil {
		return goToCJsonResult(nil, err)
	}

	res, err := (*kuzzle.Kuzzle)(k.instance).UpdateSelf(
		userData,
		SetQueryOptions(options))

	return goToCJsonResult(res, err)
}
