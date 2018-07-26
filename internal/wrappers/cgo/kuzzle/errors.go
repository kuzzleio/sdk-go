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
  #cgo CFLAGS: -I../../../headers

  #include "kuzzlesdk.h"
*/
import "C"

import (
	"github.com/kuzzleio/sdk-go/types"
)

// apply a types.KuzzleError on a json_result* C struct
func Set_json_result_error(s *C.json_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a token_validity* C struct
func Set_token_validity_error(s *C.token_validity, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a ack_result* C struct
func Set_ack_result_error(s *C.ack_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a bool_result* C struct
func Set_bool_result_error(s *C.bool_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a kuzzle_response* C struct
func Set_kuzzle_response_error(s *C.kuzzle_response, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a statistics* C struct
func Set_statistics_error(s *C.statistics_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a statistics* C struct
func Set_subscribe_error(s *C.subscribe_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a string_array_result* C struct
func Set_string_array_result_error(s *C.string_array_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a int_result* C struct
func Set_int_result_error(s *C.int_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a int_array_result* C struct
func Set_int_array_result_error(s *C.int_array_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a date_result* C struct
func Set_date_result_error(s *C.date_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a string_result* C struct
func Set_string_result_error(s *C.string_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a shards* C struct
func Set_shards_result_error(s *C.shards_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a search_result* C struct
func Set_search_result_error(s *C.search_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

// apply a types.KuzzleError on a all_statistics_result* C struct
func Set_all_statistics_error(s *C.all_statistics_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_specification_result_err(s *C.specification_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_specification_search_result_error(s *C.specification_search_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_double_result_error(s *C.double_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_json_array_result_error(s *C.json_array_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_profile_result_error(s *C.profile_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_role_result_error(s *C.role_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_search_profiles_result_error(s *C.search_profiles_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_search_roles_result_error(s *C.search_roles_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_search_users_result_error(s *C.search_users_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_user_result_error(s *C.user_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_profiles_result_error(s *C.profiles_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_roles_result_error(s *C.roles_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_user_rights_error(s *C.user_rights_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_notification_result_error(s *C.notification_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_collection_entry_error(s *C.collection_entry_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_error_result_error(s *C.error_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func Set_room_result_error(s *C.room_result, err error) {
	setErr(&s.status, &s.error, &s.stack, err)
}

func setErr(status *C.int, error **C.char, stack **C.char, err error) {
	kuzzleError := err.(types.KuzzleError)
	*status = C.int(kuzzleError.Status)
	*error = C.CString(kuzzleError.Message)

	if len(kuzzleError.Stack) > 0 {
		*stack = C.CString(kuzzleError.Stack)
	}
}
