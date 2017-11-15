package main

/*
  #cgo CFLAGS: -std=c99 -I../../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include "kuzzlesdk.h"

  static void free_char_array(char **arr, size_t length) {
    if (arr != NULL) {
      for(int i = 0; i < length; i++) {
        free(arr[i]);
      }

      free(arr);
    }
  }
*/
import "C"

import (
	"unsafe"
)

//export destroy_kuzzle_request
func destroy_kuzzle_request(st *C.kuzzle_request) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.id))
		C.free(unsafe.Pointer(st.scroll))
		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st.strategy))
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		C.free(unsafe.Pointer(st.member))
		C.free(unsafe.Pointer(st.member1))
		C.free(unsafe.Pointer(st.member2))
		C.free(unsafe.Pointer(st.unit))
		C.free(unsafe.Pointer(st.field))
		C.free(unsafe.Pointer(st.subcommand))
		C.free(unsafe.Pointer(st.pattern))
		C.free(unsafe.Pointer(st.min))
		C.free(unsafe.Pointer(st.max))
		C.free(unsafe.Pointer(st.limit))
		C.free(unsafe.Pointer(st.match))

		C.free_char_array(st.members, st.members_length)
		C.free_char_array(st.keys, st.keys_length)
		C.free_char_array(st.fields, st.fields_length)

		kuzzle_wrapper_free_json_object(st.body)
		kuzzle_wrapper_free_json_object(st.volatiles)
		kuzzle_wrapper_free_json_object(st.options)

		C.free(unsafe.Pointer(st))
	}

}

//export destroy_query_object
func destroy_query_object(st *C.query_object) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.query)
		C.free(unsafe.Pointer(st.request_id))

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_offline_queue
func destroy_offline_queue(st *C.offline_queue) {
	if st != nil && st.queries != nil {
		queries := (*[1<<30 - 1]*C.query_object)(unsafe.Pointer(st.queries))[:int(st.queries_length):int(st.queries_length)]

		for _, query := range queries {
			destroy_query_object(query)
		}

		C.free(unsafe.Pointer(st.queries))
	}

	C.free(unsafe.Pointer(st))
}

//export destroy_query_options
func destroy_query_options(st *C.query_options) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll))
		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st.refresh))
		C.free(unsafe.Pointer(st.if_exist))
		kuzzle_wrapper_free_json_object(st.volatiles)

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_room_options
func destroy_room_options(st *C.room_options) {
	if st != nil {
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		kuzzle_wrapper_free_json_object(st.volatiles)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_options
func destroy_options(st *C.options) {
	if st != nil {
		C.free(unsafe.Pointer(st.refresh))
		C.free(unsafe.Pointer(st.default_index))
		kuzzle_wrapper_free_json_object(st.headers)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_meta
func destroy_meta(st *C.meta) {
	if st != nil {
		C.free(unsafe.Pointer(st.author))
		C.free(unsafe.Pointer(st.updater))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_policy_restriction(st *C.policy_restriction) {
	if st != nil {
		C.free(unsafe.Pointer(st.index))
		C.free_char_array(st.collections, st.collections_length)
	}
}

//export destroy_policy_restriction
func destroy_policy_restriction(st *C.policy_restriction) {
	_free_policy_restriction(st)
	C.free(unsafe.Pointer(st))
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_policy(st *C.policy) {
	if st != nil {
		C.free(unsafe.Pointer(st.role_id))

		if st.restricted_to != nil {
			restrictions := (*[1<<30 - 1]C.policy_restriction)(unsafe.Pointer(st.restricted_to))[:int(st.restricted_to_length):int(st.restricted_to_length)]

			for _, restriction := range restrictions {
				_free_policy_restriction(&restriction)
			}

			C.free(unsafe.Pointer(st.restricted_to))
		}
	}
}

//export destroy_policy
func destroy_policy(st *C.policy) {
	_free_policy(st)
	C.free(unsafe.Pointer(st))
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_profile(st *C.profile) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))

		if st.policies != nil {
			policies := (*[1<<30 - 1]C.policy)(unsafe.Pointer(st.policies))[:int(st.policies_length):int(st.policies_length)]

			for _, policy := range policies {
				_free_policy(&policy)
			}

			C.free(unsafe.Pointer(st.policies))
		}
	}
}

//export destroy_profile
func destroy_profile(st *C.profile) {
	_free_profile(st)
	C.free(unsafe.Pointer(st))
}

//do not export
func _free_role(st *C.role) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		kuzzle_wrapper_free_json_object(st.controllers)
	}
}

//export destroy_role
func destroy_role(st *C.role) {
	_free_role(st)
	C.free(unsafe.Pointer(st))
}

//do not export
func _free_user(st *C.user) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		kuzzle_wrapper_free_json_object(st.content)
		C.free_char_array(st.profile_ids, st.profile_ids_length)
	}
}

//export destroy_user
func destroy_user(st *C.user) {
	_free_user(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_user_data
func destroy_user_data(st *C.user_data) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.content)
		C.free_char_array(st.profile_ids, st.profile_ids_length)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_collection
func destroy_collection(st *C.collection) {
	if st != nil {
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_document(st *C.document) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.shards))
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.collection))

		kuzzle_wrapper_free_json_object(st.content)

		destroy_meta(st.meta)
		destroy_collection(st._collection)
	}
}

//export destroy_document
func destroy_document(st *C.document) {
	_free_document(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_document_result
func destroy_document_result(st *C.document_result) {
	if st != nil {
		destroy_document(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_notification_content
func destroy_notification_content(st *C.notification_content) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		destroy_meta(st.meta)
		kuzzle_wrapper_free_json_object(st.content)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_notification_result
func destroy_notification_result(st *C.notification_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.protocol))
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		C.free(unsafe.Pointer(st.n_type))
		C.free(unsafe.Pointer(st.room_id))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))

		kuzzle_wrapper_free_json_object(st.volatiles)

		destroy_notification_content(st.result)

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_profile_result
func destroy_profile_result(st *C.profile_result) {
	if st != nil {
		destroy_profile(st.profile)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_profiles_result
func destroy_profiles_result(st *C.profiles_result) {
	if st != nil {
		if st.profiles != nil {
			profiles := (*[1<<30 - 1]C.profile)(unsafe.Pointer(st.profiles))[:int(st.profiles_length):int(st.profiles_length)]

			for _, profile := range profiles {
				_free_profile(&profile)
			}

			C.free(unsafe.Pointer(st.profiles))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_role_result
func destroy_role_result(st *C.role_result) {
	if st != nil {
		destroy_role(st.role)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_user_right(st *C.user_right) {
	if st != nil {
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.value))
	}
}

//export destroy_user_right
func destroy_user_right(st *C.user_right) {
	_free_user_right(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_user_rights_result
func destroy_user_rights_result(st *C.user_rights_result) {
	if st != nil {
		if st.user_rights != nil {
			rights := (*[1<<30 - 1]C.user_right)(unsafe.Pointer(st.user_rights))[:int(st.user_rights_length):int(st.user_rights_length)]

			for _, right := range rights {
				_free_user_right(&right)
			}

			C.free(unsafe.Pointer(st.user_rights))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_user_result
func destroy_user_result(st *C.user_result) {
	if st != nil {
		destroy_user(st.user)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_statistics(st *C.statistics) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.completed_requests)
		kuzzle_wrapper_free_json_object(st.connections)
		kuzzle_wrapper_free_json_object(st.failed_requests)
		kuzzle_wrapper_free_json_object(st.ongoing_requests)
	}
}

//export destroy_statistics
func destroy_statistics(st *C.statistics) {
	_free_statistics(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_statistics_result
func destroy_statistics_result(st *C.statistics_result) {
	if st != nil {
		destroy_statistics(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_all_statistics_result
func destroy_all_statistics_result(st *C.all_statistics_result) {
	if st != nil {
		if st.result != nil {
			stats := (*[1<<30 - 1]C.statistics)(unsafe.Pointer(st.result))

			for _, stat := range stats {
				_free_statistics(&stat)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_geopos_result
func destroy_geopos_result(st *C.geopos_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_token_validity
func destroy_token_validity(st *C.token_validity) {
	if st != nil {
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_kuzzle_response
func destroy_kuzzle_response(st *C.kuzzle_response) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.room_id))
		C.free(unsafe.Pointer(st.channel))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))

		kuzzle_wrapper_free_json_object(st.result)
		kuzzle_wrapper_free_json_object(st.volatiles)

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_json_result
func destroy_json_result(st *C.json_result) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_json_array_result
func destroy_json_array_result(st *C.json_array_result) {
	if st != nil {
		if st.result != nil {
			jobjects := (*[1<<30 - 1]*C.json_object)(unsafe.Pointer(st.result))[:int(st.result_length):int(st.result_length)]

			for _, jobject := range jobjects {
				kuzzle_wrapper_free_json_object(jobject)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_bool_result
func destroy_bool_result(st *C.bool_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_int_result
func destroy_int_result(st *C.int_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_double_result
func destroy_double_result(st *C.double_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_int_array_result
func destroy_int_array_result(st *C.int_array_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_string_result
func destroy_string_result(st *C.string_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_string_array_result
func destroy_string_array_result(st *C.string_array_result) {
	if st != nil {
		C.free_char_array(st.result, st.result_length)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_search_filters
func destroy_search_filters(st *C.search_filters) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.query)
		kuzzle_wrapper_free_json_object(st.sort)
		kuzzle_wrapper_free_json_object(st.aggregations)
		kuzzle_wrapper_free_json_object(st.search_after)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_document_search
func destroy_document_search(st *C.document_search) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll_id))

		if st.hits != nil {
			hits := (*[1<<30 - 1]C.document)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, document := range hits {
				_free_document(&document)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_profile_search
func destroy_profile_search(st *C.profile_search) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll_id))

		if st.hits != nil {
			hits := (*[1<<30 - 1]C.profile)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, profile := range hits {
				_free_profile(&profile)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_role_search
func destroy_role_search(st *C.role_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.role)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, role := range hits {
				_free_role(&role)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export destroy_ack_result
func destroy_ack_result(st *C.ack_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_shards_result
func destroy_shards_result(st *C.shards_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_specification
func destroy_specification(st *C.specification) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.fields)
		kuzzle_wrapper_free_json_object(st.validators)
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_specification_entry(st *C.specification_entry) {
	if st != nil {
		destroy_specification(st.validation)
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
	}
}

//export destroy_specification_entry
func destroy_specification_entry(st *C.specification_entry) {
	_free_specification_entry(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_specification_result
func destroy_specification_result(st *C.specification_result) {
	if st != nil {
		destroy_specification(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_search_result
func destroy_search_result(st *C.search_result) {
	if st != nil {
		destroy_document_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_search_profiles_result
func destroy_search_profiles_result(st *C.search_profiles_result) {
	if st != nil {
		destroy_profile_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_search_roles_result
func destroy_search_roles_result(st *C.search_roles_result) {
	if st != nil {
		destroy_role_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_specification_search
func destroy_specification_search(st *C.specification_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.specification_entry)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, entry := range hits {
				_free_specification_entry(&entry)
			}

			C.free(unsafe.Pointer(st.hits))
			C.free(unsafe.Pointer(st.scroll_id))
			C.free(unsafe.Pointer(st))
		}
	}
}

//export destroy_specification_search_result
func destroy_specification_search_result(st *C.specification_search_result) {
	if st != nil {
		destroy_specification_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_mapping
func destroy_mapping(st *C.mapping) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.mapping)
		destroy_collection(st.collection)
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_mapping_result
func destroy_mapping_result(st *C.mapping_result) {
	if st != nil {
		destroy_mapping(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_void_result
func destroy_void_result(st *C.void_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_collection_entry(st *C.collection_entry) {
	if st != nil {
		C.free(unsafe.Pointer(st.name))
	}
}

//export destroy_collection_entry
func destroy_collection_entry(st *C.collection_entry) {
	_free_collection_entry(st)
	C.free(unsafe.Pointer(st))
}

//export destroy_collection_entry_result
func destroy_collection_entry_result(st *C.collection_entry_result) {
	if st != nil {
		if st.result != nil {
			entries := (*[1<<30 - 1]C.collection_entry)(unsafe.Pointer(st.result))[:int(st.result_length):int(st.result_length)]

			for _, entry := range entries {
				_free_collection_entry(&entry)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_user_search
func destroy_user_search(st *C.user_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.user)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, user := range hits {
				_free_user(&user)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st))
	}
}

//export destroy_search_users_result
func destroy_search_users_result(st *C.search_users_result) {
	if st != nil {
		destroy_user_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}
